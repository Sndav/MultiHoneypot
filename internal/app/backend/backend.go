package backend

import (
	"Multi-Honeypot/config/plugin"
	"Multi-Honeypot/internal/app/backend/proxy"
	"Multi-Honeypot/internal/pkg/config"
	"Multi-Honeypot/internal/pkg/database"
	"Multi-Honeypot/internal/pkg/log"
	plugin2 "Multi-Honeypot/internal/pkg/plugin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

type HandlerFunc func(app *App)

type App struct {
	Config *config.Config
	DB     *database.DB
	Log    *log.Log
}

func NewApp(configFile string) *App {
	app := &App{}
	app.Log = log.NewLog()
	app.Config = config.NewConfig(configFile)
	app.initDB()
	return app
}

func (app *App) initDB() {
	var err error
	app.DB, err = database.NewDB(
		app.Config.Get("DB", "type"),
		app.Config.Get("DB", "connect_url"),
	)
	if err != nil {
		panic(err)
	}
}

func (app *App) run() {
	defer func() {
		if r := recover(); r != nil {
			app.Log.Log(app.Log.ERROR, "Main", r)
		}
	}()
	var wg sync.WaitGroup
	pluginList := plugin.GetPlugins()
	for i, _ := range pluginList {
		//fmt.Println(pluginItem)
		appCopy := &plugin2.AppStruct{Config: app.Config, DB: app.DB, Log: app.Log}
		wg.Add(1)
		go func(plugin plugin.Plugin) { // 建立中间代理抓取流量
			defer wg.Done()
			app.Log.Log(app.Log.INFO, "Main", "Starting ", plugin.Name)
			proxy.HandleConnection(appCopy, plugin)
		}(pluginList[i])
	}
	wg.Wait()
}

func (app *App) Start() {
	app.run()
}
