package proxy

import (
	"Multi-Honeypot/config/plugin"
	mutex2 "Multi-Honeypot/internal/pkg/mutex"
	plugin2 "Multi-Honeypot/internal/pkg/plugin"
	"Multi-Honeypot/internal/pkg/utils"
	"fmt"
	"net"
)

/*
	构建中间层代理
*/
func HandleConnection(app *plugin2.AppStruct, plugin plugin.Plugin) {
	port := app.Config.Get(plugin.Name, "port") // 插件监听的端口
	fromListener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	app.Log.Log(app.Log.INFO, plugin.Name, "Listening : ", port)

	for {
		fromConn, err := fromListener.Accept() // 接收请求

		ctx := plugin2.NewContext(app) // 建立上下文
		ip := fromConn.RemoteAddr().String()
		mutex := mutex2.NewMutex()

		app.Log.Log(app.Log.INFO, plugin.Name, "Connection From: ", ip)
		app.Log.Log(app.Log.INFO, plugin.Name, "Starting Backend: ", port)

		randomName := utils.RandStr(5) // 随机文件名
		proxySocks := fmt.Sprintf("/tmp/%s-%s.socks", plugin.Name, randomName)

		mutex.NewItem(proxySocks) // 建立互斥锁

		ctx.Mutex = mutex
		ctx.SocksFile = proxySocks
		ctx.PluginName = plugin.Name
		ctx.RemoteIP = ip

		app.Log.Log(app.Log.INFO, plugin.Name, fmt.Sprintf("Creating exchange proxy %s <------> %s", port, proxySocks))

		go plugin.BackendProcess(ctx) // 启动后端处理程序

		mutex.Wait(proxySocks) // 等待建立链路

		app.Log.Log(app.Log.INFO, plugin.Name, fmt.Sprintf("%s <---Link---> %s Success!", port, proxySocks))

		toCon, err := net.Dial("unix", proxySocks)
		if err != nil {
			panic(err)
		}

		bufferWriter, bufferExit := plugin.BufferProcess(ctx) // 获取处理函数
		go exchangeData(fromConn, toCon, bufferWriter, bufferExit, 0)
		go exchangeData(toCon, fromConn, bufferWriter, bufferExit, 1)

	}
	defer fromListener.Close()
}

func exchangeData(r, w net.Conn, bufferProcess func([]byte, int), bufferExit func(), mod int) {
	defer r.Close()
	defer w.Close()

	var buffer = make([]byte, 10000000)
	for {
		n, err := r.Read(buffer)
		if err != nil {
			bufferExit()
			break
		}

		bufferProcess(buffer[:n], mod)

		n, err = w.Write(buffer[:n])
		if err != nil {
			bufferExit()
			break
		}
	}

}
