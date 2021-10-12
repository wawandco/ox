package yarn

import "github.com/wawandco/ox/plugins/core"

var (
	_ core.Plugin = (*Plugin)(nil)
)

type Plugin struct{}

func (p *Plugin) Name() string {
	return "yarn"
}
