package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Test3Prop struct {
	gorm.Model
	Prop    string
	Test3ID uint
}
type Test3 struct {
	gorm.Model
	Name      string
	Test4ID   uint
	Test3Prop *Test3Prop
}
type Test4 struct {
	gorm.Model
	Name   string
	Test3s []Test3
}

func One2Many() {
	//CreateTableByModel(&Test3{}, &Test4{})
	//t31 := Test3{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	Name: "T3-1",
	//}
	//t32 := Test3{
	//	Model: gorm.Model{
	//		ID: 2,
	//	},
	//	Name: "T3-1",
	//}
	//t41 := Test4{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	Name:   "t41",
	//	Test3s: []Test3{t31, t32},
	//}
	//DB.Create(&t41)
	//Preload关联查询
	//var test4Domain Test4
	//DB.Preload("Test3s").First(&test4Domain)
	//fmt.Println(test4Domain)
	//Preload添加过滤条件
	//var test4Domain Test4
	//DB.Preload("Test3s", "name = ?", "T3-1").First(&test4Domain)
	//fmt.Println(test4Domain)
	//Preload传递参数
	//var test4Domain Test4
	//DB.Preload("Test3s", func(db *gorm.DB) *gorm.DB {
	//	return db.Where("name = ?", "T3-1")
	//}).First(&test4Domain)
	//fmt.Println(test4Domain)
	//test3Prop1 := Test3Prop{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	Prop:    "京津冀",
	//	Test3ID: 1,
	//}
	//DB.Create(&test3Prop1)
	//链式预加载
	//var test4 Test4
	//DB.Preload("Test3s.Test3Prop").Preload("Test3s").First(&test4)
	//fmt.Println(test4.Test3s[0].Test3Prop)
	//var test4 Test4
	//DB.Preload("Test3s.Test3Prop", "Prop = ?", "京津冀").Preload("Test3s", "ID = 1").First(&test4)
	//fmt.Println(test4.Test3s[0].Test3Prop)
	var test4Domain Test4
	DB.Preload("Test3s", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Test3Prop").Where("Prop = ?", "京津冀")
	}).First(&test4Domain)
	fmt.Println(test4Domain.Test3s[0].Test3Prop)
}
