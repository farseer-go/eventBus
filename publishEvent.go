package eventBus

import (
	"github.com/farseer-go/fs/exception"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/stopwatch"
	"math/rand"
	"strconv"
	"time"
)

// PublishEvent 阻塞发布事件
func PublishEvent(eventName string, message any) error {
	// 首先从订阅者中找到是否存在eventName
	if !subscriber.ContainsKey(eventName) {
		return flog.Errorf("需要先通过订阅事件后，才能发布事件：%s", eventName)
	}

	// 定义事件参数
	eventArgs := EventArgs{
		Id:         strconv.FormatInt(time.Now().UnixMilli(), 10) + strconv.Itoa(rand.Intn(999-100)+100),
		CreateAt:   time.Now().UnixMilli(),
		Message:    message,
		ErrorCount: 0,
	}

	// 遍历订阅者，并同步执行事件消费
	var err error
	for _, subscribeFunc := range subscriber.GetValue(eventName) {
		try := exception.Try(func() {
			sw := stopwatch.StartNew()
			subscribeFunc(message, eventArgs)
			flog.ComponentInfof("event", "%s，耗时：%s", eventName, sw.GetMillisecondsText())
		})
		try.CatchException(func(exp any) {
			err = flog.Error(exp)
		})
	}
	return err
}

// PublishEventAsync 异步发布事件
func PublishEventAsync(eventName string, message any) error {
	// 首先从订阅者中找到是否存在eventName
	if !subscriber.ContainsKey(eventName) {
		return flog.Errorf("需要先通过订阅事件后，才能发布事件：%s", eventName)
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
		go func(subscribeFunc consumerFunc) {
			try := exception.Try(func() {
				subscribeFunc(message, eventArgs)
			})
			try.CatchException(func(exp any) {
				_ = flog.Error(exp)
			})
		}(subscribeFunc)
	}
	return nil
}
