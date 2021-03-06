package build

import (
	"github.com/spf13/pflag"
	"github.com/wawandco/ox/plugins/core"
)

func (b *Command) ParseFlags(args []string) {
	for _, plugin := range b.builders {
		fp, ok := plugin.(core.FlagParser)
		if !ok {
			continue
		}

		fp.ParseFlags(args)
	}
}

func (b *Command) Flags() *pflag.FlagSet {
	fs := pflag.NewFlagSet("build", pflag.ContinueOnError)
	fs.Usage = func() {}

	for _, plugin := range b.buildPlugins {
		fp, ok := plugin.(core.FlagParser)
		if !ok {
			continue
		}

		fp.Flags().VisitAll(func(f *pflag.Flag) {
			fs.AddFlag(f)
		})
	}

	return fs
}
