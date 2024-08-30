package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
	"github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root/config"
	"github.com/khulnasoft/meshplay/meshplayctl/pkg/constants"
	"github.com/khulnasoft/meshplay/server/models"
	"github.com/layer5io/meshkit/encoding"
	"github.com/layer5io/meshkit/logger"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

const (
	// Meshplay Docker Deployment URLs
	dockerComposeWebURL         = "https://api.github.com/repos/docker/compose/releases/latest"
	defaultDockerComposeVersion = "1.24.1/docker-compose"
	dockerComposeBinaryURL      = "https://github.com/docker/compose/releases/download/"
	dockerComposeBinary         = "/usr/local/bin/docker-compose"

	// Meshplay Kubernetes Deployment URLs
	baseConfigURL = "https://raw.githubusercontent.com/layer5io/meshplay-operator/master/config/"
	OperatorURL   = baseConfigURL + "manifests/default.yaml"
	BrokerURL     = baseConfigURL + "samples/meshplay_v1alpha1_broker.yaml"
	MeshsyncURL   = baseConfigURL + "samples/meshplay_v1alpha1_meshsync.yaml"

	// Documentation URLs
	docsBaseURL                    = "https://docs.meshplay.khulnasofy.com/"
	rootUsageURL                   = docsBaseURL + "reference/meshplayctl"
	perfUsageURL                   = docsBaseURL + "reference/meshplayctl/perf"
	systemUsageURL                 = docsBaseURL + "reference/meshplayctl/system"
	systemStopURL                  = docsBaseURL + "reference/meshplayctl/system/stop"
	systemUpdateURL                = docsBaseURL + "reference/meshplayctl/system/update"
	systemResetURL                 = docsBaseURL + "reference/meshplayctl/system/reset"
	systemStatusURL                = docsBaseURL + "reference/meshplayctl/system/status"
	systemRestartURL               = docsBaseURL + "reference/meshplayctl/system/restart"
	meshUsageURL                   = docsBaseURL + "reference/meshplayctl/mesh"
	expUsageURL                    = docsBaseURL + "reference/meshplayctl/exp"
	filterUsageURL                 = docsBaseURL + "reference/meshplayctl/filter"
	filterImportURL                = docsBaseURL + "reference/meshplayctl/filter/import"
	filterDeleteURL                = docsBaseURL + "reference/meshplayctl/filter/delete"
	filterListURL                  = docsBaseURL + "reference/meshplayctl/filter/list"
	filterViewURL                  = docsBaseURL + "reference/meshplayctl/filter/view"
	patternUsageURL                = docsBaseURL + "reference/meshplayctl/pattern"
	patternViewURL                 = docsBaseURL + "reference/meshplayctl/pattern/view"
	patternExportURL               = docsBaseURL + "reference/meshplayctl/pattern/export"
	contextDeleteURL               = docsBaseURL + "reference/meshplayctl/system/context/delete"
	contextViewURL                 = docsBaseURL + "reference/meshplayctl/system/context/view"
	contextCreateURL               = docsBaseURL + "reference/meshplayctl/system/context/create"
	contextUsageURL                = docsBaseURL + "reference/meshplayctl/system/context"
	channelUsageURL                = docsBaseURL + "reference/meshplayctl/system/channel"
	channelSetURL                  = docsBaseURL + "reference/meshplayctl/system/channel/set"
	channelSwitchURL               = docsBaseURL + "reference/meshplayctl/system/channel/switch"
	channelViewURL                 = docsBaseURL + "reference/meshplayctl/system/channel/view"
	providerUsageURL               = docsBaseURL + "reference/meshplayctl/system/provider"
	providerViewURL                = docsBaseURL + "reference/meshplayctl/system/provider/view"
	providerListURL                = docsBaseURL + "reference/meshplayctl/system/provider/list"
	providerSetURL                 = docsBaseURL + "reference/meshplayctl/system/provider/set"
	providerResetURL               = docsBaseURL + "reference/meshplayctl/system/provider/reset"
	providerSwitchURL              = docsBaseURL + "reference/meshplayctl/system/provider/switch"
	tokenUsageURL                  = docsBaseURL + "reference/meshplayctl/system/token"
	modelUsageURL                  = docsBaseURL + "reference/meshplayctl/system/model"
	modelListURL                   = docsBaseURL + "reference/meshplayctl/system/model/list"
	modelImportURl                 = docsBaseURL + "reference/meshplayctl/system/model/import"
	modelViewURL                   = docsBaseURL + "reference/meshplayctl/system/model/view"
	registryUsageURL               = docsBaseURL + "reference/meshplayctl/system/registry"
	relationshipUsageURL           = docsBaseURL + "reference/meshplayctl/relationships"
	cmdRelationshipGenerateDocsURL = docsBaseURL + "reference/meshplayctl/relationships/generate"
	relationshipViewURL            = docsBaseURL + "reference/meshplayctl/relationships/view"
	workspaceUsageURL              = docsBaseURL + "reference/meshplayctl/exp/workspace"
	workspaceCreateURL             = docsBaseURL + "reference/meshplayctl/exp/workspace/create"
	workspaceListURL               = docsBaseURL + "reference/meshplayctl/exp/workspace/list"
	environmentUsageURL            = docsBaseURL + "reference/meshplayctl/exp/environment"
	environmentCreateURL           = docsBaseURL + "reference/meshplayctl/exp/environment/create"
	environmentDeleteURL           = docsBaseURL + "reference/meshplayctl/exp/environment/delete"
	environmentListURL             = docsBaseURL + "reference/meshplayctl/exp/environment/list"
	environmentViewURL             = docsBaseURL + "reference/meshplayctl/exp/environment/view"
	componentUsageURL              = docsBaseURL + "reference/meshplayctl/exp/components"
	componentListURL               = docsBaseURL + "reference/meshplayctl/exp/components/list"
	componentSearchURL             = docsBaseURL + "reference/meshplayctl/exp/components/search"
	componentViewURL               = docsBaseURL + "reference/meshplayctl/exp/components/view"
	connectionUsageURL             = docsBaseURL + "reference/meshplayctl/exp/connections"
	connectionDeleteURL            = docsBaseURL + "reference/meshplayctl/exp/connections/delete"
	connectionListURL              = docsBaseURL + "reference/meshplayctl/exp/connections/list"
	expRelationshipUsageURL        = docsBaseURL + "reference/meshplayctl/exp/relationship"
	expRelationshipGenerateURL     = docsBaseURL + "reference/meshplayctl/exp/relationship/generate"
	expRelationshipViewURL         = docsBaseURL + "reference/meshplayctl/exp/relationship/view"
	expRelationshipListURL         = docsBaseURL + "reference/meshplayctl/exp/relationship/list"

	// Meshplay Server Location
	EndpointProtocol = "http"
)

