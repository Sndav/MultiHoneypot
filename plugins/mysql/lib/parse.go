package lib

import (
	"Multi-Honeypot/plugins/mysql"
	"encoding/binary"
	"fmt"
	"log"
)

func (stm *stream) ParseTraffic(payload []byte, seq int) {

	var msg string
	switch payload[0] {

	case COM_INIT_DB:

		msg = fmt.Sprintf("USE %s;\n", payload[1:])
	case COM_DROP_DB:

		msg = fmt.Sprintf("Drop DB %s;\n", payload[1:])
	case COM_CREATE_DB, COM_QUERY:

		statement := string(payload[1:])
		msg = fmt.Sprintf("%s %s", ComQueryRequestPacket, statement)
	case COM_STMT_PREPARE:

		serverPacket := stm.findStmtPacket(stm.packets, seq+1)
		if serverPacket == nil {
			log.Println("ERR : Not found stm packet")
			return
		}

		//fetch stm id
		stmtID := binary.LittleEndian.Uint32(serverPacket.payload[1:5])
		stmt := &Stmt{
			ID:    stmtID,
			Query: string(payload[1:]),
		}

		//record stm sql
		stm.stmtMap[stmtID] = stmt
		stmt.FieldCount = binary.LittleEndian.Uint16(serverPacket.payload[5:7])
		stmt.ParamCount = binary.LittleEndian.Uint16(serverPacket.payload[7:9])
		stmt.Args = make([]interface{}, stmt.ParamCount)

		msg = PreparePacket + stmt.Query
	case COM_STMT_SEND_LONG_DATA:

		stmtID := binary.LittleEndian.Uint32(payload[1:5])
		paramId := binary.LittleEndian.Uint16(payload[5:7])
		stmt, _ := stm.stmtMap[stmtID]

		if stmt.Args[paramId] == nil {
			stmt.Args[paramId] = payload[7:]
		} else {
			if b, ok := stmt.Args[paramId].([]byte); ok {
				b = append(b, payload[7:]...)
				stmt.Args[paramId] = b
			}
		}
		return
	case COM_STMT_RESET:

		stmtID := binary.LittleEndian.Uint32(payload[1:5])
		stmt, _ := stm.stmtMap[stmtID]
		stmt.Args = make([]interface{}, stmt.ParamCount)
		return
	case COM_STMT_EXECUTE:

		var pos = 1
		stmtID := binary.LittleEndian.Uint32(payload[pos : pos+4])
		pos += 4
		var stmt *Stmt
		var ok bool
		if stmt, ok = stm.stmtMap[stmtID]; ok == false {
			log.Println("ERR : Not found stm id", stmtID)
			return
		}

		//params
		pos += 5
		if stmt.ParamCount > 0 {

			//（Null-Bitmap，len = (paramsCount + 7) / 8 byte）
			step := int((stmt.ParamCount + 7) / 8)
			nullBitmap := payload[pos : pos+step]
			pos += step

			//Parameter separator
			flag := payload[pos]

			pos++

			var pTypes []byte
			var pValues []byte

			//if flag == 1
			//n （len = paramsCount * 2 byte）
			if flag == 1 {
				pTypes = payload[pos : pos+int(stmt.ParamCount)*2]
				pos += int(stmt.ParamCount) * 2
				pValues = payload[pos:]
			}

			//bind params
			err := stmt.BindArgs(nullBitmap, pTypes, pValues)
			if err != nil {
				log.Println("ERR : Could not bind params", err)
			}
		}
		msg = string(stmt.WriteToText())
	default:
		return
	}

	msg = fmt.Sprintln(mysql.GetNowStr(true) + msg)
	fmt.Print(msg)
	stm.msg += msg
}
