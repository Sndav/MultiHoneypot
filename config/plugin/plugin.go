package plugin

import (
	"Muti-Honeypot/internal/pkg/plugin"
	"Muti-Honeypot/plugins/test"
)

type Plugin struct {
	Name string
	BackendProcess plugin.BackendHandler
	BufferProcess plugin.BufferProcessHandler
}

func GetPlugins() []Plugin {
	return []Plugin{
		{"test",test.Entrance,test.BufferProcess},
	}
}