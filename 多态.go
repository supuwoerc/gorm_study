package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Test7 struct {
	gorm.Model
	Name  string
	Test9 Test9 `gorm:"polymorphic:Owner"`
}

type Test8 struct {
	gorm.Model
	Name  string
	Test9 Test9 `gorm:"polymorphic:Owner"`
}
type Test9 struct {
	gorm.Model
	OwnerType string
	OwnerID   uint
}

func Polymorphic() {
	fmt.Println("多态")
}
