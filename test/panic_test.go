package test

import (
	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanic(t *testing.T) {
	fs.Initialize[eventBus.Module]("unit test")

	assert.Error(t, eventBus.PublishEvent("testPanic", nil))

	assert.Error(t, eventBus.PublishEventAsync("testPanic", nil))
}
