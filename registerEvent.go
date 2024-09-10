package eventBus

import (
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
)

type registerEvent struct {
	eventName string
}

func (c *registerEvent) Publish(message any) error {
	return PublishEvent(c.eventName, message)
}

func (c *registerEvent) PublishAsync(message any) {
	_ = PublishEventAsync(c.eventName, message)
}

type registerSubscribe struct {
	eventName string
}

// RegisterEvent 注册core.IEvent实现
func RegisterEvent(eventName string) *registerSubscribe {
	// 注册仓储
	container.Register(func() core.IEvent {
		return &registerEvent{eventName: eventName}
	}, eventName)

	return &registerSubscribe{
		eventName: eventName,
	}
}

// RegisterSubscribe 注册订阅者
func (receiver *registerSubscribe) RegisterSubscribe(subscribeName string, consumerFunc core.ConsumerFunc) *registerSubscribe {
	Subscribe(receiver.eventName, subscribeName, consumerFunc)
	return receiver
}
