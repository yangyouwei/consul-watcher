package api

import (
	"net/http"
)

//type UserInfo struct {
//	Name    string `json:"name"`
//	Age     int    `json:"age"`
//	Address string `json:"address"`
//}
//
//var Users = make(map[string]UserInfo)

func DoSmoeThing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world~!"))
}

//read post json
//func SETUserInfo_bodyjson(w http.ResponseWriter, r *http.Request) {
//	userjson, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	auser := UserInfo{}
//	// 将json数据解析到结构体实例中
//	err = json.Unmarshal(userjson, &auser)
//	if err != nil {
//		fmt.Println(err)
//	}
//	Users[auser.Name] = auser
//	w.Write([]byte("user set is ok"))
//}
