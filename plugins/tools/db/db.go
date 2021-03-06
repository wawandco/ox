// db packs all db operations under this top level command.
package db

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/gobuffalo/pop/v6"
	"github.com/wawandco/ox/internal/info"
	"github.com/wawandco/ox/internal/log"
	"github.com/wawandco/ox/plugins/core"
)

var _ core.Command = (*Command)(nil)
var _ core.HelpTexter = (*Command)(nil)
var _ core.PluginReceiver = (*Command)(nil)
var _ core.Subcommander = (*Command)(nil)

var ErrConnectionNotFound = errors.New("connection not found")

type Command struct {
	subcommands []core.Command
}

func (c Command) Name() string {
	return "database"
}

func (c Command) Alias() string {
	return "db"
}

func (c Command) ParentName() string {
	return ""
}

func (c Command) HelpText() string {
	return "database operation commands"
}

func (c *Command) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		log.Error("no subcommand specified, please use `db [subcommand]` to run one of the db subcommands.")
		return nil
	}

	if len(pop.Connections) == 0 {
		err := pop.LoadConfigFile()
		if err != nil {
			log.Errorf("error on db.Run: %v", err.Error())
		}
	}

	name := args[1]
	var subcommand core.Command
	for _, sub := range c.subcommands {
		if sub.Name() != name {
			continue
		}

		subcommand = sub
		break
	}

	if subcommand == nil {
		return fmt.Errorf("subcommand `%v` not found", name)
	}

	return subcommand.Run(ctx, root, args)
}

func (c *Command) Receive(pls []core.Plugin) {
	for _, plugin := range pls {
		ptool, ok := plugin.(core.Command)
		if !ok || ptool.ParentName() != c.Name() {
			continue
		}

		c.subcommands = append(c.subcommands, ptool)
	}
}

func (c *Command) Subcommands() []core.Command {
	return c.subcommands
}

func (c *Command) FindRoot() string {
	root := info.RootFolder()
	if root != "" {
		return root
	}

	root, err := os.Getwd()
	if err != nil {
		return ""
	}

	return root
}

func Plugins() []core.Plugin {
	var result []core.Plugin

	result = append(result, &Command{})
	result = append(result, &CreateCommand{})
	result = append(result, &DropCommand{})
	result = append(result, &ResetCommand{})

	return result
}
