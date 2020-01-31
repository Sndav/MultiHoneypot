package lib

import (
	"bytes"
	"io"
	"time"
)

type Mysql struct {
	port    int
	version string
	source  map[string]*stream
}

type stream struct {
	packets chan *packet
	stmtMap map[uint32]*Stmt
	msg     string
}

type packet struct {
	isClientFlow bool
	seq          int
	length       int
	payload      []byte
}

func (stm *stream) findStmtPacket(srv chan *packet, seq int) *packet {
	for {
		select {
		case packet, ok := <-stm.packets:
			if !ok {
				return nil
			}
			if packet.seq == seq {
				return packet
			}
		case <-time.After(5 * time.Second):
			return nil
		}
	}
}

func (stm *stream) resolve() {
	for {
		select {
		case packet := <-stm.packets:
			if packet.length != 0 {
				if packet.isClientFlow {
					stm.ParseTraffic(packet.payload, packet.seq)
				} else {
					stm.ParseTraffic(packet.payload, packet.seq)
				}
			}
		}
	}
}

func (m *Mysql) newPacket(buffer []byte, mod int) *packet {

	//read packet
	var payload bytes.Buffer
	var seq uint8
	var err error
	if seq, err = m.resolvePacketTo(buffer, &payload); err != nil {
		return nil
	}

	//generate new packet
	var pk = packet{
		seq:     int(seq),
		length:  payload.Len(),
		payload: payload.Bytes(),
	}
	if mod == 1 {
		pk.isClientFlow = false
	} else {
		pk.isClientFlow = true
	}

	return &pk
}

func (m *Mysql) resolvePacketTo(msg []byte, w io.Writer) (uint8, error) {

	header := msg[:4]

	length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)

	var seq uint8
	seq = header[3]
	_, _ = w.Write(msg[:int64(length)])

	return seq, nil
}