type cmdType string

const (
	cmdRoot                     cmdType = "root"
	cmdPerf                     cmdType = "perf"
	cmdMesh                     cmdType = "mesh"
	cmdSystem                   cmdType = "system"
	cmdSystemStop               cmdType = "system stop"
	cmdSystemUpdate             cmdType = "system update"
	cmdSystemReset              cmdType = "system reset"
	cmdSystemStatus             cmdType = "system status"
	cmdSystemRestart            cmdType = "system restart"
	cmdExp                      cmdType = "exp"
	cmdFilter                   cmdType = "filter"
	cmdFilterImport             cmdType = "filter import"
	cmdFilterDelete             cmdType = "filter delete"
	cmdFilterList               cmdType = "filter list"
	cmdFilterView               cmdType = "filter view"
	cmdPattern                  cmdType = "pattern"
	cmdPatternView              cmdType = "pattern view"
	cmdPatternExport            cmdType = "pattern export"
	cmdContext                  cmdType = "context"
	cmdContextDelete            cmdType = "delete"
	cmdContextCreate            cmdType = "create"
	cmdContextView              cmdType = "context view"
	cmdChannel                  cmdType = "channel"
	cmdChannelSet               cmdType = "channel set"
	cmdChannelSwitch            cmdType = "channel switch"
	cmdChannelView              cmdType = "channel view"
	cmdProvider                 cmdType = "provider"
	cmdProviderSet              cmdType = "provider set"
	cmdProviderSwitch           cmdType = "provider switch"
	cmdProviderView             cmdType = "provider view"
	cmdProviderList             cmdType = "provider list"
	cmdProviderReset            cmdType = "provider reset"
	cmdToken                    cmdType = "token"
	cmdModel                    cmdType = "model"
	cmdModelList                cmdType = "model list"
	cmdModelImport              cmdType = "model import"
	cmdModelView                cmdType = "model view"
	cmdRegistryPublish          cmdType = "registry publish"
	cmdRegistry                 cmdType = "regisry"
	cmdConnection               cmdType = "connection"
	cmdConnectionList           cmdType = "connection list"
	cmdConnectionDelete         cmdType = "connection delete"
	cmdRelationships            cmdType = "relationships"
	cmdRelationshipGenerateDocs cmdType = "relationships generate docs"
	cmdRelationshipView         cmdType = "relationship view"
	cmdRelationshipSearch       cmdType = "relationship search"
	cmdRelationshipList         cmdType = "relationship list"
	cmdWorkspace                cmdType = "workspace"
	cmdWorkspaceList            cmdType = "workspace list"
	cmdWorkspaceCreate          cmdType = "workspace create"
	cmdEnvironment              cmdType = "environment"
	cmdEnvironmentCreate        cmdType = "environment create"
	cmdEnvironmentDelete        cmdType = "environment delete"
	cmdEnvironmentList          cmdType = "environment list"
	cmdEnvironmentView          cmdType = "environment view"
	cmdComponent                cmdType = "component"
	cmdComponentList            cmdType = "component list"
	cmdComponentSearch          cmdType = "component search"
	cmdComponentView            cmdType = "component view"
	cmdExpRelationship          cmdType = "exp relationship"
	cmdExpRelationshipGenerate  cmdType = "exp relationship generate"
	cmdExpRelationshipView      cmdType = "exp relationship view"
	cmdExpRelationshipList      cmdType = "exp relationship list"
)

