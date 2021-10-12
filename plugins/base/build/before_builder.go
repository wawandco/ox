package build

import (
	"context"

	plugins "github.com/wawandco/ox/plugins/core"
)

// BeforeBuilder interface allows to identify the things
// that will run before the build process has started.
type BeforeBuilder interface {
	plugins.Plugin
	RunBeforeBuild(context.Context, string, []string) error
}
