package main

import "fmt"

func ExecSQL() {
	//执行原生sql
	DB.Exec("delete from gva_user where name = ?", "测试数据4")
	var users []User
	DB.Raw("select * from gva_user where name like ?", "%名字%").Scan(&users)
	fmt.Println(users)
}
