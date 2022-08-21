package eventBus

import (
	"github.com/farseer-go/fs/modules"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var count int

type testEventPublish struct {
	count int
}

func TestPublishEvent(t *testing.T) {
	modules.StartModules(Module{})

	Subscribe("test_event_subscribe", func(message any, ea EventArgs) {
		event := message.(testEventPublish)
		count += event.count + 1
	})

	Subscribe("test_event_subscribe", func(message any, ea EventArgs) {
		event := message.(testEventPublish)
		count += event.count + 2
	})

	PublishEvent("test_event_subscribe", testEventPublish{count: 6})
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, 15, count)

	PublishEventAsync("test_event_subscribe", testEventPublish{count: 4})
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, 26, count)
}
