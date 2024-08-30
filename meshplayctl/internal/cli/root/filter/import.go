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

package filter

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root/config"
	"github.com/khulnasoft/meshplay/meshplayctl/pkg/utils"
	"github.com/khulnasoft/meshplay/server/models"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg  string
	name string
)

var importCmd = &cobra.Command{
	Use:   "import [URI]",
	Short: "Import a WASM filter",
	Long:  "Import a WASM filter from a URI (http/s) or local filesystem path",
	Example: `
// Import a filter file from local filesystem
meshplayctl filter import /path/to/filter.wasm

// Import a filter file from a remote URI
meshplayctl filter import https://example.com/myfilter.wasm

// Add WASM configuration 
// If the string is a valid file in the filesystem, the file is read and passed as a string. Otherwise, the string is passed as is.
// Use quotes if the string contains spaces
meshplayctl filter import /path/to/filter.wasm --wasm-config [filepath|string]

// Specify the name of the filter to be imported. Use quotes if the name contains spaces
meshplayctl filter import /path/to/filter.wasm --name [string]
	`,
	Args: cobra.MinimumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		mctlCfg, err := config.GetMeshplayCtl(viper.GetViper())
		if err != nil {
			utils.Log.Error(err)
			return nil
		}

		filterURL := mctlCfg.GetBaseMeshplayURL() + "/api/filter"

		if len(args) == 0 {
			return errors.New(utils.FilterImportError("URI is required\nUse 'meshplayctl filter import --help' to display usage guide\n"))
		}

		body := models.MeshplayFilterRequestBody{
			Save:       true,
			FilterData: &models.MeshplayFilterPayload{},
		}

		uri := args[0]

		if validURL := govalidator.IsURL(uri); validURL {
			body.URL = uri
		} else {
			filterFile, err := os.ReadFile(uri)
			if err != nil {
				utils.Log.Error(utils.ErrFileRead(err))
				return nil
			}

			fileInfo, err := os.Stat(uri)
			if err != nil {
				utils.Log.Error(utils.ErrFileRead(err))
				return nil
			}

			content := filterFile

			body.FilterData.Name = fileInfo.Name()
			body.FilterData.FilterFile = content
		}

		if cfg != "" {
			// Check if the config is a file path or a string
			if _, err := os.Stat(cfg); err == nil {
				utils.Log.Info("Reading config file")
				cfgFile, err := os.ReadFile(cfg)
				if err != nil {
					utils.Log.Error(utils.ErrReadConfigFile(err))
					return nil
				}

				content := string(cfgFile)
				body.Config = content
			} else {
				utils.Log.Info("config: ")
				utils.Log.Info(cfg)
				body.Config = cfg
			}
		}

		if name != "" {
			body.FilterData.Name = name
		}

		// Convert the request body to JSON
		marshalledBody, err := json.Marshal(body)

		if err != nil {
			utils.Log.Error(utils.ErrMarshal(err))
			return nil
		}

		req, err := utils.NewRequest("POST", filterURL, bytes.NewBuffer(marshalledBody))
		if err != nil {
			utils.Log.Error(utils.ErrCreatingRequest(err))
			return nil
		}

		resp, err := utils.MakeRequest(req)
		if err != nil {
			utils.Log.Error(utils.ErrCreatingRequest(err))
			return nil
		}

		if resp.StatusCode == 200 {
			utils.Log.Info("filter imported")
		} else {
			utils.Log.Error(utils.ErrResponseStatus(resp.StatusCode))
			return nil
		}

		return nil
	},
}

func init() {
	importCmd.Flags().StringVarP(&cfg, "wasm-config", "w", "", "(optional) WASM configuration filepath/string")
	importCmd.Flags().StringVarP(&name, "name", "n", "", "(optional) filter name")
}
