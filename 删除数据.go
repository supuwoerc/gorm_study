package main

func DeleteTest() {
	//软删除
	//DB.Model(&User{}).Where("name = ?", "测试数据5").Delete(&User{})
	//硬删除
	DB.Unscoped().Model(&User{}).Where("name = ?", "测试数据5").Delete(&User{})
}
