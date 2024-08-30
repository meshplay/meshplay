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
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root/config"
	"github.com/khulnasoft/meshplay/meshplayctl/pkg/utils"
	meshkitutils "github.com/layer5io/meshkit/utils"
	meshkitkube "github.com/layer5io/meshkit/utils/kubernetes"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// runPortForward is used for port-forwarding Meshplay UI via `system dashboard`
	runPortForward bool
	localPort      int
)

// dashboardOptions holds values for command line flags that apply to the dashboard
// command.
type dashboardOptions struct {
	host    string // Host on which server is running inside the pod
	port    int    // The default port on which Meshplay service is listening
	podPort int    // Port on which server is running inside the pod
}

// newDashboardOptions initializes dashboard options with default
// values for host, port, and which dashboard to show. Also, set
// max wait time duration for 300 seconds for the dashboard to
// become available
//
// These options may be overridden on the CLI at run-time
func newDashboardOptions() *dashboardOptions {
	return &dashboardOptions{
		host:    utils.MeshplayDefaultHost,
		port:    utils.MeshplayDefaultPort,
		podPort: 8080,
	}
}

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Open Meshplay UI in browser.",
	Args:  cobra.NoArgs,
	Example: `
// Open Meshplay UI in browser
meshplayctl system dashboard

// Open Meshplay UI in browser and use port-forwarding (if default port is taken already)
meshplayctl system dashboard --port-forward

// Open Meshplay UI in browser and use port-forwarding, listen on port 9081 locally, forwarding traffic to meshplay server in the pod
meshplayctl system dashboard --port-forward -p 9081

// (optional) skip opening of MeshplayUI in browser.
meshplayctl system dashboard --skip-browser`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// check if meshplay is running or not
		mctlCfg, err := config.GetMeshplayCtl(viper.GetViper())
		if err != nil {
			utils.Log.Error(err)
			return nil
		}
		currCtx, err := mctlCfg.GetCurrentContext()
		if err != nil {
			utils.Log.Error(ErrGetCurrentContext(err))
			return nil
		}
		running, _ := utils.IsMeshplayRunning(currCtx.GetPlatform())
		if !running {
			return errors.New(`meshplay server is not running. run "meshplayctl system start" to start meshplay`)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		mctlCfg, err := config.GetMeshplayCtl(viper.GetViper())
		if err != nil {
			utils.Log.Error(err)
			return nil
		}
		currCtx, err := mctlCfg.GetCurrentContext()
		if err != nil {
			utils.Log.Error(ErrGetCurrentContext(err))
			return nil
		}
		log.Debug("Fetching Meshplay-UI endpoint")
		switch currCtx.GetPlatform() {
		case "docker":
			if runPortForward {
				log.Warn("--port-forward is not supported using Docker as Meshplay's deployment platform.")
			}
		case "kubernetes":
			client, err := meshkitkube.New([]byte(""))
			if err != nil {
				return err
			}

			// Run port forwarding for accessing Meshplay UI
			if runPortForward {
				options := newDashboardOptions()

				signals := make(chan os.Signal, 1)
				signal.Notify(signals, os.Interrupt)
				defer signal.Stop(signals)

				portforward, err := utils.NewPortForward(
					cmd.Context(),
					client,
					utils.MeshplayNamespace,
					"meshplay",
					options.host,
					localPort,
					options.podPort,
					false,
				)
				if err != nil {
					utils.Log.Error(ErrInitPortForward(err))
					return nil

				}

				if err = portforward.Init(); err != nil {
					// TODO: consider falling back to an ephemeral port if defaultPort is taken
					return ErrRunPortForward(err)
				}
				log.Info("Starting Port-forwarding for Meshplay UI")

				meshplayURL := portforward.URLFor("")

				// ticker for keeping connection alive with pod each 10 seconds
				ticker := time.NewTicker(10 * time.Second)
				go func() {
					for {
						select {
						case <-signals:
							portforward.Stop()
							ticker.Stop()
							return
						case <-ticker.C:
							keepConnectionAlive(meshplayURL)
						}
					}
				}()
				log.Info(fmt.Sprintf("Forwarding port %v -> %v", options.podPort, localPort))
				log.Info("Meshplay UI available at: ", meshplayURL)
				log.Info("Opening Meshplay UI in the default browser.")
				err = utils.NavigateToBrowser(meshplayURL)
				if err != nil {
					log.Warn("Failed to open Meshplay in browser, please point your browser to " + currCtx.GetEndpoint() + " to access Meshplay.")
				}

				<-portforward.GetStop()
				return nil
			}

			var meshplayEndpoint string
			var endpoint *meshkitutils.Endpoint
			clientset := client.KubeClient
			var opts meshkitkube.ServiceOptions
			opts.Name = "meshplay"
			opts.Namespace = utils.MeshplayNamespace
			opts.APIServerURL = client.RestConfig.Host

			endpoint, err = meshkitkube.GetServiceEndpoint(context.TODO(), clientset, &opts)
			if err != nil {
				utils.Log.Error(err) //the func return a meshkit error
				return nil
			}

			meshplayEndpoint = fmt.Sprintf("%s://%s:%d", utils.EndpointProtocol, endpoint.Internal.Address, endpoint.Internal.Port)
			currCtx.SetEndpoint(meshplayEndpoint)
			if !meshkitutils.TcpCheck(&meshkitutils.HostPort{
				Address: endpoint.Internal.Address,
				Port:    endpoint.Internal.Port,
			}, nil) {
				currCtx.SetEndpoint(fmt.Sprintf("%s://%s:%d", utils.EndpointProtocol, endpoint.External.Address, endpoint.External.Port))
				if !meshkitutils.TcpCheck(&meshkitutils.HostPort{
					Address: endpoint.External.Address,
					Port:    endpoint.External.Port,
				}, nil) {
					u, _ := url.Parse(opts.APIServerURL)
					if meshkitutils.TcpCheck(&meshkitutils.HostPort{
						Address: u.Hostname(),
						Port:    endpoint.External.Port,
					}, nil) {
						meshplayEndpoint = fmt.Sprintf("%s://%s:%d", utils.EndpointProtocol, u.Hostname(), endpoint.External.Port)
						currCtx.SetEndpoint(meshplayEndpoint)
					}
				}
			}

			if err == nil {
				err = config.UpdateContextInConfig(currCtx, mctlCfg.GetCurrentContextName())
				if err != nil {
					utils.Log.Error(err)
					return nil
				}
			}

		}

		if !skipBrowserFlag {
			log.Info("Opening Meshplay (" + currCtx.GetEndpoint() + ") in browser.")
			err = utils.NavigateToBrowser(currCtx.GetEndpoint())
			if err != nil {
				log.Warn("Failed to open Meshplay in your browser, please point your browser to " + currCtx.GetEndpoint() + " to access Meshplay.")
			}
		} else {
			log.Info("Meshplay UI available at: ", currCtx.GetEndpoint())
		}

		return nil
	},
}

// keepConnectionAlive to stop being timed out with port forwarding
func keepConnectionAlive(url string) {
	_, err := http.Get(url)
	if err != nil {
		log.Debugf("connection request failed %v", err)
	}
	log.Debugf("connection request success")
}

func init() {
	dashboardCmd.Flags().BoolVarP(&runPortForward, "port-forward", "", false, "(optional) Use port forwarding to access Meshplay UI")
	dashboardCmd.Flags().IntVarP(&localPort, "port", "p", 9081, "(optional) Local port that is not in use from which traffic is to be forwarded to the server running inside the Pod.")

	dashboardCmd.Flags().BoolVarP(&skipBrowserFlag, "skip-browser", "", false, "(optional) skip opening of MeshplayUI in browser.")
}
