package yarn

import plugins "github.com/wawandco/ox/plugins/core"

var (
	_ plugins.Plugin = (*Plugin)(nil)
)

type Plugin struct{}

func (p *Plugin) Name() string {
	return "yarn"
}
