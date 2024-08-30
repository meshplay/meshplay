package connections

import (
	"fmt"
	"net/http"

	"github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root/config"
	"github.com/khulnasoft/meshplay/meshplayctl/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteConnectionCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a connection",
	Long: `Delete
a connection`,

	Example: `
// Delete a connection
meshplayctl exp connections delete [connection_id]
`,

	Args: func(_ *cobra.Command, args []string) error {
		const errMsg = "Usage: meshplayctl exp connections delete \nRun 'meshplayctl exp connections delete --help' to see detailed help message"
		if len(args) != 1 {
			return utils.ErrInvalidArgument(errors.New(errMsg))
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		mctlCfg, err := config.GetMeshplayCtl(viper.GetViper())
		if err != nil {
			return utils.ErrLoadConfig(err)
		}

		baseUrl := mctlCfg.GetBaseMeshplayURL()
		url := fmt.Sprintf("%s/api/integrations/connections/%s", baseUrl, args[0])
		req, err := utils.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return err
		}

		resp, err := utils.MakeRequest(req)
		if err != nil {
			return err
		}

		// defers the closing of the response body after its use, ensuring that the resources are properly released.
		defer resp.Body.Close()

		// Check if the response status code is 200
		if resp.StatusCode == http.StatusOK {
			utils.Log.Info("Connection deleted successfully")
			return nil
		}

		return utils.ErrBadRequest(errors.New(fmt.Sprintf("failed to delete connection with id %s", args[0])))
	},
}
