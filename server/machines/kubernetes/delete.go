package kubernetes

import (
	"context"
	"os"
	"time"

	"github.com/gofrs/uuid"
	"github.com/khulnasoft/meshplay/server/machines"
	"github.com/khulnasoft/meshplay/server/models"
	"github.com/khulnasoft/meshkit/logger"
	"github.com/khulnasoft/meshkit/models/events"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type DeleteAction struct{}

func (da *DeleteAction) ExecuteOnEntry(ctx context.Context, machineCtx interface{}, data interface{}) (machines.EventType, *events.Event, error) {
	return machines.NoOp, nil, nil
}

func (da *DeleteAction) Execute(ctx context.Context, machineCtx interface{}, data interface{}) (machines.EventType, *events.Event, error) {
	logLevel := viper.GetInt("LOG_LEVEL")
	if viper.GetBool("DEBUG") {
		logLevel = int(logrus.DebugLevel)
	}
	// Initialize Logger instance
	log, err := logger.New("meshplay", logger.Options{
		Format:   logger.SyslogLogFormat,
		LogLevel: logLevel,
	})
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	user, _ := ctx.Value(models.UserCtxKey).(*models.User)
	sysID, _ := ctx.Value(models.SystemIDKey).(*uuid.UUID)
	provider, _ := ctx.Value(models.ProviderCtxKey).(models.Provider)
	userUUID := uuid.FromStringOrNil(user.ID)

	eventBuilder := events.NewEvent().ActedUpon(userUUID).WithCategory("connection").WithAction("update").FromSystem(*sysID).FromUser(userUUID).WithDescription("Failed to interact with the connection.")

	machinectx, err := GetMachineCtx(machineCtx, eventBuilder)
	if err != nil {
		eventBuilder.WithMetadata(map[string]interface{}{"error": err})
		return machines.NoOp, eventBuilder.Build(), err
	}

	contextID := machinectx.K8sContext.ID

	go func() {

		machinectx.MeshplayCtrlsHelper.UpdateOperatorsStatusMap(machinectx.OperatorTracker).
			UndeployDeployedOperators(machinectx.OperatorTracker).
			RemoveCtxControllerHandler(ctx, contextID)

		machinectx.MeshplayCtrlsHelper.RemoveMeshSyncDataHandler(ctx, contextID)
	}()

	_ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()
	context.AfterFunc(_ctx, func() {
		// machinectx.MeshplayCtrlsHelper.UpdateOperatorsStatusMap(machinectx.OperatorTracker)
	})

	go models.FlushMeshSyncData(ctx, machinectx.K8sContext, provider, machinectx.EventBroadcaster, user.ID, sysID, log)

	return machines.NoOp, nil, nil
}

func (da *DeleteAction) ExecuteOnExit(ctx context.Context, machineCtx interface{}, data interface{}) (machines.EventType, *events.Event, error) {
	return machines.NoOp, nil, nil
}
