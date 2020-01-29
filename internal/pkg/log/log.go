package log

import (
	"Muti-Honeypot/config/log"
	"github.com/sirupsen/logrus"
)

type logType int32

type Log struct {
	INFO    logType
	WARNING logType
	ERROR   logType
	OTHER   logType
	PANIC   logType
}

const (
	INFO    logType = 0
	WARNING logType = 1
	ERROR   logType = 2
	OTHER   logType = 3
	PANIC   logType = 4
)

func NewLog() *Log{
	log        := &Log{}
	log.INFO    = 0
	log.WARNING = 1
	log.ERROR   = 2
	log.OTHER   = 3
	log.PANIC   = 4
	return log
}

func (log *Log) Log(logTypes logType,err ...interface{}){

	switch logTypes {
	case INFO:
		logrus.Info(err...)
		break
	case WARNING:
		logrus.Warn(err...)
		break
	case ERROR:
		logrus.Error(err...)
		break
	case OTHER:
		logrus.Print(err...)
		break
	case PANIC:
		logrus.Panic(err...)
		break
	default:
		logrus.Info(err...)
		break
	}
}


func init(){
	log.LogInit()
}