const (
	HelmChartURL          = "https://meshplay.khulnasofy.com/charts/"
	HelmChartName         = "meshplay"
	HelmChartOperatorName = "meshplay-operator"
)

var (
	// ResetFlag indicates if a reset is required
	ResetFlag bool
	// SkipResetFlag indicates if fetching the updated manifest files is required
	SkipResetFlag bool
	// MeshplayDefaultHost is the default host on which Meshplay is exposed
	MeshplayDefaultHost = "localhost"
	// MeshplayDefaultPort is the default port on which Meshplay is exposed
	MeshplayDefaultPort = 9081
	// MeshplayEndpoint is the default URL in which Meshplay is exposed
	MeshplayEndpoint = fmt.Sprintf("http://%s:%v", MeshplayDefaultHost, MeshplayDefaultPort)
	// MeshplayFolder is the default relative location of the meshplay config
	// related configuration files.
	MeshplayFolder = ".meshplay"
	// DockerComposeFile is the default location within the MeshplayFolder
	// where the docker compose file is located.
	DockerComposeFile = "meshplay.yaml"
	// AuthConfigFile is the location of the auth file for performing perf testing
	AuthConfigFile = "auth.json"
	// DefaultConfigPath is the detail path to meshplayctl config
	DefaultConfigPath = "config.yaml"
	// MeshplayNamespace is the namespace to which Meshplay is deployed in the Kubernetes cluster
	MeshplayNamespace = "meshplay"
	// MeshplayDeployment is the name of a Kubernetes manifest file required to setup Meshplay
	// check https://github.com/meshplay/meshplay/tree/master/install/deployment_yamls/k8s
	MeshplayDeployment = "meshplay-deployment.yaml"
	// MeshplayService is the name of a Kubernetes manifest file required to setup Meshplay
	// check https://github.com/meshplay/meshplay/tree/master/install/deployment_yamls/k8s
	MeshplayService = "meshplay-service.yaml"
	//MeshplayOperator is the file for default Meshplay operator
	//check https://github.com/khulnasoft/meshplay-operator/blob/master/config/manifests/default.yaml
	MeshplayOperator = "default.yaml"
	//MeshplayOperatorBroker is the file for the Meshplay broker
	//check https://github.com/khulnasoft/meshplay-operator/blob/master/config/samples/meshplay_v1alpha1_broker.yaml
	MeshplayOperatorBroker = "meshplay_v1alpha1_broker.yaml"
	//MeshplayOperatorMeshsync is the file for the Meshplay Meshsync Operator
	//check https://github.com/khulnasoft/meshplay-operator/blob/master/config/samples/meshplay_v1alpha1_meshsync.yaml
	MeshplayOperatorMeshsync = "meshplay_v1alpha1_meshsync.yaml"
	// ServiceAccount is the name of a Kubernetes manifest file required to setup Meshplay
	// check https://github.com/meshplay/meshplay/tree/master/install/deployment_yamls/k8s
	ServiceAccount = "service-account.yaml"
	// To upload with param name
	ParamName = "k8sfile"
	// kubeconfig file name
	KubeConfigYaml = "kubeconfig.yaml"
	// ViperCompose is an instance of viper for docker-compose
	ViperCompose = viper.New()
	// ViperMeshconfig is an instance of viper for the meshconfig file
	ViperMeshconfig = viper.New()
	// SilentFlag skips waiting for user input and proceeds with default options
	SilentFlag bool
	// PlatformFlag sets the platform for the initial config file
	PlatformFlag string
	// Paths to kubeconfig files
	ConfigPath string
	KubeConfig string
	// KeepNamespace indicates if the namespace should be kept when Meshplay is uninstalled
	KeepNamespace bool
	// TokenFlag sets token location passed by user with --token
	TokenFlag = "Not Set"
	// global logger variable
	Log logger.Handler
	// Color for the whiteboard printer
	whiteBoardPrinter = color.New(color.FgHiBlack, color.BgWhite, color.Bold)
	//global logger error variable
	LogError logger.Handler
)

var CfgFile string

// TODO: add "meshplay-perf" as a component

// ListOfComponents returns the list of components available
var ListOfComponents = []string{}

// TemplateContext is the template context provided when creating a config file
var TemplateContext = config.Context{
	Endpoint:   EndpointProtocol + "://localhost:9081",
	Token:      "default",
	Platform:   "kubernetes",
	Components: ListOfComponents,
	Channel:    "stable",
	Version:    "latest",
	Provider:   "Meshplay",
}

