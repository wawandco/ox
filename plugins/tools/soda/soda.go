package soda

import (
	"io/fs"

	"github.com/wawandco/ox/plugins/core"
)

func Plugins(migrations fs.FS) []core.Plugin {
	pl := []core.Plugin{
		&Command{migrations: migrations},
	}

	return pl
}
