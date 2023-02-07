package eventBus

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/core"
)

// 订阅者
var subscriber collections.Dictionary[string, []core.ConsumerFunc]

// Subscribe 订阅事件
func Subscribe(eventName string, fn core.ConsumerFunc) {
	subscriber.Add(eventName, append(subscriber.GetValue(eventName), fn))
}
