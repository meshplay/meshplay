// Copyright Meshplay Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package system

import (
	"fmt"

	"github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root/config"
	"github.com/khulnasoft/meshplay/meshplayctl/pkg/utils"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var linkDocReset = map[string]string{
	"link":    "![reset-usage](/assets/img/meshplayctl/reset.png)",
	"caption": "Usage of meshplayctl system reset",
}

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset Meshplay's configuration",
	Long:  `Reset Meshplay to it's default configuration.`,
	Example: `
// Resets meshplay.yaml file with a copy from Meshplay repo
meshplayctl system reset
	`,
	Annotations: linkDocReset,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return errors.New(utils.SystemLifeCycleError(fmt.Sprintf("this command takes no arguments. See '%s --help' for more information.\n", cmd.CommandPath()), "reset"))
		}
		return resetMeshplayConfig()
	},
}

// resets meshplay config, skips conirmation if skipConfirmation is true
func resetMeshplayConfig() error {
	userResponse := false
	if utils.SilentFlag {
		userResponse = true
	} else {
		// ask user for confirmation
		userResponse = utils.AskForConfirmation("Meshplay config file will be reset to system defaults. Are you sure you want to continue")
	}
	if !userResponse {
		log.Info("Reset aborted.")
		return nil
	}

	// Get viper instance used for context
	mctlCfg, err := config.GetMeshplayCtl(viper.GetViper())
	if err != nil {
		utils.Log.Error(err)
		return nil
	}
	// get the platform, channel and the version of the current context
	// if a temp context is set using the -c flag, use it as the current context
	if tempContext != "" {
		err = mctlCfg.SetCurrentContext(tempContext)
		if err != nil {
			return ErrSettingTemporaryContext(err)
		}
	}

	currCtx, err := mctlCfg.GetCurrentContext()
	if err != nil {
		return ErrRetrievingCurrentContext(err)
	}

	log.Info("Meshplay resetting...\n")
	log.Printf("Current Context: %s", mctlCfg.GetCurrentContextName())
	log.Printf("Channel: %s", currCtx.GetChannel())
	log.Printf("Version: %s", currCtx.GetVersion())
	log.Printf("Platform: %s\n", currCtx.GetPlatform())

	// Reset the config file to the default context
	defaultContext := utils.TemplateContext
	defaultContext.Platform = currCtx.Platform
	err = config.AddContextToConfig(mctlCfg.GetCurrentContextName(), defaultContext, utils.DefaultConfigPath, true, true)
	if err != nil {
		return ErrSettingDefaultContextToConfig(err)
	}

	return fetchManifests(mctlCfg)
}

// Fetches manifests for meshplay components based on the current context
func fetchManifests(mctlCfg *config.MeshplayCtlConfig) error {
	currCtx, err := mctlCfg.GetCurrentContext()
	if err != nil {
		return ErrRetrievingCurrentContext(err)
	}

	switch currCtx.GetPlatform() {
	case "docker":

		log.Printf("Fetching default docker-compose file as per current-context: %s...", mctlCfg.GetCurrentContextName())
		err = utils.DownloadDockerComposeFile(currCtx, true)
		if err != nil {
			return ErrDownloadFile(err, utils.DockerComposeFile)
		}

		err = utils.CreateManifestsFolder()

		if err != nil {
			return ErrCreateManifestsFolder(err)
		}

		log.Printf("...fetching Meshplay Operator manifests for Kubernetes...")
		err = utils.DownloadOperatorManifest()

		if err != nil {
			return ErrDownloadFile(err, "operator manifest")
		}

		log.Info("...meshconfig (" + utils.DockerComposeFile + ") now reset to default settings.")

	case "kubernetes":

		log.Printf("Fetching Meshplay Server and Meshplay Operator manifests for  %s context...", mctlCfg.GetCurrentContextName())
		// fetch the manifest files corresponding to the version specified
		_, err := utils.FetchManifests(currCtx)

		if err != nil {
			return err
		}

		log.Info("...meshconfig has been reset to default settings.")

	default:
		return fmt.Errorf("the platform %s is not supported currently. The supported platforms are:\ndocker\nkubernetes\nPlease check %s/config.yaml file", currCtx.Platform, utils.MeshplayFolder)
	}

	return nil
}
