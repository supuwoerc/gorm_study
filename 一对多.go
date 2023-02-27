package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Test3 struct {
	gorm.Model
	Name    string
	Test4ID uint
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
	var test4Domain Test4
	DB.Preload("Test3s").First(&test4Domain)
	fmt.Println(test4Domain)
}
