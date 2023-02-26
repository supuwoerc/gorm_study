package main

func UpdateTest() {
	//update 只更新选择的字段
	//DB.Model(&User{}).Where("name = ?", "测试数据1").Update("age", 88)
	//updates 更新全部字段，结构体作为参数时零值不参与更新，map作为参数时零值参与更新
	//var user User
	//DB.Model(&User{}).Where("name = ?", "测试数据1").First(&user).Updates(&User{
	//	Name: "名字会被更新，年龄不会",
	//	Age:  0,
	//})
	DB.Model(&User{}).Where("name = ?", "测试数据2").Updates(map[string]interface{}{
		"Name": "名字和年龄都被更新",
		"Age":  0,
	})
	//save 全部保存，包括零值
	//var user User
	//DB.Model(&User{}).Where("name = ?", "测试数据1").First(&user)
	//user.Age = 99
	//DB.Save(&user)
}
