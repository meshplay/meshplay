package model

import (
	"fmt"

	"github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root/config"
	"github.com/khulnasoft/meshplay/meshplayctl/pkg/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// meshplayctl model export <designname>
var exportModal = &cobra.Command{
	Use:   "export",
	Short: "export registered models",
	Long:  "export the registered model to the specified output type",
	Example: `
// Export a model by name
meshplayctl model export [model-name] -o [oci/json/yaml] (default is oci)
meshplayctl model export [model-name] -l /home/meshplay/
meshplayctl model export [model-name] --discard-components --discard-relationships
    `,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		//Check prerequisite
		mctlCfg, err := config.GetMeshplayCtl(viper.GetViper())
		if err != nil {
			return err
		}
		err = utils.IsServerRunning(mctlCfg.GetBaseMeshplayURL())
		if err != nil {
			return err
		}
		ctx, err := mctlCfg.GetCurrentContext()
		if err != nil {
			return err
		}
		err = ctx.ValidateVersion()
		if err != nil {
			return err
		}
		return nil
	},
	Args: func(_ *cobra.Command, args []string) error {
		const errMsg = "Usage: meshplayctl model export [model-name]\nRun 'meshplayctl model export --help' to see detailed help message"
		if len(args) == 0 {
			return utils.ErrInvalidArgument(errors.New("Please provide a model name. " + errMsg))
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		mctlCfg, err := config.GetMeshplayCtl(viper.GetViper())
		if err != nil {
			log.Fatalln(err, "error processing config")
		}
		baseUrl := mctlCfg.GetBaseMeshplayURL()
		modelName := args[0]
		url := fmt.Sprintf("%s/api/meshmodels/models/%s?components=%t&relationships=%t&%s", baseUrl, modelName, !discardComponentsFlag, !discardRelationshipsFlag, utils.GetPageQueryParameter(cmd, pageNumberFlag))
		return exportModel(args[0], cmd, url, false)
	},
}
