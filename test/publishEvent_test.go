package test

import (
	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var count int

type testEventPublish struct {
	count int
}

func TestPublishEvent(t *testing.T) {
	fs.Initialize[eventBus.Module]("unit test")

	eventBus.Subscribe("test_event_subscribe", func(message any, ea eventBus.EventArgs) {
		event := message.(testEventPublish)
		count += event.count + 1
	})

	eventBus.Subscribe("test_event_subscribe", func(message any, ea eventBus.EventArgs) {
		event := message.(testEventPublish)
		count += event.count + 2
	})

	eventBus.PublishEvent("test_event_subscribe", testEventPublish{count: 6})
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, 15, count)

	eventBus.PublishEventAsync("test_event_subscribe", testEventPublish{count: 4})
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, 26, count)
}
