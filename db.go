package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

var db *sqlx.DB

func init() {
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	database, err := sqlx.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/dc")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	db = database
	// 注意这行代码要写在上面err判断的下面
	//defer db.Close()
}

func GetDB() *sqlx.DB {
	return db
}

type Ind struct {
	id          int
	gmtCreate   time.Time
	gmtModified time.Time
	metaType    string
	code        string
	name        string
	comment     string
}

func QueryIndList(code string) []Ind {
	db := GetDB()

	var indList []Ind
	err := db.Select(&indList, "SELECT id, gmt_create as gmtCreate, gmt_modifed as gmtModifed, meta_type as metaType, code, name, comment from meta_info ")
	if err != nil {
		fmt.Println("exec Ind query failed, ", err)
		return nil
	}
	return indList
}

func InsertInd(meta *Ind) {
	db := GetDB()

	sql := "insert into meta_info(id, gmt_create, gmt_modifed, meta_type, code, name, comment) values(?, ?, ?, ?, ?, ?, ?)"

	r, err := db.Exec(sql, 1, time.Now(), time.Now(), "ind", "code", "code", "code")
	if err != nil {
		fmt.Println("exec insert failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec insert failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}

func ShowTables() {
	r, err := db.Exec("SHOW TABLES")
	if (err != nil) {
		fmt.Println("查询失败：", err)
	}

	fmt.Println("表列表：", r)

}