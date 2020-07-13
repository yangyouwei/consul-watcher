package main

import (
	"flag"
	"github.com/yangyouwei/consul-watcher/api"
	"github.com/yangyouwei/consul-watcher/conflib"
	"github.com/yangyouwei/consul-watcher/loglib"
	"github.com/yangyouwei/consul-watcher/router"
)

func init()  {
	s := flag.String("c", "./conf.ini", "-c /etc/main.conf")
	flag.Parse()
	//解析参数
	if *s == "" {
		flag.Usage()
		panic("process exit!")
	}
	conflib.InitConf(s)
	loglib.InitLog()
}

func main()  {
	go api.ReceiveService()
	router.InitRouter()
}