var Services = map[string]Service{
	"meshplay": {
		Image:  "layer5/meshplay:stable-latest",
		Labels: []string{"com.centurylinklabs.watchtower.enable=true"},
		Environment: []string{
			"PROVIDER_BASE_URLS=https://meshplay.layer5.io",
			"ADAPTER_URLS=meshplay-istio:10000 meshplay-linkerd:10001 meshplay-consul:10002 meshplay-nsm:10004 meshplay-app-mesh:10005 meshplay-kuma:10007 meshplay-osm:10009 meshplay-traefik-mesh:10006 meshplay-nginx-sm:10010 meshplay-cilium:10012",
			"EVENT=meshplayLocal",
			"PORT=9081",
		},
		Volumes: []string{"$HOME/.kube:/home/appuser/.kube:ro", "$HOME/.minikube:$HOME/.minikube:ro"},
		Ports:   []string{"9081:9081"},
	},
	"meshplay-istio": {
		Image:  "layer5/meshplay-istio:stable-latest",
		Labels: []string{"com.centurylinklabs.watchtower.enable=true"},
		Ports:  []string{"10000:10000"},
	},
	"meshplay-linkerd": {
		Image:  "layer5/meshplay-linkerd:stable-latest",
		Labels: []string{"com.centurylinklabs.watchtower.enable=true"},
		Ports:  []string{"10001:10001"},
	},
	"meshplay-consul": {
		Image:  "layer5/meshplay-consul:stable-latest",
		Labels: []string{"com.centurylinklabs.watchtower.enable=true"},
		Ports:  []string{"10002:10002"},
	},
	"meshplay-nsm": {
		Image:  "layer5/meshplay-nsm:stable-latest",
		Labels: []string{"com.centurylinklabs.watchtower.enable=true"},
		Ports:  []string{"10004:10004"},
	},
	"meshplay-app-mesh": {
		Image:  "layer5/meshplay-app-mesh:stable-latest",
		Labels: []string{"com.centurylinklabs.watchtower.enable=true"},
		Ports:  []string{"10005:10005"},
	},
	"meshplay-traefik-mesh": {
		Image:  "layer5/meshplay-traefik-mesh:stable-latest",
		Labels: []string{"com.centurylinklabs.watchtower.enable=true"},
		Ports:  []string{"10006:10006"},
	},
	"meshplay-kuma": {
		Image:  "layer5/meshplay-kuma:stable-latest",
		Labels: []string{"com.centurylinklabs.watchtower.enable=true"},
		Ports:  []string{"10007:10007"},
	},
	"meshplay-osm": {
		Image:  "layer5/meshplay-osm:stable-latest",
		Labels: []string{"com.centurylinklabs.watchtower.enable=true"},
		Ports:  []string{"10009:10009"},
	},
	"meshplay-nginx-sm": {
		Image:  "layer5/meshplay-nginx-sm:stable-latest",
		Labels: []string{"com.centurylinklabs.watchtower.enable=true"},
		Ports:  []string{"10010:10010"},
	},
	"meshplay-cilium": {
		Image:  "layer5/meshplay-cilium:stable-latest",
		Labels: []string{"com.centurylinklabs.watchtower.enable=true"},
		Ports:  []string{"10012:10012"},
	},
	"watchtower": {
		Image:  "containrrr/watchtower",
		Labels: []string{"com.centurylinklabs.watchtower.enable=true"},
	},
}

// TemplateToken is the template token provided when creating a config file
var TemplateToken = config.Token{
	Name:     "default",
	Location: AuthConfigFile,
}

func BackupConfigFile(cfgFile string) {
	dir, file := filepath.Split(cfgFile)
	extension := filepath.Ext(file)
	bakLocation := filepath.Join(dir, file[:len(file)-len(extension)]+".bak.yaml")
	err := os.Rename(cfgFile, bakLocation)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Create(cfgFile)
	if err != nil {
		log.Fatal(err)
	}
}

const tokenName = "token"
const providerName = "meshplay-provider"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// StringWithCharset generates a random string with a given length
func StringWithCharset(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	// + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// SafeClose is a helper function help to close the io
func SafeClose(co io.Closer) {
	if cerr := co.Close(); cerr != nil {
		log.Error(cerr)
	}
}

func prereq() ([]byte, []byte, error) {
	ostype, err := exec.Command("uname", "-s").Output()
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not find os type")
	}

	osarch, err := exec.Command("uname", "-m").Output()
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not find os arch type")
	}

	return ostype, osarch, nil
}

// SetFileLocation to set absolute path
func SetFileLocation() error {
	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get users home directory")
	}
	MeshplayFolder = filepath.Join(home, MeshplayFolder)
	DockerComposeFile = filepath.Join(MeshplayFolder, DockerComposeFile)
	AuthConfigFile = filepath.Join(MeshplayFolder, AuthConfigFile)
	DefaultConfigPath = filepath.Join(MeshplayFolder, DefaultConfigPath)
	return nil
}

// NavigateToBroswer naviagtes to the endpoint displaying Meshplay UI in the broswer.
func NavigateToBrowser(endpoint string) error {
	err := browser.OpenURL(endpoint)
	return err
}

// UploadFileWithParams returns a request configured to upload files with other values
func UploadFileWithParams(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if err = file.Close(); err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	_, err = part.Write(fileContents)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())
	return request, nil
}

// IsValidSubcommand checks if the passed subcommand is supported by the parent command
func IsValidSubcommand(available []*cobra.Command, sub string) bool {
	for _, s := range available {
		if sub == s.CalledAs() {
			return true
		}
	}
	return false
}

