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

	traceManager := container.Resolve[trace.IManager]()
	// 这里上下文有可能会切换，所以退出程序时，要重新设置回上下文
	if traceContext, exists := traceManager.GetTraceContext(); exists {
		defer func() {
			trace.CurTraceContext.Set(traceContext)
		}()
	}

	// 事件发布链路
	var err error
	traceDetail := traceManager.TraceEventPublish(eventName)
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
		traceContext := traceManager.EntryEventConsumer(server, eventName, s.subscribeName)
		exception.Try(func() {
			s.consumerFunc(message, eventArgs)
		}).CatchException(func(exp any) {
			if traceContext.IsIgnore() { // 如果忽略了链路,则要在这里打印错误日志
				flog.Errorf("%s,%s 异常: %v", server, eventName, exp)
			}
		})
		traceManager.Push(traceContext, nil)
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
			// 创建一个事件消费入口
			traceContext := container.Resolve[trace.IManager]().EntryEventConsumer(server, eventName, s.subscribeName)
			exception.Try(func() {
				s.consumerFunc(message, eventArgs)
			}).CatchException(func(exp any) {
				if traceContext.IsIgnore() { // 如果忽略了链路,则要在这里打印错误日志
					flog.Errorf("%s,%s 异常: %v", server, eventName, exp)
				}
			})
			container.Resolve[trace.IManager]().Push(traceContext, nil)
		}(s)
	}
	return nil
}
