package test

import (
	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

var count1 int

func TestRegisterEvent(t *testing.T) {
	eventBus.RegisterEvent("testRegisterEvent", func(message any, ea core.EventArgs) {
		count1 = message.(int)
	})

	_ = container.Resolve[core.IEvent]("testRegisterEvent").Publish(3)

	assert.Equal(t, 3, count1)
}