// ContentTypeIsHTML Checks if the response is an HTML resposnse
func ContentTypeIsHTML(resp *http.Response) bool {
	ctString := strings.Split(resp.Header.Get("Content-Type"), ";")
	if len(ctString) < 1 {
		return false
	}
	if ctString[0] == "text/html" {
		return true
	}
	return false
}

// UpdateMeshplayContainers runs the update command for meshplay client
func UpdateMeshplayContainers() error {
	log.Info("Updating Meshplay now...")

	start := exec.Command("docker-compose", "-f", DockerComposeFile, "pull")
	start.Stdout = os.Stdout
	start.Stderr = os.Stderr
	if err := start.Run(); err != nil {
		return errors.Wrap(err, SystemError("failed to start meshplay"))
	}
	return nil
}

// AskForConfirmation asks the user for confirmation. A user must type in "yes" or "no" and then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as confirmations. If the input is not recognized, it will ask again. The function does not return until it gets a valid response from the user.
func AskForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]? ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

// CreateConfigFile creates config file in Meshplay Folder
func CreateConfigFile() error {
	if _, err := os.Stat(DefaultConfigPath); os.IsNotExist(err) {
		_, err := os.Create(DefaultConfigPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// ValidateURL validates url provided for meshplay backend to meshplayctl context
func ValidateURL(URL string) error {
	ParsedURL, err := url.ParseRequestURI(URL)
	if err != nil {
		return err
	}
	if ParsedURL.Scheme != "http" && ParsedURL.Scheme != "https" {
		return fmt.Errorf("%s is not a supported protocol", ParsedURL.Scheme)
	}
	return nil
}

// TruncateID shortens an id to 8 characters
func TruncateID(id string) string {
	ShortenedID := id[0:8]
	return ShortenedID
}

// PrintToTable prints the given data into a table format
func PrintToTable(header []string, data [][]string) {
	// The tables are formatted to look similar to how it looks in say `kubectl get deployments`
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header) // The header of the table
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(data) // The data in the table
	table.Render()         // Render the table
}

// PrintToTableWithFooter prints the given data into a table format but with a footer
func PrintToTableWithFooter(header []string, data [][]string, footer []string) {
	// The tables are formatted to look similar to how it looks in say `kubectl get deployments`
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header) // The header of the table
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(data) // The data in the table
	table.SetFooter(footer)
	table.Render() // Render the table
}

// ClearLine clears the last line from output
func ClearLine() {
	clearCmd := exec.Command("clear") // for UNIX-like systems
	if runtime.GOOS == "windows" {
		clearCmd = exec.Command("cmd", "/c", "cls") // for Windows
	}
	clearCmd.Stdout = os.Stdout
	err := clearCmd.Run()
	if err != nil {
		Log.Error(ErrClearLine(err))
		return
	}
}

// StringContainedInSlice returns the index in which a string is a substring in a list of strings
func StringContainedInSlice(str string, slice []string) int {
	for index, ele := range slice {
		// Return index even if only a part of the string is present
		if strings.Contains(ele, str) {
			return index
		}
	}
	return -1
}

// StringInSlice checks if a string is present in a slice
func StringInSlice(str string, slice []string) bool {
	for _, ele := range slice {
		if ele == str {
			return true
		}
	}
	return false
}

// GetID returns a array of IDs from meshplay server endpoint /api/{configurations}
func GetID(meshplayServerUrl, configuration string) ([]string, error) {
	url := meshplayServerUrl + "/api/" + configuration + "?page_size=10000"
	configType := configuration + "s"
	var idList []string
	req, err := NewRequest("GET", url, nil)
	if err != nil {
		return idList, err
	}

	res, err := MakeRequest(req)
	if err != nil {
		return idList, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return idList, ErrReadResponseBody(err)
	}
	var dat map[string]interface{}
	if err = encoding.Unmarshal(body, &dat); err != nil {
		return idList, ErrUnmarshal(errors.Wrap(err, "failed to unmarshal response body"))
	}
	if dat == nil {
		return idList, ErrNotFound(errors.New("no data found"))
	}
	if dat[configType] == nil {
		return idList, ErrNotFound(errors.New("no results found"))
	}
	for _, config := range dat[configType].([]interface{}) {
		idList = append(idList, config.(map[string]interface{})["id"].(string))
	}
	return idList, nil
}

// GetName returns a of name:id from meshplay server endpoint /api/{configurations}
func GetName(meshplayServerUrl, configuration string) (map[string]string, error) {
	url := meshplayServerUrl + "/api/" + configuration + "?page_size=10000"
	configType := configuration + "s"
	nameIdMap := make(map[string]string)
	req, err := NewRequest("GET", url, nil)
	if err != nil {
		return nameIdMap, err
	}

	res, err := MakeRequest(req)
	if err != nil {
		return nameIdMap, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nameIdMap, ErrReadResponseBody(err)
	}
	var dat map[string]interface{}
	if err = encoding.Unmarshal(body, &dat); err != nil {
		return nameIdMap, ErrUnmarshal(errors.Wrap(err, "failed to unmarshal response body"))
	}
	if dat == nil {
		return nameIdMap, ErrNotFound(errors.New("no data found"))
	}
	if dat[configType] == nil {
		return nameIdMap, ErrNotFound(errors.New("no results found"))
	}
	for _, config := range dat[configType].([]interface{}) {
		nameIdMap[config.(map[string]interface{})["name"].(string)] = config.(map[string]interface{})["id"].(string)
	}
	return nameIdMap, nil
}

// Delete configuration from meshplay server endpoint /api/{configurations}/{id}
func DeleteConfiguration(meshplayServerUrl, id, configuration string) error {
	url := meshplayServerUrl + "/api/" + configuration + "/" + id
	req, err := NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	_, err = MakeRequest(req)
	if err != nil {
		return err
	}
	return nil
}

// ValidId - Check if args is a valid ID or a valid ID prefix and returns the full ID
func ValidId(meshplayServerUrl, args string, configuration string) (string, bool, error) {
	isID := false
	configID, err := GetID(meshplayServerUrl, configuration)
	if err == nil {
		for _, id := range configID {
			if strings.HasPrefix(id, args) {
				args = id
			}
		}
	} else {
		return "", false, err
	}
	isID, err = regexp.MatchString("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$", args)
	if err != nil {
		return "", false, ErrInvalidNameOrID(err)
	}
	return args, isID, nil
}

// ValidId - Check if args is a valid name or a valid name prefix and returns the full name and ID
func ValidName(meshplayServerUrl, args string, configuration string) (string, string, bool, error) {
	isName := false
	nameIdMap, err := GetName(meshplayServerUrl, configuration)

	if err != nil {
		return "", "", false, err
	}

	fullName := ""
	ID := ""

	for name := range nameIdMap {
		if strings.HasPrefix(name, args) {
			fullName = name
			ID = nameIdMap[name]
			isName = true
		}
	}

	return fullName, ID, isName, nil
}

// AskForInput asks the user for an input and checks if it is in the available values
func AskForInput(prompt string, allowed []string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s %s: ", prompt, allowed)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if StringInSlice(response, allowed) {
			return response
		}
		log.Fatalf("Invalid respose %s. Allowed responses %s", response, allowed)
	}
}

