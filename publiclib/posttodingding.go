package publiclib

import (
	"fmt"
	"github.com/yangyouwei/consul-watcher/conflib"
	"github.com/yangyouwei/consul-watcher/loglib"
	"io/ioutil"
	"net/http"
	"strings"
)

func POSTMgToDingding(s *Service)  {
	url := conflib.Dingding
	responestr := ""
	servernum := 0
	for n,addr_port := range s.Servers{
		a := addr_port.Addr+":"+addr_port.Port+";"
		responestr = responestr+a
		servernum = n+1
	}

	//mes := fmt.Sprintf("test :server-st change! anzhi-server服务出现变化,当前运行数量:「%v」.服务器地址:「%s」",servernum,responestr)
	mes := fmt.Sprintf(`{"msgtype": "text","text": {"content": "server-st change!：%v服务出现变化,当前运行数量:%v,服务器地址：%s"}}`,s.Name,servernum,responestr)
	//fmt.Println(mes)

	resp, err := http.Post(url, "application/json", strings.NewReader(mes))
	if err != nil {
		loglib.Mylog.Println(err)
		loglib.Mylog.Println("send messages to dingding  fail.")
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		loglib.Mylog.Println(err)
		loglib.Mylog.Println("curl dyups, add server fail.")
		return
	}
	loglib.Mylog.Println("dingding respone: ",string(body))
}