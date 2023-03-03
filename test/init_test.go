package test

import (
	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs"
)

func init() {
	fs.Initialize[eventBus.Module]("unit test")
}
