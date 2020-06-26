package loglib

import (
	"github.com/yangyouwei/consul-watcher/conflib"
	"log"
	"os"
	"path"
)

var Mylog *log.Logger

func InitLog() {
	process_name := path.Base(os.Args[0])
	if conflib.Mainconf.LogFileDir == "" {
		conflib.Mainconf.LogFileDir = process_name+".log"
	}
		if conflib.Mainconf.LogBool {
			logFile, err := os.OpenFile(conflib.Mainconf.LogFileDir, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				log.Fatal(err)
			}
			Mylog = log.New(logFile, "["+process_name+"] ", log.Ldate|log.Ltime|log.LstdFlags)
		} else {
			Mylog = log.New(os.Stdout, "["+process_name+"] ", log.Ldate|log.Ltime|log.LstdFlags)
		}
}
