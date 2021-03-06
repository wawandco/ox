package help

import (
	"context"
	"strings"
	"testing"

	"github.com/wawandco/ox/plugins/core"
)

func TestFindCommand(t *testing.T) {
	hp := Command{
		commands: []core.Command{},
	}

	migrate := &subPl{}
	pop := &testPlugin{}
	pop.Receive([]core.Plugin{
		migrate,
	})

	hp.commands = append(hp.commands, pop)

	t.Run("not enough arguments", func(*testing.T) {
		result, names := hp.findCommand([]string{"help"})
		if result != nil || names != nil {
			t.Fatal("Should be nil")
		}
	})

	t.Run("top level command", func(t *testing.T) {
		result, names := hp.findCommand([]string{"help", "database"})
		expected := []string{
			"database",
		}

		if result.Name() != "database" || strings.Join(names, " ") != strings.Join(expected, " ") {
			t.Fatal("didn't find our guy")
		}
	})

	t.Run("subcommand lookup", func(*testing.T) {
		result, names := hp.findCommand([]string{"help", "database", "migrate"})
		expected := []string{
			"database",
			"migrate",
		}

		ht, ok := result.(core.HelpTexter)
		if result.Name() != "migrate" || !ok || ht.HelpText() != migrate.HelpText() || strings.Join(names, " ") != strings.Join(expected, " ") {
			t.Fatal("didn't find our guy")
		}
	})

	t.Run("extra args on non-subcommander", func(*testing.T) {
		result, names := hp.findCommand([]string{"help", "database", "migrate", "other", "thing"})
		expected := []string{
			"database",
			"migrate",
		}
		ht, ok := result.(core.HelpTexter)
		if result.Name() != "migrate" || !ok || ht.HelpText() != migrate.HelpText() || strings.Join(names, " ") != strings.Join(expected, " ") {
			t.Fatal("didn't find our guy")
		}
	})

}

type testPlugin struct {
	subcommands []core.Command
}

func (tp testPlugin) Name() string {
	return "database"
}

func (tp testPlugin) ParentName() string {
	return ""
}

func (tp testPlugin) HelpText() string {
	return "pop help text"
}

func (tp *testPlugin) Run(ctx context.Context, root string, args []string) error {
	return nil
}

func (tp *testPlugin) Receive(pls []core.Plugin) {
	for _, pl := range pls {
		c, ok := pl.(core.Command)
		if !ok || c.ParentName() != tp.Name() {
			continue
		}

		tp.subcommands = append(tp.subcommands, c)
	}
}

func (tp *testPlugin) Subcommands() []core.Command {
	return tp.subcommands
}

type subPl struct{}

func (tp subPl) Name() string {
	return "migrate"
}

func (tp subPl) ParentName() string {
	return "database"
}

func (tp subPl) HelpText() string {
	return "migrate help text"
}

func (tp subPl) Run(ctx context.Context, root string, args []string) error {
	return nil
}
