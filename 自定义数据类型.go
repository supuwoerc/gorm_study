package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type StringArray []string
type CustomJsonInfo struct {
	Name string
	Age  int
}

func (c CustomJsonInfo) Value() (driver.Value, error) {
	res, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return string(res), nil
}
func (c *CustomJsonInfo) Scan(value interface{}) error {
	bytes := value.([]byte)
	err := json.Unmarshal(bytes, c)
	if err != nil {
		return err
	}
	return nil
}
func (c StringArray) Value() (driver.Value, error) {
	str := strings.Join(c, ",")
	return str, nil
}
func (c *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("类型转换失败")
	}
	res := string(bytes)
	*c = strings.Split(res, ",")
	return nil
}

type CustomJson struct {
	ID             uint
	CustomJsonInfo CustomJsonInfo
	Strs           StringArray
}

func CustomSchema() {
	//CreateTableByModel(&CustomJson{})
	//DB.Create(&CustomJson{
	//	CustomJsonInfo: CustomJsonInfo{
	//		Name: "hhhh",
	//		Age:  18,
	//	},
	//	Strs: []string{"x", "y", "z"},
	//})
	var c CustomJson
	DB.First(&c)
	fmt.Println(c)
}
