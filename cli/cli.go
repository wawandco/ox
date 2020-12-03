package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/paganotoni/oxpecker/plugins"
	"github.com/paganotoni/oxpecker/plugins/cli/fixer"
	"github.com/paganotoni/oxpecker/plugins/cli/help"
	"github.com/paganotoni/oxpecker/plugins/cli/version"
)

// cli is the CLI wrapper for our tool. It is in charge
// for articulating different commands, finding it and also
// taking care of the CLI iteraction.
type cli struct {
	root    string
	Plugins []plugins.Plugin
}

// findCommand looks in the plugins for a command
// with the passed name.
func (c *cli) findCommand(name string) plugins.Command {
	for _, cm := range c.Plugins {
		// We skip subcommands on this case
		// those will be wired by the parent command implementing
		// Receive.
		if _, ok := cm.(plugins.Subcommand); ok {
			continue
		}

		command, ok := cm.(plugins.Command)
		if !ok {
			continue
		}

		pluginName := command.Name()
		if pn, ok := cm.(plugins.CommandNamer); ok {
			pluginName = pn.CommandName()
		}

		if pluginName != name {
			continue
		}

		return command
	}

	return nil
}

// Runs the CLI
func (c *cli) Run(ctx context.Context, pwd string, args []string) error {

	// IMPORTANT: Incorporate the plugin system by taking a look at this.
	// https://github.com/gobuffalo/buffalo-cli/blob/81f172713e1182412f27a0b128160386e04cd39b/internal/garlic/run.go#L28

	// Not sure if we should do this here or somewhere
	// else, these are some environment variables to be set
	// and other things to check.
	os.Setenv("GO111MODULE", "on") // Modules must be ON
	os.Setenv("CGO_ENABLED", "0")  // CGO disabled

	if len(args) < 2 {
		fmt.Println("no command provided, please provide one")
		return nil
	}

	// Passing args and plugins to those plugins that require them
	for _, plugin := range c.Plugins {
		pf, ok := plugin.(plugins.FlagParser)
		if ok {
			err := pf.ParseFlags(args[1:])
			if err != nil {
				fmt.Println(err)
			}
		}

		pr, ok := plugin.(plugins.PluginReceiver)
		if ok {
			pr.Receive(c.Plugins)
		}
	}

	command := c.findCommand(args[1])
	if command == nil {
		// TODO: print help ?
		fmt.Printf("did not find %s command\n", args[1])
		return nil
	}

	return command.Run(ctx, c.root, args[1:])
}

// New creates a CLI with the passed root and plugins. This becomes handy
// when specifying your own plugins.
func New() *cli {
	c := &cli{
		Plugins: []plugins.Plugin{
			&help.Help{},
			&version.Version{},
			&fixer.Fixer{},
		},
	}

	return c
}
