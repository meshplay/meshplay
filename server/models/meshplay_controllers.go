package models

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"github.com/layer5io/meshkit/broker/nats"
	"github.com/layer5io/meshkit/database"
	"github.com/layer5io/meshkit/logger"
	"github.com/layer5io/meshkit/models/controllers"
	"github.com/layer5io/meshkit/utils"
	meshplaykube "github.com/layer5io/meshkit/utils/kubernetes"
	"github.com/spf13/viper"
)

const (
	ChartRepo                     = "https://meshplay.github.io/meshplay.khulnasofy.com/charts"
	MeshplayServerBrokerConnection = "meshplay-server"
)

type MeshplayControllerStatusAndVersion struct {
	Status  controllers.MeshplayControllerStatus
	Version string
}

type MeshplayController int

const (
	MeshplayBroker MeshplayController = iota
	Meshsync
	MeshplayOperator
)

type MeshplayControllersHelper struct {
	// context that is being manged by a particular controllerHelper instance
	contextID string
	//  controller handlers for a particular context
	// this will be used as the source of truth
	ctxControllerHandlers map[MeshplayController]controllers.IMeshplayController

	// operator status for a particular context
	ctxOperatorStatus controllers.MeshplayControllerStatus

	// meshsync data handler for a particular context
	ctxMeshsyncDataHandler *MeshsyncDataHandler

	log          logger.Handler
	oprDepConfig controllers.OperatorDeploymentConfig
	dbHandler    *database.Handler
}

func (mch *MeshplayControllersHelper) GetControllerHandlersForEachContext() map[MeshplayController]controllers.IMeshplayController {
	return mch.ctxControllerHandlers
}

func (mch *MeshplayControllersHelper) GetMeshSyncDataHandlersForEachContext() *MeshsyncDataHandler {
	return mch.ctxMeshsyncDataHandler
}

func (mch *MeshplayControllersHelper) GetOperatorsStatusMap() controllers.MeshplayControllerStatus {
	return mch.ctxOperatorStatus
}

func NewMeshplayControllersHelper(log logger.Handler, operatorDepConfig controllers.OperatorDeploymentConfig, dbHandler *database.Handler) *MeshplayControllersHelper {
	return &MeshplayControllersHelper{
		ctxControllerHandlers: make(map[MeshplayController]controllers.IMeshplayController),
		log:                   log,
		oprDepConfig:          operatorDepConfig,
		ctxOperatorStatus:     controllers.Unknown,
		// The nil check is performed for the ctxMeshsyncDataHandler and if it is nil, then a new dataHandler for the context is assigned.
		// The presence of a handler for a context in a map indicate that the meshsync data for that context is properly being handled.
		// Resetting this value results in again subscribing to the Broker.
		ctxMeshsyncDataHandler: nil,
		dbHandler:              dbHandler,
	}
}

// initializes Meshsync data handler for the contexts for whom it has not been
// initialized yet. Apart from updating the map, it also runs the handler after
// updating the map. The presence of a handler for a context in a map indicate that
// the meshsync data for that context is properly being handled
func (mch *MeshplayControllersHelper) AddMeshsynDataHandlers(ctx context.Context, k8scontext K8sContext, userID, meshplayInstanceID uuid.UUID, provider Provider) *MeshplayControllersHelper {
	// only checking those contexts whose MeshplayConrollers are active
	// go func(mch *MeshplayControllersHelper) {

	ctxID := k8scontext.ID
	if mch.ctxMeshsyncDataHandler == nil {
		controllerHandlers := mch.ctxControllerHandlers

		// brokerStatus := controllerHandlers[MeshplayBroker].GetStatus()
		// do something if broker is being deployed , maybe try again after sometime
		brokerEndpoint, err := controllerHandlers[MeshplayBroker].GetPublicEndpoint()
		if brokerEndpoint == "" {
			if err != nil {
				mch.log.Warn(err)
			}
			mch.log.Info(fmt.Sprintf("Meshplay Broker unreachable for Kubernetes context (%v)", ctxID))
			return mch
		}
		brokerHandler, err := nats.New(nats.Options{
			// URLS: []string{"localhost:4222"},
			URLS:           []string{brokerEndpoint},
			ConnectionName: MeshplayServerBrokerConnection,
			Username:       "",
			Password:       "",
			ReconnectWait:  2 * time.Second,
			MaxReconnect:   60,
		})
		if err != nil {
			mch.log.Warn(err)
			mch.log.Info(fmt.Sprintf("MeshSync not configured for Kubernetes context (%v) due to '%v'", ctxID, err.Error()))
			return mch
		}
		mch.log.Info(fmt.Sprintf("Connected to Meshplay Broker (%v) for Kubernetes context (%v)", brokerEndpoint, ctxID))
		token, _ := ctx.Value(TokenCtxKey).(string)
		msDataHandler := NewMeshsyncDataHandler(brokerHandler, *mch.dbHandler, mch.log, provider, userID, uuid.FromStringOrNil(k8scontext.ConnectionID), meshplayInstanceID, token)
		err = msDataHandler.Run()
		if err != nil {
			mch.log.Warn(err)
			mch.log.Info(fmt.Sprintf("Unable to connect MeshSync for Kubernetes context (%s) due to: %s", ctxID, err.Error()))
			return mch
		}
		mch.ctxMeshsyncDataHandler = msDataHandler
		mch.log.Info(fmt.Sprintf("MeshSync connected for Kubernetes context (%s)", ctxID))

	}

	// }(mch)

	return mch
}

