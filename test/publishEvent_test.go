package test

import (
	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs"
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
	"time"
)

var count int32

type testEventPublish struct {
	count int32
}

func TestPublishEvent(t *testing.T) {
	fs.Initialize[eventBus.Module]("unit test")

	eventBus.Subscribe("test_event_subscribe", func(message any, ea eventBus.EventArgs) {
		event := message.(testEventPublish)
		atomic.AddInt32(&count, event.count+1)
	})

	eventBus.Subscribe("test_event_subscribe", func(message any, ea eventBus.EventArgs) {
		event := message.(testEventPublish)
		atomic.AddInt32(&count, event.count+2)
	})

	eventBus.PublishEvent("test_event_subscribe", testEventPublish{count: 6})
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, int32(15), count)

	eventBus.PublishEventAsync("test_event_subscribe", testEventPublish{count: 4})
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, int32(26), atomic.LoadInt32(&count))
}
