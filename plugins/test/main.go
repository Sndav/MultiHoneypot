package test

import (
	"Muti-Honeypot/internal/pkg/plugin"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func Entrance(app *plugin.AppStruct,conn net.Listener ,args ...interface{}){
	for {
		fromConn, err := conn.Accept()
		if err != nil {
			panic(err)
		}
		go Read(fromConn)
	}
}

func Read(r net.Conn) {
	defer r.Close()

	var buffer = make([]byte, 100000)
	for {
		_, err := r.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:]))
	}

}

func BufferProcess (app *plugin.AppStruct, ip string) (func([]byte,int),func()) {
	// 处理单个请求，返回 buffer 记录函数
	pluginName := app.Keys["pluginName"].(string)
	prefix := md5sum(pluginName+ip)[:5] + "-"
	tmpFile, err := ioutil.TempFile(os.TempDir(),prefix)
	if err != nil{
		panic(err)
	}
	app.Log.Log(app.Log.INFO,"Making temp file: ",tmpFile.Name())
	bufferWriter := func (bytes []byte,mod int) {
		switch mod {
		case 0:
			_,_ = tmpFile.WriteString("-----Client------\n")
			app.Log.Log(app.Log.INFO,"Client sent msg")
			break
		case 1:
			_,_ = tmpFile.WriteString("-----Server------\n")
			app.Log.Log(app.Log.INFO,"Server sent msg")
			break
		}

		n,err := tmpFile.Write(bytes)
		app.Log.Log(app.Log.INFO,"Writed to ",tmpFile.Name()," ", n," bytes")
		if err != nil{
			panic(err)
		}
	}
	bufferExit := func(){
		app.Log.Log(app.Log.INFO,tmpFile.Name()," Closed")
		_ = tmpFile.Close()
		if fileExists(tmpFile.Name()){
			file,err := os.Open(tmpFile.Name())
			if err != nil {
				log.Fatal(err)
			}

			defer file.Close()
			fileinfo,err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}

			fileSize := fileinfo.Size()
			buffer := make([]byte,fileSize)

			bytesread,err := file.Read(buffer)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("bytes read:",bytesread)
			err = app.DB.InsertPhishing(pluginName,ip,string(buffer))
			if err != nil{
				panic(err)
			}

			_ = os.Remove(tmpFile.Name())
		}
	}
	return bufferWriter, bufferExit
}

func md5sum(content string) string{
	h := md5.New()
	_, err := io.WriteString(h,content)
	if err != nil{
		panic(err)
	}
	return fmt.Sprintf("%x",h.Sum(nil))
}

func fileExists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}