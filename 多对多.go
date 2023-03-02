package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Info struct {
	gorm.Model
	Prop    string
	Test5Id uint
}
type Test5 struct {
	gorm.Model
	Info   *Info
	Test6s []*Test6 `gorm:"many2many:test5_test6;"`
}

type Test6 struct {
	gorm.Model
	Name   string
	Test5s []*Test5 `gorm:"many2many:test5_test6;"`
}

func Many2Many() {
	//CreateTableByModel(&Test5{}, &Test6{}, &Info{})
	//info_1 := Info{
	//	Prop: "我是info_1",
	//}
	//t6_1 := Test6{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	Name:   "我是t6_1",
	//	Test5s: nil,
	//}
	//t6_2 := Test6{
	//	Model: gorm.Model{
	//		ID: 2,
	//	},
	//	Name: "我是t6_2",
	//}
	//t5_1 := Test5{
	//	Info:   &info_1,
	//	Test6s: []*Test6{&t6_1, &t6_2},
	//}
	//info_2 := Info{
	//	Prop: "info_2",
	//}
	//t5_2 := Test5{
	//	Info:   &info_2,
	//	Test6s: []*Test6{&t6_1, &t6_2},
	//}
	//DB.Create(&t5_1)
	//DB.Create(&t5_2)
	//var test5 Test5
	//DB.Preload("Test6s").Preload("Info").Where("id = 2").First(&test5)
	//fmt.Println(len(test5.Test6s))
	//fmt.Println(test5.Info)
	//var test6s []*Test6
	//DB.Model(&Test5{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//}).Preload("Test5s.Info").Association("Test6s").Find(&test6s)
	//fmt.Println(test6s[0])
	//fmt.Println(test6s[0].Test5s[0])
	//fmt.Println(test6s[0].Test5s[0].Info)

	var test6s []*Test6
	test5 := Test5{
		Model: gorm.Model{
			ID: 1,
		},
	}
	DB.Model(&test5).Preload("Test5s", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Info").Where("prop = ?", "我是info_1")
	}).Where("name = ?", "我是t6_2").Association("Test6s").Find(&test6s)
	fmt.Println(test6s)
	//fmt.Println(test6s[0])
	//fmt.Println(len(test6s[0].Test5s))
	//fmt.Println(test6s[0].Test5s[0].Info)
}
