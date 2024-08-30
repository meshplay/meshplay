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
	"os"

	"github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root/config"
	"github.com/khulnasoft/meshplay/meshplayctl/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var (
	providerFlag string
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate to a Meshplay Server",
	Long: `
Authenticate to the Local or a Remote Provider of a Meshplay Server

The authentication mode is web-based browser flow`,
	Args: cobra.MinimumNArgs(0),
	Example: `
// Login with the Meshplay Provider of your choice: the Local Provider or a Remote Provider.
meshplayctl system login

// Login with the Meshplay Provider by specifying it via -p or --provider flag.
meshplayctl system login -p Meshplay
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mctlCfg, err := config.GetMeshplayCtl(viper.GetViper())
		if err != nil {
			return errors.Wrap(err, "error processing config")
		}

		currCtx, err := mctlCfg.GetCurrentContext()
		if err != nil {
			return err
		}

		isRunning, err := utils.IsMeshplayRunning(currCtx.GetPlatform())
		if err != nil {
			log.Error("failed to check Meshplay Server status: ", err)
			return nil
		}

		if !isRunning {
			log.Error(`Meshplay Server is not running. Run "meshplayctl system start" to start Meshplay.`)
			return nil
		}

		var tokenData []byte
		if providerFlag != "" {
			var provider = providerFlag
			tokenData, err = utils.InitiateLogin(mctlCfg, provider)
		} else {
			tokenData, err = utils.InitiateLogin(mctlCfg, "")
		}

		if err != nil {
			log.Printf("authentication failed: Unable to reach Meshplay server at %s. Verify your environment's readiness for a Meshplay deployment by running `meshplayctl system check`.", mctlCfg.GetBaseMeshplayURL())
			return nil
		}

		log.Println("successfully authenticated")

		token, err := mctlCfg.GetTokenForContext(mctlCfg.GetCurrentContextName())
		if err != nil {
			// Attempt to create token if it doesn't already exists
			token.Location = utils.AuthConfigFile

			// Write new entry in the config
			if err := config.AddTokenToConfig(token, utils.DefaultConfigPath); err != nil {
				log.Error("failed to find token path for the current context")
				return nil
			}
		}

		if err := os.WriteFile(token.GetLocation(), tokenData, 0666); err != nil {
			log.Error("failed to write the token to the filesystem: ", err)
		}

		return nil
	},
}

func init() {
	loginCmd.PersistentFlags().StringVarP(&providerFlag, "provider", "p", "", "login Meshplay with specified provider")
}
