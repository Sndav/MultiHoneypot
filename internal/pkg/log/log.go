package log

import (
	"Multi-Honeypot/config/log"
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

func NewLog() *Log {
	log := &Log{}
	log.INFO = 0
	log.WARNING = 1
	log.ERROR = 2
	log.OTHER = 3
	log.PANIC = 4
	return log
}

func (log *Log) Log(logTypes logType, processName string, msg ...interface{}) {
	logHandler := logrus.WithFields(logrus.Fields{
		"process": processName,
	})
	switch logTypes {
	case INFO:
		logHandler.Info(msg...)
		break
	case WARNING:
		logHandler.Warn(msg...)
		break
	case ERROR:
		logHandler.Error(msg...)
		break
	case OTHER:
		logHandler.Print(msg...)
		break
	case PANIC:
		logHandler.Panic(msg...)
		break
	default:
		logHandler.Info(msg...)
		break
	}
}

func init() {
	log.LogInit()
}
