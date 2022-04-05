package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 启动ES

	// 启动Redis

	//	测试MySQL

	//ind := Ind{
	//	id:          1,
	//	gmtCreate:   time.Now(),
	//	gmtModified: time.Now(),
	//	metaType:    "ind",
	//	code:        "code",
	//	name:        "name",
	//	comment:     "code",
	//}
	//InsertInd(&ind)

	//indList := QueryIndList("code")
	//fmt.Println(indList)

	ShowTables()
}
