package fix

import (
	"context"

	"github.com/wawandco/ox/plugins/core"
)

// Fixer interface is created for those commands that fill fix certain
// part of our code to match versions or compile correctly. Fixers are
// a good way to transition from old versions of the tools into newer
// ones
type Fixer interface {
	core.Plugin
	Fix(context.Context, string, []string) error
}
