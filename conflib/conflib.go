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

type UpsConf struct {
	 Upstpl string
	 UpstreamPath string
	 DyupsUrl string
}

var Dingding string

var Upsconf UpsConf

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
	Upsconf.GETCONF(cfg)
	getdingding(cfg)
	fmt.Println("config init success.")
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

func (u *UpsConf)GETCONF(cfg *goconfig.ConfigFile)  {
	key, err := cfg.GetSection("ups")
	if err != nil {
		log.Panic(err)
	}
	for k, v := range key {
		switch k {
		case "tpl_path":
			u.Upstpl = v
		case "upstream_path":
			u.UpstreamPath = v
		case "dyups_url":
			u.DyupsUrl = v
		}
	}
}

func getdingding(cfg *goconfig.ConfigFile)  {
	key, err := cfg.GetSection("dingding")
	if err != nil {
		log.Panic(err)
	}
	for k, v := range key {
		switch k {
		case "token_url":
			Dingding = v
		}
	}
}

func stringtobool(s string) bool {
	if s == "true" {
		return true
	}
	return false
}