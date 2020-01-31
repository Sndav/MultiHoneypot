package plugin

import (
	"Multi-Honeypot/internal/pkg/config"
	"Multi-Honeypot/internal/pkg/database"
	"Multi-Honeypot/internal/pkg/log"
	"Multi-Honeypot/internal/pkg/mutex"
)

type AppStruct struct {
	Config *config.Config
	DB     *database.DB
	Log    *log.Log
}

type Context struct {
	App        *AppStruct
	Mutex      *mutex.Mutex
	PluginName string
	SocksFile  string
	RemoteIP   string
	Keys       map[string]interface{}
}

type BackendHandler func(*Context)
type BufferProcessHandler func(*Context) (func([]byte, int), func())
type BufferExit func(*Context) func()

func NewContext(app *AppStruct) *Context {
	ctx := &Context{App: app}
	ctx.Keys = make(map[string]interface{})
	return ctx
}

// 上下文变量获取
func (c *Context) Set(key string, val interface{}) {
	c.Keys[key] = val
}

func (c *Context) Get(key string) interface{} {
	return c.Keys[key]
}
