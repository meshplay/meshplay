package connections

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root/config"
	"github.com/khulnasoft/meshplay/meshplayctl/pkg/utils"
	"github.com/khulnasoft/meshplay/server/models/connections"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listConnectionsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the connections",
	Long:  `List all the connections`,
	Example: `
// List all the connections
meshplayctl exp connections list

// List all the connections with page number
meshplayctl exp connections list --page 2
`,

	Args: func(_ *cobra.Command, args []string) error {
		const errMsg = "Usage: meshplayctl exp connection list \nRun 'meshplayctl exp connection list --help' to see detailed help message"
		if len(args) != 0 {
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
		var url string
		if cmd.Flags().Changed("page") {
			url = fmt.Sprintf("%s/api/integrations/connections?page=%d", baseUrl, pageNumberFlag)
		} else {
			url = fmt.Sprintf("%s/api/integrations/connections?pagesize=all", baseUrl)
		}
		req, err := utils.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			utils.Log.Error(err)
			return nil
		}

		resp, err := utils.MakeRequest(req)
		if err != nil {
			utils.Log.Error(err)
			return nil
		}

		// defers the closing of the response body after its use, ensuring that the resources are properly released.
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			utils.Log.Error(err)
			return nil
		}

		connectionsResponse := &connections.ConnectionPage{}
		err = json.Unmarshal(data, connectionsResponse)
		if err != nil {
			utils.Log.Error(err)
			return nil
		}

		header := []string{"id", "Name", "Type", "Status"}
		rows := [][]string{}

		for _, connection := range connectionsResponse.Connections {
			if len(connection.Name) > 0 {
				rows = append(rows, []string{connection.ID.String(), connection.Name, connection.Type, string(connection.Status)})
			}
		}

		if len(rows) == 0 {
			// if no connection is found
			fmt.Println("No connection(s) found")
			return nil
		}

		if cmd.Flags().Changed("page") {
			utils.PrintToTable(header, rows)
		} else {
			maxRowsPerPage := 25
			err := utils.HandlePagination(maxRowsPerPage, "connections", rows, header)
			if err != nil {
				utils.Log.Error(err)
				return err
			}
		}

		return nil
	},
}
