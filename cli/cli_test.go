package cli

import (
	"testing"

	"github.com/wawandco/ox/lifecycle/build"
	"github.com/wawandco/ox/lifecycle/dev"
	"github.com/wawandco/ox/lifecycle/generate"
	"github.com/wawandco/ox/lifecycle/test"
	"github.com/wawandco/ox/plugins"
	"github.com/wawandco/ox/tools/cli/help"
	"github.com/wawandco/ox/tools/cli/version"
	"github.com/wawandco/ox/tools/db"
)

func Test_CliTestingAliaser(t *testing.T) {
	plugins := []plugins.Plugin{
		&generate.Command{},
		&build.Command{},
		&dev.Command{},
		&version.Command{},
		&help.Command{},
		&test.Command{},
		&db.Command{},
	}

	c := &cli{
		plugins,
	}

	tcases := []struct {
		commandAlias string
		nameExpected string
	}{
		{"g", "generate"},
		{"b", "build"},
		{"d", "dev"},
		{"v", "version"},
		{"h", "help"},
		{"db", "database"},
		{"t", "test"},
	}

	for _, ca := range tcases {
		command := c.findCommand(ca.commandAlias)
		if command == nil {
			t.Errorf("Command %s not found", ca.commandAlias)
			continue
		}

		if command.Name() != ca.nameExpected {
			t.Errorf("Not equal")
		}
	}

}
