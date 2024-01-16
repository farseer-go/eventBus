package eventBus

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/core"
)

type subscribeConsumer struct {
	subscribeName string            // 消费者名称
	consumerFunc  core.ConsumerFunc // 消费者处理函数
}

// 订阅者
var subscriber collections.Dictionary[string, []subscribeConsumer]

// Subscribe 订阅事件
func Subscribe(eventName, subscribeName string, consumerFunc core.ConsumerFunc) {
	subscriber.Add(eventName, append(subscriber.GetValue(eventName), subscribeConsumer{subscribeName: subscribeName, consumerFunc: consumerFunc}))
}
