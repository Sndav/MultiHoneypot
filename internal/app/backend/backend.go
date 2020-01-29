package backend

import (
	"Muti-Honeypot/config/plugin"
	"Muti-Honeypot/internal/app/backend/proxy"
	"Muti-Honeypot/internal/pkg/config"
	"Muti-Honeypot/internal/pkg/database"
	"Muti-Honeypot/internal/pkg/log"
	plugin2 "Muti-Honeypot/internal/pkg/plugin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

type HandlerFunc func(app *App)

type App struct {
	Config   *config.Config
	DB       *database.DB
	Log      *log.Log
	Keys     map[string]interface{}
}


func NewApp(configFile string) *App{
	app := &App{}
	app.Log = log.NewLog()
	app.Config = config.NewConfig(configFile)
	app.Keys = make(map[string]interface{})
	app.initDB()
	return app
}

func (app *App) initDB(){
	var err error
	app.DB , err = database.NewDB(
		app.Config.Get("DB","type"),
		app.Config.Get("DB","connect_url"),
	)
	if err != nil{
		panic(err)
	}
}

func (app *App) run(){
	defer func(){
		if r := recover(); r != nil {
			app.Log.Log(app.Log.ERROR,r)
		}
	}()
	var wg sync.WaitGroup
	pluginList := plugin.GetPlugins()
	for _,item := range pluginList{
		appCopy := &plugin2.AppStruct{Config: app.Config, DB: app.DB, Log: app.Log, Keys: app.Keys}
		app.Log.Log(app.Log.INFO,"Starting ",item.Name)
		wg.Add(1)
		go func(){ // 建立中间代理抓取流量
			defer wg.Done()
			proxy.HandleConnection(appCopy,item)
		}()
	}
	wg.Wait()
}

func (app *App) Start(){
	app.run()
}