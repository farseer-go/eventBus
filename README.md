# eventBus
以事件驱动的方式来解耦业务逻辑，在`DDD`中，事件总线是必然用到的技术。

当两个业务模块相互之间有业务关联，但又不希望在代码结构上直接依赖。

则可以使用事件驱动的方式来解耦相互之间的依赖。

## What are the functions?
* eventBus（事件总线）
    * struct
        * EventArgs （事件属性）
    * func
        * PublishEvent （阻塞发布事件）
        * PublishEventAsync （异步发布事件）
        * Subscribe （订阅事件）

## Getting Started
订阅事件
```go
type NewUser struct {
    UserName string
}

eventBus.Subscribe("new_user_event", func (message any, ea EventArgs) {
    user := message.(NewUser)
    // do.....
})
```

发布事件
```go
// 同步（阻塞）
eventBus.PublishEvent("new_user_event", newUser{UserName: "steden"})

// or 异步（非阻塞）
eventBus.PublishEventAsync("new_user_event", newUser{UserName: "steden"})
```