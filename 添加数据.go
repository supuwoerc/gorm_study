package main

import (
	"fmt"
	"strconv"
)

// CreateTest 创建一条数据
func CreateTest() {
	//选择特定字段 DB.Select()
	//剔除特定字段 DB.Omit()
	var users []*User
	for i := 0; i < 5; i++ {
		users = append(users, &User{
			Name: "测试数据" + strconv.Itoa(i+1),
			Age:  uint(10 + i),
		})
	}
	res := DB.Select("Name", "Age").Create(users)
	if res.Error != nil {
		fmt.Println("创建失败：" + res.Error.Error())
	} else {
		fmt.Println("创建成，影响行数：" + strconv.FormatInt(res.RowsAffected, 10))
	}
}
