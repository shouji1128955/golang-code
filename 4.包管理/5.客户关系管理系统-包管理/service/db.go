package service

import (
	"cms/config"
	. "cms/model"
	"encoding/json"
	"fmt"
	"os"
)

type FsDBservice struct {
}

// 6)
func (cs *FsDBservice) Write(data []Customer) {
	//保存写入到文件
	jsonBytes, _ := json.Marshal(data)
	os.WriteFile(config.JsonPath, jsonBytes, 666)
	fmt.Println("[INFO]文件保存成功!")
}

func (cs *FsDBservice) Read() []Customer {
	jsonByte, _ := os.ReadFile(config.JsonPath)
	//进行序列化
	var data []Customer
	json.Unmarshal(jsonByte, &data)
	return data
	//jsonWrite, _ := json.Marshal(Customers)
	//os.WriteFile("customers.json", jsonWrite, 666)
}