// ParseURLGithub checks URL and returns raw repo, path, error
func ParseURLGithub(URL string) (string, string, error) {
	// GitHub URL:
	// - https://github.com/meshplay/meshplay/blob/master/.goreleaser.yml
	// - https://raw.githubusercontent.com/layer5io/meshplay/master/.goreleaser.yml
	parsedURL, err := url.Parse(URL)
	if err != nil {
		return "", "", ErrParsingUrl(fmt.Errorf("failed to retrieve file from URL: %s", URL))
	}
	host := parsedURL.Host
	path := parsedURL.Path
	path = strings.Replace(path, "/blob/", "/", 1)
	paths := strings.Split(path, "/")
	if host == "github.com" {
		if len(paths) < 5 {
			return "", "", ErrParsingUrl(fmt.Errorf("failed to retrieve file from URL: %s", URL))
		}
		resURL := "https://" + host + strings.Join(paths[:4], "/")
		return resURL, strings.Join(paths[4:], "/"), nil
	} else if host == "raw.githubusercontent.com" {
		if len(paths) < 5 {
			return "", "", ErrParsingUrl(fmt.Errorf("failed to retrieve file from URL: %s", URL))
		}
		resURL := "https://" + "raw.githubusercontent.com" + path
		return resURL, "", nil
	}
	return URL, "", ErrParsingUrl(errors.New("only github urls are supported"))
}

// PrintToTableInStringFormat prints the given data into a table format but return as a string
func PrintToTableInStringFormat(header []string, data [][]string) string {
	// The tables are formatted to look similar to how it looks in say `kubectl get deployments`
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader(header) // The header of the table
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(data) // The data in the table
	table.Render()         // Render the table

	return tableString.String()
}

// Indicate an ongoing Process at a given time on CLI
func CreateDefaultSpinner(suffix string, finalMsg string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)

	s.Suffix = " " + suffix
	s.FinalMSG = finalMsg + "\n"
	return s
}

