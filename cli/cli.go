package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/wawandco/ox/internal/info"
	"github.com/wawandco/ox/internal/log"
	"github.com/wawandco/ox/plugins/base/content"
	"github.com/wawandco/ox/plugins/core"
)

// cli is the CLI wrapper for our tool. It is in charge
// of articulating different commands, finding it and also
// taking care of the CLI interaction.
type cli struct {
	Plugins []core.Plugin
}

// findCommand looks in the plugins for a command
// with the passed name.
func (c *cli) findCommand(name string) core.Command {
	for _, cm := range c.Plugins {
		// We skip subcommands on this case
		// those will be wired by the parent command implementing
		// Receive.
		command, ok := cm.(core.Command)
		if !ok || command.ParentName() != "" {
			continue
		}

		alias, ok := cm.(core.Aliaser)
		if ok && alias.Alias() == name {
			return command
		}

		if command.Name() == name {
			return command
		}

	}

	return nil
}

// Wrap Runs the CLI or cmd/ox/main.go
func (c *cli) Wrap(ctx context.Context, args []string) error {
	path := filepath.Join("cmd", "ox", "main.go")
	_, err := os.Stat(path)

	if name := info.ModuleName(); err != nil || name == "" || name == "github.com/wawandco/ox" {
		return c.Run(ctx, args)
	}

	// Fix command should not do the wrapping logic
	// we need to find a way to make this more generic.
	if args[1] == "fix" {
		return c.Run(ctx, args)
	}

	log.Infof("Using %v \n", path)

	cmd := exec.CommandContext(ctx, "go", "run")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	args[0] = path
	cmd.Args = append(cmd.Args, args...)

	return cmd.Run()
}

func (c *cli) Run(ctx context.Context, args []string) error {
	if len(args) < 2 {
		fmt.Println(content.Banner)
		log.Error("no command provided, please provide one")

		return nil
	}

	// Passing args and plugins to those plugins that require them
	for _, plugin := range c.Plugins {
		pf, ok := plugin.(core.FlagParser)
		if ok {
			pf.ParseFlags(args[1:])
		}

		pr, ok := plugin.(core.PluginReceiver)
		if ok {
			pr.Receive(c.Plugins)
		}
	}

	command := c.findCommand(args[1])
	if command == nil {
		fmt.Println(content.Banner)
		log.Infof("did not find %s command\n", args[1])
		return nil
	}

	// Commands that require running within the ox directory
	// may require its root to be determined with the go.mod. However,
	// some other commands may want to determine the root by themselves,
	// doing os.Getwd or something similar. The latter ones are RootFinders.
	root := info.RootFolder()
	rf, ok := command.(core.RootFinder)
	if root == "" && !ok {
		return errors.New("go.mod not found")
	}

	if root == "" {
		root = rf.FindRoot()
	}

	return command.Run(ctx, root, args[1:])
}

// Use passed Plugins by appending these to the
// plugins list inside the CLI.
func (c *cli) Use(plugins ...core.Plugin) {
	c.Plugins = append(c.Plugins, plugins...)
}

// Remove looks in the plugins list and removes plugins that
// match passed names.
func (c *cli) Remove(names ...string) {
	result := make([]core.Plugin, 0)
	for _, pl := range c.Plugins {
		var found bool
		for _, restricted := range names {
			if pl.Name() == restricted {
				found = true
			}
		}

		if found {
			continue
		}

		result = append(result, pl)
	}

	c.Plugins = result
}

// Clear the plugin list of the CLI.
func (c *cli) Clear() {
	c.Plugins = []core.Plugin{}
}
