package api

import (
	"github.com/yangyouwei/consul-watcher/loglib"
	"github.com/yangyouwei/consul-watcher/publiclib"
	"io/ioutil"
	"net/http"
)

func PostFromConsul(w http.ResponseWriter, r *http.Request) {
		res, err := ioutil.ReadAll(r.Body)
		if err != nil {
			loglib.Mylog.Println(err)
			return
		}
	service,err := publiclib.GetRes(string(res))

	if service.Name == "" {
		loglib.Mylog.Println("decode consul post json erro")
		w.Write([]byte("decode consul post json error"))
		return
	}

	w.Write([]byte("ok"))

	//生成配置文件
	publiclib.GenConf(service)
	//curl dyups api
	publiclib.PostUps(service)
}