// Get Meshplay Session Data/Details (Adapters)
func GetSessionData(mctlCfg *config.MeshplayCtlConfig) (*models.Preference, error) {
	path := mctlCfg.GetBaseMeshplayURL() + "/api/system/sync"
	method := "GET"
	client := &http.Client{}
	req, err := NewRequest(method, path, nil)
	if err != nil {
		return nil, ErrCreatingRequest(err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, ErrRequestResponse(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, ErrReadResponseBody(err)
	}

	prefs := &models.Preference{}
	err = encoding.Unmarshal(body, prefs)
	if err != nil {
		return nil, errors.New("Failed to process JSON data. Please sign into Meshplay")
	}

	return prefs, nil
}

// ContainsStringPrefix takes a string slice and a string and returns true if it is present
func ContainsStringPrefix(arr []string, str string) bool {
	for _, el := range arr {
		if strings.HasPrefix(el, str) {
			return true
		}
	}

	return false
}

// TransformYAML takes in:
//
//	yamlByt - YAML Byte slice that needs to be modified
//	transform - function that will be executed on that value, the returned value will replace the current value
//	keys - takes in a series of keys which are supposed to be nested, numbers can also be passed to access an array
func TransformYAML(yamlByt []byte, transform func(interface{}) (interface{}, error), keys ...string) ([]byte, error) {
	var data map[string]interface{}

	err := yaml.Unmarshal(yamlByt, &data)
	if err != nil {
		return nil, err
	}

	data = RecursiveCastMapStringInterfaceToMapStringInterface(data)

	val, ok := MapGet(data, keys...)
	if !ok {
		return nil, fmt.Errorf("invalid path")
	}

	transformed, err := transform(val)
	if err != nil {
		return nil, err
	}

	MapSet(data, transformed, keys...)

	return yaml.Marshal(data)
}

// MapGet takes in the map keys - each key goes one level deeper in the map
func MapGet(mp map[string]interface{}, key ...string) (interface{}, bool) {
	if mp == nil {
		return nil, false
	}

	if len(key) == 0 {
		return mp, true
	}

	if len(key) == 1 {
		val, ok := mp[key[0]]
		return val, ok
	}

	val, ok := mp[key[0]]
	if !ok {
		return mp, false
	}

	switch v := val.(type) {
	case map[string]interface{}:
		return MapGet(v, key[1:]...)
	case []interface{}:
		// Check if we can find key in the nested structure
		if len(key) < 2 {
			return mp, false
		}

		// Check if the key[1] is of type uint, if it is then
		keyNum, err := strconv.Atoi(key[1])
		if err != nil {
			return mp, false
		}

		if keyNum >= len(v) {
			return mp, false
		}

		valMapM, ok := v[keyNum].(map[string]interface{})
		if !ok {
			return mp, false
		}

		return MapGet(valMapM, key[2:]...)
	case []map[string]interface{}:
		// Check if we can find key in the nested structure
		if len(key) < 2 {
			return mp, false
		}

		// Check if the key[1] is of type uint, if it is then
		keyNum, err := strconv.Atoi(key[1])
		if err != nil {
			return mp, false
		}

		if keyNum >= len(v) {
			return mp, false
		}

		return MapGet(v[keyNum], key[2:]...)
	}

	return mp, true
}

// MapSet takes in the map that needs to be manipulated, the value that needs to
// be assgined to be assigned and the key - each key goes one level deeper in the map
func MapSet(mp map[string]interface{}, value interface{}, key ...string) {
	var _mapSet func(map[string]interface{}, interface{}, ...string) map[string]interface{}

	_mapSet = func(mp map[string]interface{}, value interface{}, key ...string) map[string]interface{} {
		if mp == nil {
			return nil
		}

		if len(key) == 0 {
			return mp
		}

		if len(key) == 1 {
			mp[key[0]] = value
			return mp
		}

		val, ok := mp[key[0]]
		if !ok {
			return mp
		}

		switch v := val.(type) {
		case map[string]interface{}:
			mp[key[0]] = _mapSet(v, value, key[1:]...)
			return mp
		case []interface{}:
			// Check if we can find key in the nested structure
			if len(key) < 2 {
				return mp
			}

			// Check if the key[1] is of type uint, if it is then
			keyNum, err := strconv.Atoi(key[1])
			if err != nil {
				return mp
			}

			if keyNum >= len(v) {
				return mp
			}

			valMapM, ok := v[keyNum].(map[string]interface{})
			if !ok {
				return mp
			}

			v[keyNum] = _mapSet(valMapM, value, key[2:]...)

			mp[key[0]] = v

			return mp
		case []map[string]interface{}:
			// Check if we can find key in the nested structure
			if len(key) < 2 {
				return mp
			}

			// Check if the key[1] is of type uint, if it is then
			keyNum, err := strconv.Atoi(key[1])
			if err != nil {
				return mp
			}

			if keyNum >= len(v) {
				return mp
			}

			v[keyNum] = _mapSet(v[keyNum], value, key[2:]...)

			mp[key[0]] = v

			return mp
		}

		return mp
	}

	_mapSet(mp, value, key...)
}

// RecursiveCastMapStringInterfaceToMapStringInterface will convert a
// map[string]interface{} recursively => map[string]interface{}
func RecursiveCastMapStringInterfaceToMapStringInterface(in map[string]interface{}) map[string]interface{} {
	res := ConvertMapInterfaceMapString(in)
	out, ok := res.(map[string]interface{})
	if !ok {
		fmt.Println("failed to cast")
	}

	return out
}

// ConvertMapInterfaceMapString converts map[interface{}]interface{} => map[string]interface{}
//
// It will also convert []interface{} => []string
func ConvertMapInterfaceMapString(v interface{}) interface{} {
	switch x := v.(type) {
	case map[interface{}]interface{}:
		m := map[string]interface{}{}
		for k, v2 := range x {
			switch k2 := k.(type) {
			case string:
				m[k2] = ConvertMapInterfaceMapString(v2)
			default:
				m[fmt.Sprint(k)] = ConvertMapInterfaceMapString(v2)
			}
		}
		v = m

	case []interface{}:
		for i, v2 := range x {
			x[i] = ConvertMapInterfaceMapString(v2)
		}

	case map[string]interface{}:
		for k, v2 := range x {
			x[k] = ConvertMapInterfaceMapString(v2)
		}
	}

	return v
}

// SetOverrideValues returns the value overrides based on current context to install/upgrade helm chart
func SetOverrideValues(ctx *config.Context, meshplayImageVersion, callbackURL, providerURL string) map[string]interface{} {
	// first initialize all the components' "enabled" field to false
	// this matches to the components listed in install/kubernetes/helm/meshplay/values.yaml
	valueOverrides := map[string]interface{}{
		"meshplay-istio": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-cilium": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-linkerd": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-consul": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-kuma": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-nsm": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-nginx-sm": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-traefik-mesh": map[string]interface{}{
			"enabled": false,
		},
		"meshplay-app-mesh": map[string]interface{}{
			"enabled": false,
		},
	}

	// set the "enabled" field to true only for the components listed in the context
	for _, component := range ctx.GetComponents() {
		if _, ok := valueOverrides[component]; ok {
			valueOverrides[component] = map[string]interface{}{
				"enabled": true,
			}
		}
	}

	// set the meshplay image version
	valueOverrides["image"] = map[string]interface{}{
		"tag": ctx.GetChannel() + "-" + meshplayImageVersion,
	}

	// set the provider
	if ctx.GetProvider() != "" {
		valueOverrides["env"] = map[string]interface{}{
			constants.ProviderENV: ctx.GetProvider(),
		}
	}

	if callbackURL != "" {
		valueOverrides["env"] = map[string]interface{}{
			constants.CallbackURLENV: callbackURL,
		}
	}

	if providerURL != "" {
		valueOverrides["env"] = map[string]interface{}{
			constants.ProviderURLsENV: providerURL,
		}
	}

	// disable the operator
	if ctx.GetOperatorStatus() == "disabled" {
		if _, ok := valueOverrides["env"]; !ok {
			valueOverrides["env"] = map[string]interface{}{}
		}
		envOverrides := valueOverrides["env"].(map[string]interface{})
		envOverrides["DISABLE_OPERATOR"] = "'true'"
	}

	return valueOverrides
}

// CheckFileExists checks if the given file exists in system or not
func CheckFileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, fmt.Errorf("%s does not exist", name)
	}
	return false, errors.Wrap(err, fmt.Sprintf("Failed to read/fetch the file %s", name))
}

