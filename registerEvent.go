package eventBus

import (
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
)

type registerEvent struct {
	eventName string
}

func (c *registerEvent) Publish(message any) {
	_ = PublishEvent(c.eventName, message)
}

// RegisterEvent 注册core.IEvent实现
func RegisterEvent(eventName string, fns ...consumerFunc) {
	// 注册仓储
	container.Register(func() core.IEvent {
		return &registerEvent{eventName: eventName}
	}, eventName)

	// 同时订阅消费
	for i := 0; i < len(fns); i++ {
		Subscribe(eventName, fns[i])
	}
}
