package publiclib

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/yangyouwei/consul-watcher/loglib"
	"strconv"
)

type Server struct {
	Addr string
	Port string
}

type Service struct {
	Name string
	Servers []Server
}

func GetRes(jsonstr string ) (*Service,error) {
	n := gjson.Get(jsonstr,`#`)
	if fmt.Sprint(n.Type) == "Null" {
		loglib.Mylog.Println("parse json error")
		var val error = errors.New("parse json error")
		return &Service{},val
	}
	numstr := fmt.Sprint(n)
	num,err := strconv.Atoi(numstr)
	if err != nil {
		loglib.Mylog.Println("parse json error")
		var val error = errors.New("parse json error")
		return &Service{},val
	}

	s := Service{}

	for i:= 0;i<num ;i++ {
		addr := gjson.Get(jsonstr,fmt.Sprint(i)+".Service.Address")
		port := gjson.Get(jsonstr,fmt.Sprint(i)+".Service.Port")
		name := gjson.Get(jsonstr,fmt.Sprint(i)+".Service.Service")
		if fmt.Sprint(name) == "Null" {
			loglib.Mylog.Println("parse json error")
			var val error = errors.New("parse json error")
			return &Service{},val
		}

		server_add_port := Server{fmt.Sprint(addr),fmt.Sprint(port)}
		s.Name = fmt.Sprint(name)
		s.Servers = append(s.Servers,server_add_port)
	}

	return &s,nil
}