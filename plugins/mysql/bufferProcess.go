package mysql

import (
	"Multi-Honeypot/internal/pkg/plugin"
)

func BufferProcess(ctx *plugin.Context) (func([]byte, int), func()) {

	stream := NewStream()

	go stream.Consumer()
	bufferWriter := func(bytes []byte, mod int) {
		packet := NewPacket(bytes, mod)
		stream.packets <- packet
	}

	bufferExit := func() {
		err := ctx.App.DB.InsertPhishing(ctx.PluginName, ctx.RemoteIP, stream.msg)
		if err != nil {
			panic(err)
		}
	}
	return bufferWriter, bufferExit
}
