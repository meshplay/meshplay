package resolver

import (
	"github.com/khulnasoft/meshplay/server/internal/graphql/model"
	"github.com/khulnasoft/meshplay/server/models"
	"github.com/khulnasoft/meshkit/broker"
	"github.com/khulnasoft/meshkit/logger"
	"github.com/khulnasoft/meshkit/utils/broadcast"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Log                          logger.Handler
	BrokerConn                   broker.Handler
	MeshSyncChannelPerK8sContext map[string]chan struct{}
	Config                       *models.HandlerConfig
	Broadcast                    broadcast.Broadcaster
	// operatorChannel         chan *model.OperatorStatus
	// performanceChannel  chan *model.PerfPageResult
	brokerChannel       chan *broker.Message
	addonChannel        chan []*model.AddonList
	controlPlaneChannel chan []*model.ControlPlane
	dataPlaneChannel    chan []*model.DataPlane
}
