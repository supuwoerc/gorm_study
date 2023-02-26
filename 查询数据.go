package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// FindTest 查询数据：https://gorm.io/zh_CN/docs/query.html
func FindTest() {
	//采用map接收返回的数据
	//var result map[string]interface{}
	//采用结构体接收返回的数据
	var result User
	//First 按照主键排序，查询到第一个（需要在语句中体现model，不然gorm不知道主键）
	//Last 按照主键排序，查询到最后一个 (需要在语句中体现model，不然gorm不知道主键)
	//Take  查询到表里面第一条数据 (不需要需要在语句中体现model，只需要告知表名就行：model或者直接Table("表名"))
	//err := DB.Model(&User{}).First(&result, 99).Error
	//Where添加查询条件：拼接查询条件
	//err := DB.Where("name = ?", "测试数据1").First(&result).Error
	//Where添加查询条件：使用结构体作为条件查询
	//err := DB.Where(User{Name: "测试数据2"}).First(&result).Error
	//Where添加查询条件：使用map作为查询条件
	//err := DB.Where(map[string]interface{}{"name": "测试数据3"}).First(&result).Error
	//Or 添加多个查询条件
	//err := DB.Where(map[string]interface{}{"name": "测试数据3-notin"}).Or("name = ?", "测试数据2").First(&result).Error
	//Not 排除条件
	//err := DB.Where(map[string]interface{}{"name": "测试数据3"}).Not("name = ?", "测试数据2").First(&result).Error
	//Order 排序
	//err := DB.Order("name desc").Take(&result).Error
	//limit&offset添加偏移和结果数量
	//err := DB.Model(&User{}).Offset(2).Limit(1).Find(&result).Error
	//err := DB.Model(&User{}).First(&result, "name =  ?", "测试数据3").Error
	//err := DB.Model(&User{}).First(&result, &User{Name: "测试数据4"}).Error
	err := DB.Model(&User{}).First(&result, map[string]interface{}{"name": "测试数据5"}).Error
	//type CustomResult struct {
	//	Name string
	//}
	//var customResult CustomResult
	//err := DB.Model(&User{}).First(&customResult).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("未查询到数据")
	}
	fmt.Println(result)
}
