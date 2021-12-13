package plugins

import (
	"os"

	"github.com/wawandco/ox/plugins/base"
	"github.com/wawandco/ox/plugins/tools/db"
	"github.com/wawandco/ox/plugins/tools/docker"
	"github.com/wawandco/ox/plugins/tools/envy"
	"github.com/wawandco/ox/plugins/tools/flect"
	"github.com/wawandco/ox/plugins/tools/git"
	"github.com/wawandco/ox/plugins/tools/grift"
	"github.com/wawandco/ox/plugins/tools/node"
	"github.com/wawandco/ox/plugins/tools/npm"

	"github.com/wawandco/ox/plugins/tools/ox"
	"github.com/wawandco/ox/plugins/tools/ox/action"
	"github.com/wawandco/ox/plugins/tools/ox/app"
	"github.com/wawandco/ox/plugins/tools/ox/assets"
	"github.com/wawandco/ox/plugins/tools/ox/cmd"
	"github.com/wawandco/ox/plugins/tools/ox/config"
	"github.com/wawandco/ox/plugins/tools/ox/embedded"
	"github.com/wawandco/ox/plugins/tools/ox/middleware"
	"github.com/wawandco/ox/plugins/tools/ox/model"
	"github.com/wawandco/ox/plugins/tools/ox/render"
	"github.com/wawandco/ox/plugins/tools/ox/resource"
	"github.com/wawandco/ox/plugins/tools/ox/template"
	"github.com/wawandco/ox/plugins/tools/refresh"
	"github.com/wawandco/ox/plugins/tools/soda"
	"github.com/wawandco/ox/plugins/tools/soda/fizz"
	"github.com/wawandco/ox/plugins/tools/soda/sql"
	"github.com/wawandco/ox/plugins/tools/standard"
	"github.com/wawandco/ox/plugins/tools/webpack"
	"github.com/wawandco/ox/plugins/tools/yarn"
)

// Default plugins for applications base. While ox
// has other plugins this list is the base that is used across most of
// the apps we build and maintain.
var Default = append(base.Plugins,
	&webpack.Plugin{},
	&refresh.Plugin{},
	&yarn.Plugin{},
	&npm.Plugin{},
	&envy.Developer{},
	&db.CreateCommand{},
	&db.DropCommand{},
	&db.ResetCommand{},

	// Application base commands.
	&db.Command{},
	&grift.Command{},

	// Builders
	&node.Builder{},
	&standard.Builder{},

	// Fixers
	&ox.InstallFixer{},
	&render.Fixer{},
	&embedded.Fixer{},

	// Expressions to be replaced
	&ox.ExpressionsFixer{},
	&model.Fixer{},
	&ox.ReplaceImportsFixer{},
	&standard.GoModTidyFixer{},
	&standard.GoImportsFixer{},

	// Generators
	&cmd.Generator{},
	&template.Generator{},
	&model.Generator{},
	&action.Generator{},
	&resource.Generator{},
	&grift.Generator{},
	&soda.Generator{},

	// Initializer
	&embedded.Initializer{},
	&model.Initializer{},
	&render.Initializer{},
	&refresh.Initializer{},
	&template.Initializer{},
	&flect.Initializer{},
	&docker.Initializer{},
	&action.Initializer{},
	&middleware.Initializer{},
	&cmd.Initializer{},
	&config.Initializer{},
	&docker.Initializer{},
	&app.Initializer{},
	&standard.Initializer{},
	&grift.Initializer{},
	&assets.Initializer{},
	&soda.Initializer{},
	&git.Initializer{},

	// &standard.AfterInitializer{},
	&standard.GetBuffalo{},
	&standard.ModTidy{},
	&yarn.AfterInitializer{},
	&npm.AfterInitializer{},
	&git.AfterInitializer{},

	// Testers
	&standard.Tester{},
	&envy.Tester{},

	// migrate command
	soda.NewCommand(os.DirFS("migrations")),
	&fizz.Creator{},
	&sql.Creator{},
)