func Contains(key string, col []string) int {
	for i, n := range col {
		if n == key {
			return i
		}
	}
	return -1
}

// HandlePagination handles interactive pagination and prints the content in the terminal.
// It takes the page size, data to paginate, header for the data table, and an optional footer.
// If no footer is provided, it will be omitted.
// Pagination allows users to navigate through the data using Enter or ↓ to continue,
// Esc or Ctrl+C (Ctrl+Cmd for OS users) to exit.
func HandlePagination(pageSize int, component string, data [][]string, header []string, footer ...[]string) error {

	startIndex := 0
	endIndex := min(len(data), startIndex+pageSize)
	for {
		// Clear the entire terminal screen
		ClearLine()

		// Print number of filter files and current page number
		whiteBoardPrinter.Print("Total number of ", component, ":", len(data))
		fmt.Println()
		whiteBoardPrinter.Print("Page: ", startIndex/pageSize+1)
		fmt.Println()

		whiteBoardPrinter.Println("Press Enter or ↓ to continue. Press Esc or Ctrl+C to exit.")

		if len(footer) > 0 {
			PrintToTableWithFooter(header, data[startIndex:endIndex], footer[0])
		} else {
			PrintToTable(header, data[startIndex:endIndex])
		}
		keysEvents, err := keyboard.GetKeys(10)
		if err != nil {
			return err
		}

		defer func() {
			_ = keyboard.Close()
		}()

		event := <-keysEvents
		if event.Err != nil {
			Log.Error(fmt.Errorf("unable to capture keyboard events"))
			break
		}

		if event.Key == keyboard.KeyEsc || event.Key == keyboard.KeyCtrlC {
			break
		}

		if event.Key == keyboard.KeyEnter || event.Key == keyboard.KeyArrowDown {
			startIndex += pageSize
			endIndex = min(len(data), startIndex+pageSize)
		}

		if startIndex >= len(data) {
			break
		}
	}
	return nil
}

func FindInSlice(key string, items []string) (int, bool) {
	for idx, item := range items {
		if item == key {
			return idx, true
		}
	}
	return -1, false
}

func DisplayCount(component string, count int64) {
	whiteBoardPrinter.Println("Total number of ", component, ":", count)
}

func GetPageQueryParameter(cmd *cobra.Command, page int) string {
	if !cmd.Flags().Changed("page") {
		return "pagesize=all"
	}
	return fmt.Sprintf("page=%d", page)
}
