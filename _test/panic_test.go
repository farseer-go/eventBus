package test

import (
	"github.com/farseer-go/eventBus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanic(t *testing.T) {
	assert.Error(t, eventBus.PublishEvent("testPanic", nil))
	assert.Error(t, eventBus.PublishEventAsync("testPanic", nil))
}
