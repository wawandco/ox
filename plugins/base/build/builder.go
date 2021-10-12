package build

import (
	"context"

	plugins "github.com/wawandco/ox/plugins/core"
)

// Builder interface allows to set the build steps to be run.
type Builder interface {
	plugins.Plugin
	Build(context.Context, string, []string) error
}