func (mch *MeshplayControllersHelper) RemoveMeshSyncDataHandler(ctx context.Context, contextID string) {

	mch.ctxMeshsyncDataHandler = nil
}

// attach a MeshplayController for each context if
// 1. the config is valid
// 2. if it is not already attached
func (mch *MeshplayControllersHelper) AddCtxControllerHandlers(ctx K8sContext) *MeshplayControllersHelper {
	// go func(mch *MeshplayControllersHelper) {

	// resetting this value as a specific controller handler instance does not have any significance opposed to
	// a MeshsyncDataHandler instance where it signifies whether or not a listener is attached

	cfg, _ := ctx.GenerateKubeConfig()
	client, err := meshplaykube.New(cfg)
	// means that the config is invalid
	if err != nil {
		mch.log.Error(err)
	}

	mch.ctxControllerHandlers = map[MeshplayController]controllers.IMeshplayController{
		MeshplayBroker:   controllers.NewMeshplayBrokerHandler(client),
		MeshplayOperator: controllers.NewMeshplayOperatorHandler(client, mch.oprDepConfig),
		Meshsync:        controllers.NewMeshsyncHandler(client),
	}

	// }(mch)
	return mch
}

func (mch *MeshplayControllersHelper) RemoveCtxControllerHandler(ctx context.Context, contextID string) {
	mch.ctxControllerHandlers = nil
}

// update the status of MeshplayOperator in all the contexts
// for whom MeshplayControllers are attached
// should be called after AddCtxControllerHandlers
func (mch *MeshplayControllersHelper) UpdateOperatorsStatusMap(ot *OperatorTracker) *MeshplayControllersHelper {
	// go func(mch *MeshplayControllersHelper) {

	if ot.IsUndeployed(mch.contextID) {
		mch.ctxOperatorStatus = controllers.Undeployed
	} else {
		if mch.ctxControllerHandlers != nil {
			operatorHandler, ok := mch.ctxControllerHandlers[MeshplayOperator]
			if ok {
				mch.ctxOperatorStatus = operatorHandler.GetStatus()
			}
		}
	}

	// }(mch)
	return mch
}

type OperatorTracker struct {
	ctxIDtoDeploymentStatus map[string]bool
	mx                      sync.Mutex
	DisableOperator         bool
}

func NewOperatorTracker(disabled bool) *OperatorTracker {
	return &OperatorTracker{
		ctxIDtoDeploymentStatus: make(map[string]bool),
		mx:                      sync.Mutex{},
		DisableOperator:         disabled,
	}
}

func (ot *OperatorTracker) Undeployed(ctxID string, undeployed bool) {
	if ot.DisableOperator { //no-op when operator is disabled
		return
	}
	if ot.ctxIDtoDeploymentStatus == nil {
		ot.ctxIDtoDeploymentStatus = make(map[string]bool)
	}
	ot.ctxIDtoDeploymentStatus[ctxID] = undeployed
}

func (ot *OperatorTracker) IsUndeployed(ctxID string) bool {
	if ot.DisableOperator { //Return true everytime so that operators stay in undeployed state across all contexts
		return true
	}
	if ot.ctxIDtoDeploymentStatus == nil {
		ot.ctxIDtoDeploymentStatus = make(map[string]bool)
		return false
	}
	return ot.ctxIDtoDeploymentStatus[ctxID]
}

