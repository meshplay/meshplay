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

var listModelCmd = &cobra.Command{
	Use:   "list",
	Short: "list registered models",
	Long:  "list name of all registered models",
	Example: `
// View list of models
meshplayctl model list

// View list of models with specified page number (25 models per page)
meshplayctl model list --page 2

// View number of available models in Meshplay
meshplayctl model list --count
    `,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return errors.New(utils.SystemModelSubError("this command takes no arguments\n", "list"))
		}
		mctlCfg, err := config.GetMeshplayCtl(viper.GetViper())
		if err != nil {
			log.Fatalln(err, "error processing config")
		}

		baseUrl := mctlCfg.GetBaseMeshplayURL()
		url := fmt.Sprintf("%s/api/meshmodels/models?%s", baseUrl, utils.GetPageQueryParameter(cmd, pageNumberFlag))

		return listModel(cmd, url, false)
	},
}
