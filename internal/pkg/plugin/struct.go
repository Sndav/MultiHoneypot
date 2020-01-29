package plugin

import (
	"Muti-Honeypot/internal/pkg/config"
	"Muti-Honeypot/internal/pkg/database"
	"Muti-Honeypot/internal/pkg/log"
	"net"
)

type AppStruct struct {
	Config   *config.Config
	DB       *database.DB
	Log      *log.Log
	Keys     map[string]interface{}
}

type BackendHandler func(*AppStruct, net.Listener ,...interface{})
type BufferProcessHandler func(*AppStruct, string) (func([]byte,int),func())
type BufferExit func(*AppStruct, string) func()
