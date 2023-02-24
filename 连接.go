package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	username := "root"     //账号
	password := "12345678" //密码
	host := "127.0.0.1"    //地址
	port := 3306           //端口
	dbname := "gorm_study" //数据库名
	//可能会遇到tcp连接的问题：https://github.com/Masterminds/glide/issues/999
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(connectString))
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	} else {
		fmt.Println("数据库连接成功")
		DB = db
	}
}
func main() {
	fmt.Println("main is running...")
}
