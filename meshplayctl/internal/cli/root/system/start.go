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
	"bufio"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root/config"
	"github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root/constants"
	pkgconstants "github.com/khulnasoft/meshplay/meshplayctl/pkg/constants"
	"github.com/khulnasoft/meshplay/meshplayctl/pkg/utils"

	dockerCmd "github.com/docker/cli/cli/command"
	cliconfig "github.com/docker/cli/cli/config"
	cliflags "github.com/docker/cli/cli/flags"
	"github.com/docker/docker/api/types"
	dockerconfig "github.com/docker/docker/cli/config"

	meshkitutils "github.com/khulnasoft/meshkit/utils"
	meshkitkube "github.com/khulnasoft/meshkit/utils/kubernetes"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	skipUpdateFlag  bool
	skipBrowserFlag bool
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Meshplay",
	Long:  `Start Meshplay and each of its cloud native components.`,
	Args:  cobra.NoArgs,
	Example: `
// Start meshplay
meshplayctl system start

// (optional) skip opening of MeshplayUI in browser.
meshplayctl system start --skip-browser

// (optional) skip checking for new updates available in Meshplay.
meshplayctl system start --skip-update

// Reset Meshplay's configuration file to default settings.
meshplayctl system start --reset

// Specify Platform to deploy Meshplay to.
meshplayctl system start -p docker

// Specify Provider to use.
meshplayctl system start --provider Meshplay
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
			return err
		}
		cfg, err := config.GetMeshplayCtl(viper.GetViper())
		if err != nil {
			utils.Log.Error(err)
			return nil
		}
		ctx, err := cfg.GetCurrentContext()
		if err != nil {
			utils.Log.Error(ErrGetCurrentContext(err))
			return nil
		}
		err = ctx.ValidateVersion()
		if err != nil {
			utils.Log.Error(err)
			return nil
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := start(); err != nil {
			return errors.Wrap(err, utils.SystemError("failed to start Meshplay"))
		}
		return nil
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		latestVersions, err := meshkitutils.GetLatestReleaseTagsSorted(pkgconstants.GetMeshplayGitHubOrg(), pkgconstants.GetMeshplayGitHubRepo())
		version := constants.GetMeshplayctlVersion()
		if err == nil {
			if len(latestVersions) == 0 {
				log.Warn("no versions found for Meshplay")
				return
			}
			latest := latestVersions[len(latestVersions)-1]
			if latest != version {
				log.Printf("A new release of meshplayctl is available: %s → %s", version, latest)
				log.Printf("https://github.com/meshplay/meshplay/releases/tag/%s", latest)
				log.Print("Check https://docs.meshplay.khulnasofy.com/installation/upgrades#upgrading-meshplay-cli for instructions on how to update meshplayctl\n")
			}
		}
	},
}

func start() error {
	if _, err := os.Stat(utils.MeshplayFolder); os.IsNotExist(err) {
		if err := os.Mkdir(utils.MeshplayFolder, 0777); err != nil {
			return ErrCreateDir(err, utils.MeshplayFolder)
		}
	}

	// Get viper instance used for context
	mctlCfg, err := config.GetMeshplayCtl(viper.GetViper())
	if err != nil {
		return errors.Wrap(err, "error processing config")
	}
	// get the platform, channel and the version of the current context
	// if a temp context is set using the -c flag, use it as the current context
	if tempContext != "" {
		err = mctlCfg.SetCurrentContext(tempContext)
		if err != nil {
			return errors.Wrap(err, "failed to set temporary context")
		}
	}

	currCtx, err := mctlCfg.GetCurrentContext()
	if err != nil {
		return err
	}
	meshplayImageVersion := currCtx.GetVersion()
	if currCtx.GetChannel() == "stable" && currCtx.GetVersion() == "latest" {
		meshplayImageVersion = "latest"
	}

	if utils.PlatformFlag != "" {
		if utils.PlatformFlag == "docker" || utils.PlatformFlag == "kubernetes" {
			currCtx.SetPlatform(utils.PlatformFlag)
		} else {
			return ErrUnsupportedPlatform(utils.PlatformFlag, utils.CfgFile)
		}
	}

	if providerFlag != "" {
		currCtx.SetProvider(providerFlag)
	}

	// update the context to config
	err = config.UpdateContextInConfig(currCtx, mctlCfg.GetCurrentContextName())
	if err != nil {
		return err
	}

	// Reset Meshplay config file to default settings
	if utils.ResetFlag {
		err := resetMeshplayConfig()
		if err != nil {
			return ErrResetMeshconfig(err)
		}
	}

	callbackURL := viper.GetString(pkgconstants.CallbackURLENV)
	providerURL := viper.GetString(pkgconstants.ProviderURLsENV)
	// deploy to platform specified in the config.yaml
	switch currCtx.GetPlatform() {
	case "docker":
		// download the docker-compose.yaml file corresponding to the current version
		if err := utils.DownloadDockerComposeFile(currCtx, true); err != nil {
			return ErrDownloadFile(err, utils.DockerComposeFile)
		}

		// viper instance used for docker compose
		utils.ViperCompose.SetConfigFile(utils.DockerComposeFile)
		err = utils.ViperCompose.ReadInConfig()
		if err != nil {
			return err
		}

		compose := &utils.DockerCompose{}
		err = utils.ViperCompose.Unmarshal(&compose)
		if err != nil {
			return ErrUnmarshalDockerCompose(err, utils.DockerComposeFile)
		}

		//changing the port mapping in docker compose
		//extracting the custom user port from config.yaml
		userPort := strings.Split(currCtx.GetEndpoint(), ":")
		//extracting container port from the docker-compose
		containerPort := strings.Split(utils.Services["meshplay"].Ports[0], ":")
		userPortMapping := userPort[len(userPort)-1] + ":" + containerPort[len(containerPort)-1]
		utils.Services["meshplay"].Ports[0] = userPortMapping

		RequiredService := []string{"meshplay", "watchtower"}

		AllowedServices := map[string]utils.Service{}
		for _, v := range currCtx.GetComponents() {
			if utils.Services[v].Image == "" {
				log.Fatalf("Invalid component specified %s", v)
			}

			temp, ok := utils.Services[v]
			if !ok {
				return errors.New(fmt.Sprintf("No Docker Compose service exists for Meshplay component `%s`.", v))
			}

			spliter := strings.Split(temp.Image, ":")
			temp.Image = fmt.Sprintf("%s:%s-%s", spliter[0], currCtx.GetChannel(), "latest")
			utils.Services[v] = temp
			AllowedServices[v] = utils.Services[v]
			utils.ViperCompose.Set(fmt.Sprintf("services.%s", v), utils.Services[v])
			err = utils.ViperCompose.WriteConfig()
			if err != nil {
				// failure while adding a service to docker compose file is not a fatal error
				// meshplayctl will continue deploying with required services (meshplay, watchtower)
				log.Infof("Encountered an error while adding `%s` service to Docker Compose file. Verify permission to write to `.meshplay/meshplay.yaml` file.", v)
			}
		}

		for _, v := range RequiredService {
			if v == "watchtower" {
				AllowedServices[v] = utils.Services[v]
				continue
			}

			temp, ok := utils.Services[v]
			if !ok {
				return errors.New("unable to extract meshplay version")
			}

			spliter := strings.Split(temp.Image, ":")
			temp.Image = fmt.Sprintf("%s:%s-%s", spliter[0], currCtx.GetChannel(), "latest")
			if v == "meshplay" {
				callbackEnvVaridx, ok := utils.FindInSlice(pkgconstants.CallbackURLENV, temp.Environment)
				if !ok {
					temp.Environment = append(temp.Environment, fmt.Sprintf("%s=%s", pkgconstants.CallbackURLENV, callbackURL))
				} else if callbackURL != "" {
					if ok {
						temp.Environment[callbackEnvVaridx] = fmt.Sprintf("%s=%s", pkgconstants.CallbackURLENV, callbackURL)
					}
				}

				providerEnvVar := currCtx.GetProvider()
				// If user has specified provider using --provider flag use that.
				if providerFlag != "" {
					providerEnvVar = providerFlag
				}
				proivderEnvVaridx, ok := utils.FindInSlice(pkgconstants.ProviderENV, temp.Environment)

				if !ok {
					temp.Environment = append(temp.Environment, fmt.Sprintf("%s=%s", pkgconstants.ProviderENV, providerEnvVar))
				} else if providerEnvVar != "" {
					temp.Environment[proivderEnvVaridx] = fmt.Sprintf("%s=%s", pkgconstants.ProviderENV, providerEnvVar)
				}

				temp.Image = fmt.Sprintf("%s:%s-%s", spliter[0], currCtx.GetChannel(), meshplayImageVersion)
			}
			utils.Services[v] = temp
			AllowedServices[v] = utils.Services[v]
		}

		//////// FLAGS
		// Control whether to pull for new Meshplay container images
		if skipUpdateFlag {
			log.Info("Skipping Meshplay update...")
		} else {
			err := utils.UpdateMeshplayContainers()
			if err != nil {
				return errors.Wrap(err, utils.SystemError("failed to update Meshplay containers"))
			}
		}

		var endpoint meshkitutils.HostPort

		userResponse := false

		//skip asking confirmation if -y flag used or host in meshconfig is already localhost
		if utils.SilentFlag || strings.HasSuffix(userPort[1], "localhost") {
			userResponse = true
		} else {
			// ask user for confirmation
			userResponse = utils.AskForConfirmation("The endpoint address will be changed to localhost. Are you sure you want to continue?")
		}

		if userResponse {
			endpoint.Address = utils.EndpointProtocol + "://localhost"
			currCtx.SetEndpoint(endpoint.Address + ":" + userPort[len(userPort)-1])

			err = config.UpdateContextInConfig(currCtx, mctlCfg.GetCurrentContextName())
			if err != nil {
				return err
			}
		} else {
			endpoint.Address = userPort[0]
		}

		tempPort, err := strconv.Atoi(userPort[len(userPort)-1])
		if err != nil {
			return err
		}
		endpoint.Port = int32(tempPort)

		// group, err := user.LookupGroup("docker")
		// if err != nil {
		// 	return errors.Wrap(err, utils.SystemError("unable to get GID of docker group"))
		// }

		// // Create the group_add option and add GID of docker group to meshplay container
		// groupAdd := viper.GetStringSlice("services.meshplay.group_add")
		// groupAdd = append(groupAdd, group.Gid)
		// utils.ViperCompose.Set("services.meshplay.group_add", groupAdd)

		// // Write the modified configuration back to the Docker Compose file
		// if err := utils.ViperCompose.WriteConfig(); err != nil {
		// 	return errors.Wrap(err, utils.SystemError("unable to add group_add option. Meshplay Server cannot perform this privileged action"))
		// }

		log.Info("Starting Meshplay...")
		start := exec.Command("docker-compose", "-f", utils.DockerComposeFile, "up", "-d")
		start.Stdout = os.Stdout
		start.Stderr = os.Stderr

		if err := start.Run(); err != nil {
			return errors.Wrap(err, utils.SystemError("failed to run meshplay server"))
		}

		checkFlag := 0 //flag to check

		// Get the Docker configuration
		dockerCfg, err := cliconfig.Load(dockerconfig.Dir())
		if err != nil {
			return ErrCreatingDockerClient(err)
		}

		//connection to docker-client
		cli, err := dockerCmd.NewAPIClientFromFlags(cliflags.NewClientOptions(), dockerCfg)
		if err != nil {
			utils.Log.Error(ErrCreatingDockerClient(err))
			return err
		}

		containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
		if err != nil {
			return errors.Wrap(err, utils.SystemError("failed to fetch the list of containers"))
		}

		var mockEndpoint *meshkitutils.MockOptions
		mockEndpoint = nil

		res := meshkitutils.TcpCheck(&endpoint, mockEndpoint)
		if res {
			return errors.New("the endpoint is not accessible")
		}

		//check for container meshplay_meshplay_1 running status
		for _, container := range containers {
			if container.Names[0] == "/meshplay_meshplay_1" {
				//check flag to check successful deployment
				checkFlag = 0
				break
			}

			checkFlag = 1
		}

		//if meshplay_meshplay_1 failed to start showing logs
		//code for logs
		if checkFlag == 1 {
			log.Info("Starting Meshplay logging . . .")
			cmdlog := exec.Command("docker-compose", "-f", utils.DockerComposeFile, "logs", "-f")
			cmdReader, err := cmdlog.StdoutPipe()
			if err != nil {
				return errors.Wrap(err, utils.SystemError("failed to create stdout pipe"))
			}
			scanner := bufio.NewScanner(cmdReader)
			go func() {
				for scanner.Scan() {
					log.Println(scanner.Text())
				}
			}()
			if err := cmdlog.Start(); err != nil {
				return errors.Wrap(err, utils.SystemError("failed to start logging"))
			}
			if err := cmdlog.Wait(); err != nil {
				return errors.Wrap(err, utils.SystemError("failed to wait for command to execute"))
			}
		}

	case "kubernetes":
		kubeClient, err := meshkitkube.New([]byte(""))
		if err != nil {
			return err
		}

		log.Info("Starting Meshplay...")

		spinner := utils.CreateDefaultSpinner("Deploying Meshplay on Kubernetes", "\nMeshplay deployed on Kubernetes.")
		spinner.Start()

		if err := utils.CreateManifestsFolder(); err != nil {
			utils.Log.Error(ErrCreateManifestsFolder(err))
			return err
		}

		// Applying Meshplay Helm charts for installing Meshplay
		if err = applyHelmCharts(kubeClient, currCtx, meshplayImageVersion, false, meshkitkube.INSTALL, callbackURL, providerURL); err != nil {
			return err
		}

		// checking if Meshplay is ready
		time.Sleep(10 * time.Second) // sleeping 10 seconds to countermeasure time to apply helm charts
		ready, err := meshplayReadinessHealthCheck()
		if err != nil {
			log.Info(err)
		}

		spinner.Stop()

		if !ready {
			log.Info("\nFew Meshplay pods have not come up yet.\nPlease check the status of the pods by executing “meshplayctl system status” and Meshplay-UI endpoint with “meshplayctl system dashboard” before using meshplay.")
			return nil
		}
		log.Info("Meshplay is starting...")

		// switch to default case if the platform specified is not supported
	default:
		return fmt.Errorf("the platform %s is not supported currently. The supported platforms are:\ndocker\nkubernetes\nPlease check %s/config.yaml file", currCtx.GetPlatform(), utils.MeshplayFolder)
	}

	// execute dashboard command to fetch and navigate to Meshplay UI
	return dashboardCmd.RunE(nil, nil)
}

func init() {
	startCmd.PersistentFlags().StringVarP(&utils.PlatformFlag, "platform", "p", "", "platform to deploy Meshplay to.")
	startCmd.Flags().BoolVarP(&skipUpdateFlag, "skip-update", "", false, "(optional) skip checking for new Meshplay's container images.")
	startCmd.Flags().BoolVarP(&utils.ResetFlag, "reset", "", false, "(optional) reset Meshplay's configuration file to default settings.")
	startCmd.Flags().BoolVarP(&skipBrowserFlag, "skip-browser", "", false, "(optional) skip opening of MeshplayUI in browser.")
	startCmd.PersistentFlags().StringVar(&providerFlag, "provider", "", "(optional) Defaults to the provider specified in the current context")
}

// Apply Meshplay helm charts
func applyHelmCharts(kubeClient *meshkitkube.Client, currCtx *config.Context, meshplayImageVersion string, dryRun bool, act meshkitkube.HelmChartAction, callbackURL, providerURL string) error {
	// get value overrides to install the helm chart
	overrideValues := utils.SetOverrideValues(currCtx, meshplayImageVersion, callbackURL, providerURL)

	// install the helm charts with specified override values
	var chartVersion string
	if meshplayImageVersion != "latest" {
		chartVersion = meshplayImageVersion
	}
	action := "install"
	if act == meshkitkube.UNINSTALL {
		action = "uninstall"
	}
	errServer := kubeClient.ApplyHelmChart(meshkitkube.ApplyHelmChartConfig{
		Namespace:       utils.MeshplayNamespace,
		ReleaseName:     "meshplay",
		CreateNamespace: true,
		ChartLocation: meshkitkube.HelmChartLocation{
			Repository: utils.HelmChartURL,
			Chart:      utils.HelmChartName,
			Version:    chartVersion,
		},
		OverrideValues: overrideValues,
		Action:         act,
		// the helm chart will be downloaded to ~/.meshplay/manifests if it doesn't exist
		DownloadLocation: path.Join(utils.MeshplayFolder, utils.ManifestsFolder),
		DryRun:           dryRun,
	})
	errOperator := kubeClient.ApplyHelmChart(meshkitkube.ApplyHelmChartConfig{
		Namespace:       utils.MeshplayNamespace,
		ReleaseName:     "meshplay-operator",
		CreateNamespace: true,
		ChartLocation: meshkitkube.HelmChartLocation{
			Repository: utils.HelmChartURL,
			Chart:      utils.HelmChartOperatorName,
			Version:    chartVersion,
		},
		Action: act,
		// the helm chart will be downloaded to ~/.meshplay/manifests if it doesn't exist
		DownloadLocation: path.Join(utils.MeshplayFolder, utils.ManifestsFolder),
		DryRun:           dryRun,
	})
	if errServer != nil && errOperator != nil {
		return fmt.Errorf("could not %s meshplay server: %s\ncould not %s meshplay-operator: %s", action, errServer.Error(), action, errOperator.Error())
	}
	if errServer != nil {
		return fmt.Errorf("%s success for operator but failed for meshplay server: %s", action, errServer.Error())
	}
	if errOperator != nil {
		return fmt.Errorf("%s success for meshplay server but failed for meshplay operator: %s", action, errOperator.Error())
	}
	return nil
}
