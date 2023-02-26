package main

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type User struct {
	gorm.Model
	Name string `gorm:"default:default_name"`
	Age  uint   `gorm:"comment:年龄"`
}

// CreateTableByModel 自动迁移数据库模型
func CreateTableByModel(model ...interface{}) {
	for _, v := range model {
		if DB.Migrator().HasTable(v) {
			_ = DB.Migrator().DropTable(v)
			fmt.Println("删除数据库表" + reflect.TypeOf(v).String())
		}
	}
	err := DB.AutoMigrate(model...) //这里不要粗心不展开  model是一个切片
	if err != nil {
		fmt.Println("迁移失败" + err.Error())
	} else {
		var names = make([]string, 0)
		for _, v := range model {
			names = append(names, reflect.TypeOf(v).String())
		}
		fmt.Println(strings.Join(names, "") + "-->迁移成功成功")
	}
}
