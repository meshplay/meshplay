package model

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	operatorv1alpha1 "github.com/khulnasoft/meshplay-operator/api/v1alpha1"
	operatorClient "github.com/khulnasoft/meshplay-operator/pkg/client"
	"github.com/khulnasoft/meshplay/server/models"
	brokerpkg "github.com/khulnasoft/meshkit/broker"
	"github.com/khulnasoft/meshkit/broker/nats"
	"github.com/khulnasoft/meshkit/logger"

	"github.com/khulnasoft/meshkit/models/controllers"
	"github.com/khulnasoft/meshkit/utils"
	meshplaykube "github.com/khulnasoft/meshkit/utils/kubernetes"
	kubeerror "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	Namespace                = "meshplay"
	RequestSubject           = "meshplay.meshsync.request"
	MeshsyncSubject          = "meshplay.meshsync.core"
	BrokerQueue              = "meshplay"
	MeshSyncBrokerConnection = "meshsync"
)

type Connections struct {
	Connections []connection `json:"connections"`
}

type connection struct {
	Name string `json:"name"`
}

func Initialize(client *meshplaykube.Client, delete bool, adapterTracker models.AdaptersTrackerInterface) error {
	// installOperator
	err := installUsingHelm(client, delete, adapterTracker)
	if err != nil {
		return err
	}

	return nil
}

func GetOperator(kubeclient *meshplaykube.Client) (string, string, error) {
	if kubeclient == nil || kubeclient.KubeClient == nil {
		return "", "", ErrMeshplayClientNil
	}

	dep, err := kubeclient.KubeClient.AppsV1().Deployments("meshplay").Get(context.TODO(), "meshplay-operator", metav1.GetOptions{})
	if err != nil && !kubeerror.IsNotFound(err) {
		return "", "", ErrMeshplayClient(err)
	}

	version := ""
	if err == nil {
		for _, container := range dep.Spec.Template.Spec.Containers {
			if container.Name == "manager" {
				version = strings.Split(container.Image, ":")[1]
			}
		}
	}

	return dep.ObjectMeta.Name, version, nil
}

func GetBrokerInfo(broker controllers.IMeshplayController, log logger.Handler) OperatorControllerStatus {
	brokerStatus := broker.GetStatus().String()
	monitorEndpoint, err := broker.GetEndpointForPort("monitor")
	log.Debug("broker monitor endpoint", monitorEndpoint, err)

	if brokerStatus == controllers.Connected.String() {
		brokerEndpoint, _ := broker.GetPublicEndpoint()
		brokerStatus = fmt.Sprintf("%s %s", brokerStatus, brokerEndpoint)
	}
	// change the type of IMeshplayController GetName() to to models.MeshplayController
	// and use MeshplayControllersStatusListItem instead of OperatorControllerStatus
	brokerControllerStatus := OperatorControllerStatus{
		Name:   broker.GetName(),
		Status: Status(brokerStatus),
	}

	brokerControllerStatus.Version, _ = broker.GetVersion()

	return brokerControllerStatus
}

func GetMeshSyncInfo(meshsync controllers.IMeshplayController, broker controllers.IMeshplayController, log logger.Handler) OperatorControllerStatus {
	meshsyncStatus := meshsync.GetStatus().String()

	// Debug block
	if broker != nil {
		monitorEndpoint, err := broker.GetEndpointForPort("monitor")
		log.Debug("broker monitor endpoint", monitorEndpoint, err)
	}

	// change the type of IMeshplayController GetName() to to models.MeshplayController
	// and use MeshplayControllersStatusListItem instead of OperatorControllerStatus
	if meshsyncStatus == controllers.Connected.String() && broker != nil {
		brokerEndpoint, _ := broker.GetPublicEndpoint()
		meshsyncStatus = fmt.Sprintf("%s %s", meshsyncStatus, brokerEndpoint)
	}
	version, _ := meshsync.GetVersion()
	meshsyncControllerStatus := OperatorControllerStatus{
		Name:    meshsync.GetName(),
		Version: version,
		Status:  Status(meshsyncStatus),
	}

	return meshsyncControllerStatus
}

func SubscribeToBroker(_ models.Provider, meshplayKubeClient *meshplaykube.Client, datach chan *brokerpkg.Message, brokerConn brokerpkg.Handler, ct *K8sConnectionTracker) (string, error) {
	var broker *operatorv1alpha1.Broker
	var endpoints []string
	if ct != nil {
		endpoints = ct.ListBrokerEndpoints()
	}
	meshplayclient, err := operatorClient.New(&meshplayKubeClient.RestConfig)
	if err != nil {
		if meshplayclient == nil {
			return "", ErrMeshplayClientNil
		}
		return "", ErrMeshplayClient(err)
	}

	timeout := 60
	for timeout > 0 {
		broker, err = meshplayclient.CoreV1Alpha1().Brokers(Namespace).Get(context.Background(), "meshplay-broker", metav1.GetOptions{})
		if err == nil && broker.Status.Endpoint.External != "" {
			break
		}

		timeout--
		time.Sleep(1 * time.Second)
	}

	endpoint := broker.Status.Endpoint.Internal
	if len(strings.Split(broker.Status.Endpoint.Internal, ":")) > 1 {
		port, _ := strconv.Atoi(strings.Split(broker.Status.Endpoint.Internal, ":")[1])
		if !utils.TcpCheck(&utils.HostPort{
			Address: strings.Split(broker.Status.Endpoint.Internal, ":")[0],
			Port:    int32(port),
		}, nil) {
			endpoint = broker.Status.Endpoint.External
			port, _ = strconv.Atoi(strings.Split(broker.Status.Endpoint.External, ":")[1])
			if !utils.TcpCheck(&utils.HostPort{
				Address: strings.Split(broker.Status.Endpoint.External, ":")[0],
				Port:    int32(port),
			}, nil) {
				if !utils.TcpCheck(&utils.HostPort{
					Address: "host.docker.internal",
					Port:    int32(port),
				}, nil) {
					u, _ := url.Parse(meshplayKubeClient.RestConfig.Host)
					if utils.TcpCheck(&utils.HostPort{
						Address: u.Hostname(),
						Port:    int32(port),
					}, nil) {
						endpoint = fmt.Sprintf("%s:%d", u.Hostname(), int32(port))
					}
				} else {
					endpoint = fmt.Sprintf("host.docker.internal:%d", int32(port))
				}
			}
		}
	}
	endpoints = append(endpoints, endpoint)
	// subscribing to nats
	conn, err := nats.New(nats.Options{
		URLS:           endpoints,
		ConnectionName: "meshplay",
		Username:       "",
		Password:       "",
		ReconnectWait:  2 * time.Second,
		MaxReconnect:   5,
	})
	// Hack for minikube based clusters
	if err != nil && conn == nil {
		return endpoint, err
	}
	defer func() {
		if conn == nil {
			return
		}
		available := make(map[string]bool)
		for _, server := range conn.ConnectedEndpoints() {
			available[server] = true
		}
		ct.ResetEndpoints(available)
	}()
	conn.DeepCopyInto(brokerConn)
	err = brokerConn.SubscribeWithChannel(MeshsyncSubject, BrokerQueue, datach)
	if err != nil {
		return endpoint, ErrSubscribeChannel(err)
	}

	err = brokerConn.Publish(RequestSubject, &brokerpkg.Message{
		Request: &brokerpkg.RequestObject{
			Entity: brokerpkg.ReSyncDiscoveryEntity,
		},
	})

	if err != nil {
		return endpoint, ErrPublishBroker(err)
	}

	return endpoint, nil
}
