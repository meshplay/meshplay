package extensions

import (
	"github.com/khulnasoft/meshplay/server/machines"
	"github.com/khulnasoft/meshkit/broker"
	"github.com/khulnasoft/meshkit/database"
	"github.com/khulnasoft/meshkit/logger"
)

type ExtensionInput struct {
	DBHandler            *database.Handler
	MeshSyncChannel      chan struct{}
	Logger               logger.Handler
	BrokerConn           broker.Handler
	K8sConnectionTracker *machines.ConnectionToStateMachineInstanceTracker
}
