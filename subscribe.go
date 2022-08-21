package eventBus

import "github.com/farseer-go/collections"

// 订阅者
var subscriber collections.Dictionary[string, []consumerFunc]

// 订阅者的函数
type consumerFunc func(message any, ea EventArgs)

// Subscribe 订阅事件
func Subscribe(eventName string, fn consumerFunc) {
	subscriber.Add(eventName, append(subscriber.GetValue(eventName), fn))
}
