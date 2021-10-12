// base package contains the base plugins and hooks to the
// ox CLI, these plugins are the base commands of the CLI.
package base

import (
	"github.com/wawandco/ox/plugins/base/build"
	"github.com/wawandco/ox/plugins/base/dev"
	"github.com/wawandco/ox/plugins/base/fix"
	"github.com/wawandco/ox/plugins/base/generate"
	"github.com/wawandco/ox/plugins/base/help"
	"github.com/wawandco/ox/plugins/base/new"
	"github.com/wawandco/ox/plugins/base/test"
	"github.com/wawandco/ox/plugins/base/version"
	"github.com/wawandco/ox/plugins/core"
)

// Plugins that should be base to any OX CLI app, these provide the
// foundation for the app to work and its where teams would hook
// their own plugins.
var Plugins = []core.Plugin{
	&build.Command{},
	&dev.Command{},
	&test.Command{},
	&fix.Command{},
	&generate.Command{},
	&new.Command{},
	&version.Command{},
	&help.Command{},
}
