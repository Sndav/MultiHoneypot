package proxy

import (
	"Muti-Honeypot/config/plugin"
	plugin2 "Muti-Honeypot/internal/pkg/plugin"
	"Muti-Honeypot/internal/pkg/utils"
	"fmt"
	"net"
)


func HandleConnection(app *plugin2.AppStruct, plugin plugin.Plugin) {
	port := app.Config.Get(plugin.Name,"port") // 插件监听的端口
	fromListener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	app.Log.Log(app.Log.INFO,"Listening : ",port)

	for {
		fromConn, err := fromListener.Accept()
		ip := fromConn.RemoteAddr().String()

		app.Log.Log(app.Log.INFO,"Connection From: ",ip)
		app.Log.Log(app.Log.INFO,"Starting Backend: ",port)

		randomName := utils.RandStr(5)                  // 随机文件名
		proxySocks := fmt.Sprintf("/tmp/%s-%s.socks",plugin.Name,randomName)

		app.Log.Log(app.Log.INFO,fmt.Sprintf("Creating Exchange Proxy %s <------> %s", port, proxySocks))

		app.Keys["pluginName"] = plugin.Name

		ClientConn, err := net.Listen("unix",proxySocks) // 建立代理链接
		if err != nil{
			panic(err)
		}

		app.Log.Log(app.Log.INFO,"Success")

		go plugin.BackendProcess(app, ClientConn)   // 建立处理协程

		toCon, err := net.Dial("unix", proxySocks)
		if err != nil {
			panic(err)
		}
		bufferWriter,bufferExit := plugin.BufferProcess(app,ip)
		go exchangeData(fromConn, toCon, bufferWriter,bufferExit , 0)
		go exchangeData(toCon, fromConn, bufferWriter,bufferExit , 1)

	}
	defer fromListener.Close()
}

func exchangeData(r, w net.Conn, bufferProcess func([]byte,int) , bufferExit func(),mod int) {
	defer r.Close()
	defer w.Close()

	var buffer = make([]byte, 100000)
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