// looks at the status of Meshplay Operator for each cluster and takes necessary action.
// it will deploy the operator only when it is in NotDeployed state
func (mch *MeshplayControllersHelper) DeployUndeployedOperators(ot *OperatorTracker) *MeshplayControllersHelper {
	if ot.DisableOperator { //Return true everytime so that operators stay in undeployed state across all contexts
		return mch
	}
	// go func(mch *MeshplayControllersHelper) {

	if mch.ctxOperatorStatus == controllers.NotDeployed {
		if mch.ctxControllerHandlers != nil {
			operatorHandler, ok := mch.ctxControllerHandlers[MeshplayOperator]
			if ok {
				err := operatorHandler.Deploy(false)

				if err != nil {
					mch.log.Error(err)
				}
			}
		}
	}

	// }(mch)

	return mch
}

func (mch *MeshplayControllersHelper) UndeployDeployedOperators(ot *OperatorTracker) *MeshplayControllersHelper {
	// go func(mch *MeshplayControllersHelper) {

	oprStatus := mch.ctxOperatorStatus

	if oprStatus != controllers.Undeployed {

		if mch.ctxControllerHandlers != nil {
			operatorHandler, ok := mch.ctxControllerHandlers[MeshplayOperator]
			if ok {
				err := operatorHandler.Undeploy()

				if err != nil {
					mch.log.Error(err)
				}
			}
		}
	}

	// }(mch)
	return mch
}

func NewOperatorDeploymentConfig(adapterTracker AdaptersTrackerInterface) controllers.OperatorDeploymentConfig {
	// get meshplay release version
	meshplayReleaseVersion := viper.GetString("BUILD")
	if meshplayReleaseVersion == "" || meshplayReleaseVersion == "Not Set" || meshplayReleaseVersion == "edge-latest" {
		_, latestRelease, err := CheckLatestVersion(meshplayReleaseVersion)
		// if unable to fetch latest release tag, meshkit helm functions handle
		// this automatically fetch the latest one
		if err != nil {
			// mch.log.Error(fmt.Errorf("Couldn't check release tag: %s. Will use latest version", err))
			meshplayReleaseVersion = ""
		} else {
			meshplayReleaseVersion = latestRelease
		}
	}

	return controllers.OperatorDeploymentConfig{
		MeshplayReleaseVersion: meshplayReleaseVersion,
		GetHelmOverrides: func(delete bool) map[string]interface{} {
			return setOverrideValues(delete, adapterTracker)
		},
		HelmChartRepo: ChartRepo,
	}
}

// checkLatestVersion takes in the current server version compares it with the target
// and returns the (isOutdated, latestVersion, error)
func CheckLatestVersion(serverVersion string) (*bool, string, error) {
	// Inform user of the latest release version
	versions, err := utils.GetLatestReleaseTagsSorted("meshplay", "meshplay")
	latestVersion := versions[len(versions)-1]
	isOutdated := false
	if err != nil {
		return nil, "", ErrCreateOperatorDeploymentConfig(err)
	}
	// Compare current running Meshplay server version to the latest available Meshplay release on GitHub.
	if latestVersion != serverVersion {
		isOutdated = true
		return &isOutdated, latestVersion, nil
	}

	return &isOutdated, latestVersion, nil
}

// setOverrideValues detects the currently insalled adapters and sets appropriate
// overrides so as to not uninstall them. It also sets override values for
// operator so that it can be enabled or disabled depending on the need
func setOverrideValues(delete bool, adapterTracker AdaptersTrackerInterface) map[string]interface{} {
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

// setOverrideValues detects the currently insalled adapters and sets appropriate
// overrides so as to not uninstall them.
func SetOverrideValuesForMeshplayDeploy(adapters []Adapter, adapter Adapter, install bool) map[string]interface{} {
	installedAdapters := make([]string, 0)
	for _, adapter := range adapters {
		if adapter.Name != "" {
			installedAdapters = append(installedAdapters, strings.Split(adapter.Location, ":")[0])
		}
	}

	overrideValues := map[string]interface{}{
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

	for _, adapter := range installedAdapters {
		if _, ok := overrideValues[adapter]; ok {
			overrideValues[adapter] = map[string]interface{}{
				"enabled": true,
			}
		}
	}

	// based on deploy/undeploy action change the status of adapter override
	if _, ok := overrideValues[strings.Split(adapter.Location, ":")[0]]; ok {
		overrideValues[strings.Split(adapter.Location, ":")[0]] = map[string]interface{}{
			"enabled": install,
		}
	}

	return overrideValues
}
