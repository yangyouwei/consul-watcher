package publiclib

import (
	"github.com/yangyouwei/consul-watcher/conflib"
	"github.com/yangyouwei/consul-watcher/loglib"
	"io/ioutil"
	"net/http"
	"strings"
)

func PostUps(s *Service)  {
	url := conflib.Upsconf.DyupsUrl+s.Name
	//fmt.Println(url)
	//fmt.Println(s)
	responestr := ""
	for _,addr_port := range s.Servers{
		a := "server "+addr_port.Addr+":"+addr_port.Port+";"
		responestr = responestr+a
	}

	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(responestr))
	if err != nil {
		loglib.Mylog.Println(err)
		loglib.Mylog.Println("curl dyups, add server fail.")
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		loglib.Mylog.Println(err)
		loglib.Mylog.Println("curl dyups, add server fail.")
		return
	}
	loglib.Mylog.Println("dyups repone: ",string(body))
}