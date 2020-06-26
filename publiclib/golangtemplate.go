package publiclib

import (
	"bytes"
	"github.com/yangyouwei/consul-watcher/conflib"
	"github.com/yangyouwei/consul-watcher/loglib"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func ReadTpl(p string) string {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		loglib.Mylog.Println(err)
	}
	return string(b)
}


func GenConf(s *Service)  {
	// 读取模板
	templateText := ReadTpl(conflib.Upsconf.Upstpl)
	buffer := new(bytes.Buffer)
	t := template.Must(template.New("upstream").Parse(templateText))
	err := t.Execute(buffer, s)
	if err != nil {
		loglib.Mylog.Println("Executing template:", err)
		return
	}

	//创建配置文件
	fd, err := os.OpenFile(conflib.Upsconf.UpstreamPath + s.Name+".conf", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fd.Close()
	fd.Write(buffer.Bytes())
	loglib.Mylog.Println("save cnf to file ,services: ",s)
}