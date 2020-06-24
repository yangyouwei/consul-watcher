package conflib

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
	"path/filepath"
)

type MainConf struct {
	ListenPort string
	LogBool bool
	LogFileDir string
}

var Mainconf MainConf

func InitConf(s *string)  {
	cstr, err := filepath.Abs(*s)
	if err != nil {
		log.Panicln(err)
	}
	//初始化配置文件
	fmt.Println(cstr)
	cfg, err := goconfig.LoadConfigFile(cstr)
	if err != nil {
		log.Println("读取配置文件失败[config.ini]")
		log.Panic(err)
	}
	Mainconf.GETCONF(cfg)
	log.Println("config init success.")
}

func (m *MainConf)GETCONF(cfg *goconfig.ConfigFile)  {
	key, err := cfg.GetSection("main")
	if err != nil {
		log.Panic(err)
	}
	for k, v := range key {
		switch k {
		case "listen_port":
			m.ListenPort = v
		case "log":
			m.LogBool = stringtobool(v)
		case "log_file_dir":
			m.LogFileDir = v
		}
	}
}

func stringtobool(s string) bool {
	if s == "true" {
		return true
	}
	return false
}