package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Addr struct {
	Provine string
	City    string
	Zhen    string
}
type Customer struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr
}

func main() {

	//反序列化,就是把第9节的序列化创建的文件读取到内存中并且初始化
	jsonBytes, _ := os.ReadFile("customers.json")
	fmt.Println(string(jsonBytes))
	//var customers []map[string]interface{} //不确定类型可以使用空接口类型
	var customers []Customer
	json.Unmarshal(jsonBytes, &customers) //需要把地址传给定义好的变量进行赋值

	fmt.Println(customers)
	fmt.Println(customers[1].Addr.Provine)

	for i, customerstruct := range customers {
		fmt.Printf("序号:%d,姓名:%s 年龄:%d 省份:%s 城市:%s\n", i, customerstruct.Name, customerstruct.Age, customerstruct.Addr.Provine, customerstruct.City)

	}
}
