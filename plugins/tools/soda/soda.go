package soda

import (
	"github.com/gobuffalo/packd"
	"github.com/wawandco/ox/plugins/core"
)

func Plugins(migrations packd.Box) []core.Plugin {
	pl := []core.Plugin{
		&Command{migrations: migrations},
	}

	return pl
}
