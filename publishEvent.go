package eventBus

import (
	"github.com/farseer-go/fs/exception"
	"github.com/farseer-go/fs/flog"
	"math/rand"
	"strconv"
	"time"
)

// PublishEvent 阻塞发布事件
func PublishEvent(eventName string, message any) {
	// 首先从订阅者中找到是否存在eventName
	if !subscriber.ContainsKey(eventName) {
		return
		//panic("未找到事件名称：" + eventName + "，需要先通过订阅事件后，才能发布事件")
	}

	// 定义事件参数
	eventArgs := EventArgs{
		Id:         strconv.FormatInt(time.Now().UnixMilli(), 10) + strconv.Itoa(rand.Intn(999-100)+100),
		CreateAt:   time.Now().UnixMilli(),
		Message:    message,
		ErrorCount: 0,
	}

	// 遍历订阅者，并异步执行事件消费
	for _, subscribeFunc := range subscriber.GetValue(eventName) {
		try := exception.Try(func() {
			subscribeFunc(message, eventArgs)
		})
		try.CatchException(func(exp any) {
			flog.Error(exp)
		})
	}
}

// PublishEventAsync 异步发布事件
func PublishEventAsync(eventName string, message any) {
	// 首先从订阅者中找到是否存在eventName
	if !subscriber.ContainsKey(eventName) {
		panic("未找到事件名称：" + eventName + "，需要先通过订阅事件后，才能发布事件")
	}

	// 定义事件参数
	eventArgs := EventArgs{
		Id:         strconv.FormatInt(time.Now().UnixMilli(), 10) + strconv.Itoa(rand.Intn(999-100)+100),
		CreateAt:   time.Now().UnixMilli(),
		Message:    message,
		ErrorCount: 0,
	}

	// 遍历订阅者，并异步执行事件消费
	for _, subscribeFunc := range subscriber.GetValue(eventName) {
		go subscribeFunc(message, eventArgs)
	}
}
