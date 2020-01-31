package mysql

import (
	"fmt"
)

type stream struct {
	packets    chan *packet
	packetList map[int]packet
	msg        string
}

type packet struct {
	isClientFlow bool
	seq          int
	length       int
	payload      []byte
}

func NewStream() *stream {
	s := &stream{}
	s.packets = make(chan *packet)
	s.packetList = make(map[int]packet)
	return s
}

func NewPacket(buf []byte, mod int) *packet {
	p := &packet{}
	if mod == 1 {
		p.isClientFlow = false
	} else {
		p.isClientFlow = true
	}
	seq := buf[3]
	p.seq = int(seq)
	p.length = len(buf) - 4
	p.payload = buf[4:]
	return p
}

func (stream *stream) Consumer() {
	for {
		packet := <-stream.packets
		stream.packetList[packet.seq] = *packet
		stream.DoParse(*packet)
	}
}

func (stm *stream) DoParse(packet packet) {
	payload := packet.payload

	var msg string
	switch payload[0] {

	case COM_INIT_DB:

		msg = fmt.Sprintf("USE %s;\n", payload[1:])
	case COM_DROP_DB:

		msg = fmt.Sprintf("Drop DB %s;\n", payload[1:])
	case COM_CREATE_DB, COM_QUERY:

		statement := string(payload[1:])
		msg = fmt.Sprintf("%s %s", ComQueryRequestPacket, statement)

	default:
		return
	}
	fmt.Println("MySQL: ", msg)
	stm.msg += msg + "\n"
}
