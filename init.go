package eventBus

import "github.com/farseer-go/collections"

// 订阅者
var subscriber collections.Dictionary[string, []consumerFunc]

func initSubscriber() {
	subscriber = collections.NewDictionary[string, []consumerFunc]()
}
