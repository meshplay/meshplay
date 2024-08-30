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

var (
	silentFlagSet bool
)

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Stop, then start Meshplay",
	Long:  `Restart all Meshplay containers / pods.`,
	Example: `
// Restart all Meshplay containers, their instances and their connected volumes
meshplayctl system restart

// (optional) skip checking for new updates available in Meshplay.
meshplayctl system restart --skip-update
	`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		//Check prerequisite
		hcOptions := &HealthCheckOptions{
			IsPreRunE:  true,
			PrintLogs:  false,
			Subcommand: cmd.Use,
		}
		hc, err := NewHealthChecker(hcOptions)
		if err != nil {
			return ErrHealthCheckFailed(err)
		}
		// execute healthchecks
		err = hc.RunPreflightHealthChecks()
		if err != nil {
			cmd.SilenceUsage = true
		}

		return err
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New(utils.SystemLifeCycleError(fmt.Sprintf("restart takes only one flag. See '%s --help' for more information.\n", cmd.CommandPath()), "restart"))
		}
		return restart()
	},
}

func restart() error {
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

	currPlatform := currCtx.GetPlatform()

	running, err := utils.AreMeshplayComponentsRunning(currPlatform)
	if err != nil {
		return err
	}
	if !running { // Meshplay is not running
		if err := start(); err != nil {
			return ErrRestartMeshplay(err)
		}
	} else {
		if currPlatform == "kubernetes" {
			userResponse := false
			if utils.SilentFlag {
				userResponse = true
			} else {
				// ask user for confirmation
				userResponse = utils.AskForConfirmation("Meshplay deployments will be deleted from your cluster. Are you sure you want to continue")
			}
			if !userResponse {
				log.Info("Restart aborted.")
				return nil
			}
			// take a backup of silentFlag value to pass it to start() function later
			silentFlagSet = utils.SilentFlag
			// skips asking for confirmation in the stop() function
			utils.SilentFlag = true
		}

		log.Info("Restarting Meshplay...")

		if err := stop(); err != nil {
			return ErrRestartMeshplay(err)
		}

		// reset the silent flag to avoid overriding the flag for start command
		utils.SilentFlag = silentFlagSet

		if err := start(); err != nil {
			return ErrRestartMeshplay(err)
		}
	}
	return nil
}

func init() {
	restartCmd.Flags().BoolVarP(&skipUpdateFlag, "skip-update", "", false, "(optional) skip checking for new Meshplay's container images.")
	restartCmd.Flags().StringVar(&providerFlag, "provider", "", "Provider to use with the Meshplay server")
}
