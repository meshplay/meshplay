package channels

import "github.com/khulnasoft/meshkit/broker"

var (
	BrokerPublish   = "broker-publish"
	BrokerSubscribe = "broker-subscribe"
)

type BrokerPublishPayload struct {
	Subject string
	Data    *broker.Message
}

func NewBrokerSubscribeChannel() BrokerSubscribeChannel {
	return make(chan *broker.Message)
}

type BrokerSubscribeChannel chan *broker.Message

func (ch BrokerSubscribeChannel) Stop() {
	<-ch
}

func NewBrokerPublishChannel() BrokerPublishChannel {
	return make(chan *BrokerPublishPayload)
}

type BrokerPublishChannel chan *BrokerPublishPayload

func (ch BrokerPublishChannel) Stop() {
	<-ch
}
