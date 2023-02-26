package main

import (
	"gorm.io/gorm"
)

type Test1 struct {
	gorm.Model
	Name    string
	Test2Id uint
}

type Test2 struct {
	gorm.Model
	Name  string
	Test1 Test1
}

func One2One() {
	//CreateTableByModel(&Test2{}, &Test1{})
	//DB.Create(&Test2{
	//	Name: "Test2-测试1",
	//	Test1: Test1{
	//		Name: "Test1-测试",
	//	},
	//})
	//var test2 Test2
	//预加载
	//DB.Preload("Test1").First(&test2)
	//fmt.Println(test2)
	//创建具有关联关系的数据
	test1 := Test1{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "哈哈",
	}
	test2 := Test2{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "嘿嘿",
	}
	//DB.Create(&test2)
	//DB.Create(&test1)
	DB.Model(&test2).Association("Test1").Append(&test1)
}
