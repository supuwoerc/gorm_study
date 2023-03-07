package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Test7 struct {
	gorm.Model
	Name   string
	Test9  []Test9  `gorm:"polymorphic:Owner;polymorphicValue:t7;"`
	Test10 []Test10 `gorm:"foreignKey:Test7ForeignKey;references:Name"`
	Test11 []Test11 `gorm:"many2many:test7_test11;foreignKey:Name;joinForeignKey:t7;references:T11Name;joinReferences:t11"`
}

type Test8 struct {
	gorm.Model
	Name  string
	Test9 []Test9 `gorm:"polymorphic:Owner;polymorphicValue:t8"`
}
type Test9 struct {
	gorm.Model
	Name      string
	OwnerType string
	OwnerID   uint
}
type Test10 struct {
	gorm.Model
	Name            string
	Test7ForeignKey string
}
type Test11 struct {
	gorm.Model
	T11Name string
}

func Polymorphic() {
	fmt.Println("多态")
	CreateTableByModel(&Test7{}, &Test8{}, &Test9{}, &Test10{}, &Test11{})
	//DB.Create(&Test8{
	//	Name: "我是test8_1",
	//	Test9: []Test9{{
	//		Name: "Test9_1",
	//	}},
	//})
	//DB.Create(&Test7{
	//	Name: "Test7_1",
	//	Test9: []Test9{{
	//		Name: "Test9_2",
	//	}, {
	//		Name: "Test9_3",
	//	}},
	//})
	//指定外键名称和引用字段
	//DB.Create(&Test7{
	//	Name:  "Test7_1",
	//	Test9: nil,
	//	Test10: []Test10{{
	//		Name: "Test10_1",
	//	}, {
	//		Name: "Test10_2",
	//	}},
	//})
	//多对多外键设置
	DB.Create(&Test7{
		Name:   "Test77777",
		Test9:  nil,
		Test10: nil,
		Test11: []Test11{
			{T11Name: "T11-1"},
			{T11Name: "T11-2"},
			{T11Name: "T11-3"},
		},
	})
}
