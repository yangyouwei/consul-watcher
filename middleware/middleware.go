package middleware

import (
	"github.com/yangyouwei/consul-watcher/loglib"
	"net/http"
)

//request logging
func LoggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loglib.Mylog.Println("request domain: ", r.Host, "URL :", r.URL)
		next.ServeHTTP(w, r)
	})
}

//crossdomain conf
func CrossDomainMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                                                              //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Connection, User-Agent, Cookie,Action, Module") //header的类型
		w.Header().Set("content-type", "application/json")                                                                                              //返回数据格式是json
		//header("Access-Control-Allow-Credentials : true");
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}