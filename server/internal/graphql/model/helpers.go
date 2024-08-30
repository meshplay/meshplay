package model

import (
	"context"
	"database/sql"
	"strings"
	"sync"

	"github.com/khulnasoft/meshplay/server/handlers"
	"github.com/khulnasoft/meshplay/server/models"
	"github.com/layer5io/meshkit/broker"
	"github.com/layer5io/meshkit/logger"
	"github.com/layer5io/meshkit/models/controllers"
	meshplaykube "github.com/layer5io/meshkit/utils/kubernetes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// to be moved elsewhere
const (
	chartRepo = "https://meshplay.github.io/meshplay.io/charts"
)

var (
	controlPlaneNamespace = map[MeshType][]string{
		MeshTypeIstio:              {"istio-system"},
		MeshTypeLinkerd:            {"linkerd-system"},
		MeshTypeCiliumServiceMesh:  {"kube-system"},
		MeshTypeConsul:             {"consul-system"},
		MeshTypeTraefikMesh:        {"traefik-system"},
		MeshTypeKuma:               {"kuma-system"},
		MeshTypeNginxServiceMesh:   {"nginx-system"},
		MeshTypeNetworkServiceMesh: {"nsm-system"},
		MeshTypeAppMesh:            {"appmesh-system"},
		//Any namespace added or appended above should also be appended on the AllMesh array
		MeshTypeAllMesh: {"istio-system", "linkerd-system", "consul-system", "traefik-system", "kuma-system", "nginx-system", "nsm-system", "appmesh-system"},
	}

	addonPortSelector = map[string]string{
		"grafana":          "service",
		"prometheus":       "http",
		"jaeger-collector": "jaeger-collector-http",
		"kiali":            "http",
		"zipkin":           "http-query",
	}
)
var (
	//TODO: Add the image orgs of other control plane pods. This change is backwards compatible and wont break anything
	controlPlaneImageOrgs = map[MeshType][]string{
		MeshTypeCiliumServiceMesh: {"cilium"},
	}
)

// installs operator
// To be depricated
func installUsingHelm(client *meshplaykube.Client, delete bool, _ models.AdaptersTrackerInterface) error {
	// retrieving meshplay's version to apply the appropriate chart
	meshplayReleaseVersion := viper.GetString("BUILD")
	if meshplayReleaseVersion == "" || meshplayReleaseVersion == "Not Set" || meshplayReleaseVersion == "edge-latest" {
		_, latestRelease, err := handlers.CheckLatestVersion(meshplayReleaseVersion)
		// if unable to fetch latest release tag, meshkit helm functions handle
		// this automatically fetch the latest one
		if err != nil {
			logrus.Errorf("Couldn't check release tag: %s. Will use latest version", err)
			meshplayReleaseVersion = ""
		} else {
			meshplayReleaseVersion = latestRelease
		}
	}
	var (
		act   = meshplaykube.INSTALL
		chart = "meshplay-operator"
	)
	if delete {
		act = meshplaykube.UNINSTALL
	}
	// a basic check to see if meshplay is installed in cluster
	// this helps decide what chart should be used for installing operator
	if viper.GetString("KUBERNETES_SERVICE_HOST") != "" {
		// act = meshplaykube.UPGRADE
		chart = "meshplay"
	}

	err := client.ApplyHelmChart(meshplaykube.ApplyHelmChartConfig{
		Namespace:   "meshplay",
		ReleaseName: "meshplay-operator",
		ChartLocation: meshplaykube.HelmChartLocation{
			Repository: chartRepo,
			Chart:      chart,
			Version:    meshplayReleaseVersion,
		},
		// CreateNamespace doesn't have any effect when the action is UNINSTALL
		CreateNamespace: true,
		Action:          act,
	})
	if err != nil {
		return ErrApplyHelmChart(err)
	}

	return nil
}

// SetOverrideValues detects the currently insalled adapters and sets appropriate
// overrides so as to not uninstall them. It also sets override values for
// operator so that it can be enabled or disabled depending on the need

// to be depricated
func SetOverrideValues(delete bool, adapterTracker models.AdaptersTrackerInterface) map[string]interface{} {
	installedAdapters := make([]string, 0)
	adapters := adapterTracker.GetAdapters(context.TODO())

	for _, adapter := range adapters {
		if adapter.Name != "" {
			installedAdapters = append(installedAdapters, strings.Split(adapter.Location, ":")[0])
		}
	}

	overrideValues := map[string]interface{}{
		"fullnameOverride": "meshplay-operator",
		"meshplay": map[string]interface{}{
			"enabled": false,
		},
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
		"meshplay-operator": map[string]interface{}{
			"enabled": true,
		},
	}

	for _, adapter := range installedAdapters {
		if _, ok := overrideValues[adapter]; ok {
			overrideValues[adapter] = map[string]interface{}{
				"enabled": true,
			}
		}
	}

	if delete {
		overrideValues["meshplay-operator"] = map[string]interface{}{
			"enabled": false,
		}
	}

	return overrideValues
}

// K8sConnectionTracker keeps track of BrokerURLs per kubernetes context
type K8sConnectionTracker struct {
	mx              sync.Mutex
	contextToBroker map[string]string //ContextID -> BrokerURL
}

func NewK8sConnctionTracker() *K8sConnectionTracker {
	return &K8sConnectionTracker{
		contextToBroker: make(map[string]string),
	}
}
func (k *K8sConnectionTracker) Set(id string, url string) {
	k.mx.Lock()
	defer k.mx.Unlock()
	k.contextToBroker[id] = url
}

// Takes a set of endpoints and discard the current endpoint if its not present in the set
func (k *K8sConnectionTracker) ResetEndpoints(available map[string]bool) {
	k.mx.Lock()
	defer k.mx.Unlock()
	c := make(map[string]string)
	for id, url := range k.contextToBroker {
		if available[url] {
			c[id] = url
		}
	}
	k.contextToBroker = c
}
func (k *K8sConnectionTracker) ListBrokerEndpoints() (a []string) {
	k.mx.Lock()
	defer k.mx.Unlock()
	for _, v := range k.contextToBroker {
		a = append(a, v)
	}
	return
}
func (k *K8sConnectionTracker) Get(id string) (url string) {
	k.mx.Lock()
	defer k.mx.Unlock()
	url = k.contextToBroker[id]
	return
}

// Takes the meshkit Logger and logs a comma separated list of currently tracked Broker Endpoints
func (k *K8sConnectionTracker) Log(l logger.Handler) {
	var e = "Connected broker endpoints : "
	k.mx.Lock()
	defer k.mx.Unlock()
	for _, v := range k.contextToBroker {
		e += v + ", "
	}
	l.Info(strings.TrimSuffix(e, ", "))
}

func GetInternalController(controller models.MeshplayController) MeshplayController {
	switch controller {
	case models.MeshplayBroker:
		return MeshplayControllerBroker
	case models.MeshplayOperator:
		return MeshplayControllerOperator
	case models.Meshsync:
		return MeshplayControllerMeshsync
	}
	return ""
}

func GetInternalControllerStatus(status controllers.MeshplayControllerStatus) MeshplayControllerStatus {
	switch status {
	case controllers.Deployed:
		return MeshplayControllerStatusDeployed

	case controllers.NotDeployed:
		return MeshplayControllerStatusNotdeployed

	case controllers.Deploying:
		return MeshplayControllerStatusDeploying

	case controllers.Unknown:
		return MeshplayControllerStatusUnkown

	case controllers.Undeployed:
		return MeshplayControllerStatusUndeployed

	case controllers.Enabled:
		return MeshplayControllerStatusEnabled

	case controllers.Running:
		return MeshplayControllerStatusRunning

	case controllers.Connected:
		return MeshplayControllerStatusConnected
	}
	return ""
}

func CheckIfBrokerEventExistsInArray(event broker.EventType, events []broker.EventType) bool {
	for _, e := range events {
		if e == event {
			return true
		}
	}
	return false
}

func GetMeshplayBrokerEventTypesFromArray(events []MeshSyncEventType) []broker.EventType {
	var brokerEvents []broker.EventType
	for _, event := range events {
		brokerEvents = append(brokerEvents, GetMeshplayBrokerEventTypes(event))
	}
	return brokerEvents
}

func GetMeshplayBrokerEventTypes(event MeshSyncEventType) broker.EventType {
	switch event {
	case MeshSyncEventTypeAdded:
		return broker.Add
	case MeshSyncEventTypeDeleted:
		return broker.Delete
	case MeshSyncEventTypeModified:
		return broker.Update
	}
	return ""
}

// SelectivelyFetchNamespaces fetches the an array of namespaces from DB based on ClusterIDs (or KubernetesServerIDs) passed as param
func SelectivelyFetchNamespaces(cids []string, provider models.Provider) ([]string, error) {
	namespaces := make([]string, 0)
	var rows *sql.Rows
	var err error
	rows, err = provider.GetGenericPersister().Raw("SELECT DISTINCT rom.name as name FROM kubernetes_resources kr LEFT JOIN kubernetes_resource_object_meta rom ON kr.id = rom.id WHERE kr.kind = 'Namespace' AND kr.cluster_id IN ?", cids).Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}

		namespaces = append(namespaces, name)
	}
	return namespaces, nil
}
