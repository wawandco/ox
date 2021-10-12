package soda

import (
	"github.com/gobuffalo/packd"
	plugins "github.com/wawandco/ox/plugins/core"
)

func Plugins(migrations packd.Box) []plugins.Plugin {
	pl := []plugins.Plugin{
		&Command{migrations: migrations},
	}

	return pl
}
