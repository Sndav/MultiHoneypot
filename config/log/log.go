package log

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func LogInit(){
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
