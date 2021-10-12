package build

import (
	"context"

	"github.com/wawandco/ox/plugins/core"
)

// Builder interface allows to set the build steps to be run.
type Builder interface {
	core.Plugin
	Build(context.Context, string, []string) error
}
