package eventBus

// 订阅者
var subscriber map[string][]consumerFunc

func initSubscriber() {
	subscriber = make(map[string][]consumerFunc)
}
