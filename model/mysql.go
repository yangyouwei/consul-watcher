package model

//import (
//	"database/sql"
//	"fmt"
//	"github.com/yangyouwei/go-api/loglib"
//	"log"
//)
//
//var (
//	//mysql
//	MysqlString string = conflib.Mysql_conf_str.Username + ":" + conflib.Mysql_conf_str.Password + "@tcp(" + conflib.Mysql_conf_str.Ipaddress + ":" + conflib.Mysql_conf_str.Port + ")/" + conflib.Mysql_conf_str.DatabaseName
//	//mysql 连接池
//	Db *sql.DB
//)
//
//func init()  {
//	Db, err := sql.Open("mysql",MysqlString)
//	check(err)
//}
//
//func check(err error) {
//	if err != nil {
//		fmt.Println(err)
//		loglib.Mylog.Panic(err)
//	}
//}
//
////USE DB
//func SaveChapter(c chan ChapterInof,db *sql.DB)  {
//	trytime := 0
//	redo:
//		chpatertable := fmt.Sprintf("INSERT %v ( booksId, chapterName, size,content,chapterId) VALUES (?,?,?,?,?)", chapternum)
//		stmt, err := db.Prepare(chpatertable)
//		if trytime == 3 {
//			continue
//		}
//		if err != nil {
//			loglib.Mylog.Println(err)
//			trytime = trytime+1
//			goto redo
//		}
//
//		res, err := stmt.Exec(chater.Bookid,chater.Chaptername,chater.Size,chater.Content,chater.Chapterid)
//		if err != nil {
//			loglib.Mylog.Println(err)
//			trytime = trytime + 1
//			goto redo
//		}
//
//		_, err = res.LastInsertId() //必须是自增id的才可以正确返回。
//		if err != nil {
//			log.Println(err)
//		}
//		defer stmt.Close()
//		stmt.Close()
//}