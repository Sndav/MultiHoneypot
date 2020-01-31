package plugin

import (
	"Multi-Honeypot/internal/pkg/plugin"
	"Multi-Honeypot/plugins/mysql"
	"Multi-Honeypot/plugins/test"
)

type Plugin struct {
	Name           string
	BackendProcess plugin.BackendHandler
	BufferProcess  plugin.BufferProcessHandler
}

func GetPlugins() []Plugin {
	return []Plugin{
		{"test", test.Entrance, test.BufferProcess},
		{"mysql", mysql.Backend, mysql.BufferProcess},
	}
}
