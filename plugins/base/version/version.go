package version

import (
	"context"
	"fmt"
	"os"

	"github.com/wawandco/ox/plugins/base/content"
	"github.com/wawandco/ox/plugins/core"
)

var (
	// The version of the CLI
	version = "v0.12.0-beta.2"
)

var (
	// Version is a Command
	_ core.Command = (*Command)(nil)
)

// Command command will print X version.
type Command struct{}

func (b Command) Name() string {
	return "version"
}

func (c Command) Alias() string {
	return "v"
}

func (b Command) ParentName() string {
	return ""
}

func (b Command) HelpText() string {
	return "returns the current version of ox CLI"
}

// Run prints the version of the ox cli
func (b *Command) Run(ctx context.Context, root string, args []string) error {
	fmt.Println(content.Banner)
	fmt.Printf("Version %v\n", version)

	return nil
}

func (b *Command) FindRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	return wd
}
