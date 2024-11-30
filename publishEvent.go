package eventBus

import (
	"fmt"
	"strconv"
	"time"

	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/exception"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/sonyflake"
	"github.com/farseer-go/fs/trace"
)

// PublishEvent 阻塞发布事件
func PublishEvent(eventName string, message any) error {
	// 首先从订阅者中找到是否存在eventName
	if !subscriber.ContainsKey(eventName) {
		return flog.Errorf("需要先通过订阅事件后，才能发布事件：%s", eventName)
	}

	// 这里上下文有可能会切换，所以退出程序时，要重新设置回上下文
	if traceContext := trace.CurTraceContext.Get(); traceContext != nil {
		defer func() {
			trace.CurTraceContext.Set(traceContext)
		}()
	}

	// 事件发布链路
	var err error
	traceDetail := container.Resolve[trace.IManager]().TraceEventPublish(eventName)
	defer func() { traceDetail.End(err) }()

	// 定义事件参数
	eventArgs := core.EventArgs{
		Id:         strconv.FormatInt(sonyflake.GenerateId(), 10),
		CreateAt:   time.Now().UnixMilli(),
		Message:    message,
		ErrorCount: 0,
		EventName:  eventName,
	}

	// 遍历订阅者，并同步执行事件消费
	server := fmt.Sprintf("本地Event/%s/%s/%v", core.AppName, core.AppIp, core.AppId)
	for _, s := range subscriber.GetValue(eventName) {
		// 创建一个事件消费入口
		eventTraceContext := container.Resolve[trace.IManager]().EntryEventConsumer(server, eventName, s.subscribeName)
		try := exception.Try(func() {
			s.consumerFunc(message, eventArgs)
		})
		try.CatchException(func(exp any) {
			err = flog.Error(exp)
		})
		container.Resolve[trace.IManager]().Push(eventTraceContext, err)
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
	eventArgs := core.EventArgs{
		Id:         strconv.FormatInt(sonyflake.GenerateId(), 10),
		CreateAt:   time.Now().UnixMilli(),
		Message:    message,
		ErrorCount: 0,
		EventName:  eventName,
	}

	// 遍历订阅者，并异步执行事件消费
	server := fmt.Sprintf("本地Event/%s/%s/%v", core.AppName, core.AppIp, core.AppId)
	for _, s := range subscriber.GetValue(eventName) {
		go func(s subscribeConsumer) {
			var err error
			// 创建一个事件消费入口
			eventTraceContext := container.Resolve[trace.IManager]().EntryEventConsumer(server, eventName, s.subscribeName)
			try := exception.Try(func() {
				s.consumerFunc(message, eventArgs)
			})
			try.CatchException(func(exp any) {
				err = flog.Error(exp)
			})
			container.Resolve[trace.IManager]().Push(eventTraceContext, err)
		}(s)
	}
	return nil
}
