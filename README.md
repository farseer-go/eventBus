# eventBus 事件总线
> 包：`"github.com/farseer-go/eventBus"`
> 
> 模块：`eventBus.Module`

- `Document`
    - [English](https://farseer-go.gitee.io/en-us/)
    - [中文](https://farseer-go.gitee.io/)
    - [English](https://farseer-go.github.io/doc/en-us/)
- Source
    - [github](https://github.com/farseer-go/fs)


![](https://img.shields.io/github/stars/farseer-go?style=social)
![](https://img.shields.io/github/license/farseer-go/eventBus)
![](https://img.shields.io/github/go-mod/go-version/farseer-go/eventBus)
![](https://img.shields.io/github/v/release/farseer-go/eventBus)
[![codecov](https://img.shields.io/codecov/c/github/farseer-go/eventBus)](https://codecov.io/gh/farseer-go/eventBus)
![](https://img.shields.io/github/languages/code-size/farseer-go/eventBus)
![](https://img.shields.io/github/directory-file-count/farseer-go/eventBus)
![](https://goreportcard.com/badge/github.com/farseer-go/eventBus)

## 概述
以事件驱动的方式来解耦业务逻辑，在`DDD`中，事件总线是必然用到的技术。

当两个业务模块相互之间有业务关联，但又不希望在代码结构上直接依赖。

则可以使用事件驱动的方式来解耦相互之间的依赖。

## 1、发布事件
本着farseer-go极简、优雅风格，使用eventBus组件也是非常简单的：

_函数定义_
```go
// 发布事件（同步、阻塞）
func PublishEvent(eventName string, message any)

// 发布事件（异步）
func PublishEventAsync(eventName string, message any)
```
- `eventName`：事件名称
- `message`：事件消息

_演示：_
```go
type newUser struct {
    UserName string
}

func main() {
    fs.Initialize[eventBus.Module]("queue生产消息演示")

    // 同步（阻塞）
    eventBus.PublishEvent("new_user_event", newUser{UserName: "steden"})

    // or 异步（非阻塞）
    eventBus.PublishEventAsync("new_user_event", newUser{UserName: "steden"})
}
```

## 2、订阅事件
_函数定义_
```go
// 订阅
func Subscribe(eventName string, fn consumerFunc)
// 回调函数
type consumerFunc func(message any, ea EventArgs)
```
- `eventName`：事件名称
- `fn`：事件回调函数
- `message`：事件消息
- `ea`：事件参数

_演示：_
```go
type newUser struct {
    UserName string
}

func main() {
    fs.Initialize[eventBus.Module]("queue生产消息演示")

    eventBus.Subscribe("new_user_event", func (message any, ea EventArgs) {
        user := message.(NewUser)
        // do.....
    })
}
```