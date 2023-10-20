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
	c1 := Customer{"ilync", 23, Addr{"北京市", "昌平", "沙河镇"}}
	c2 := Customer{"wangsu", 23, Addr{"北京市", "昌平", "沙河镇"}}
	c3 := Customer{"rain", 23, Addr{"北京市", "昌平", "沙河镇"}}
	customers := []Customer{c1, c2, c3}
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println("============")
	fmt.Println(customers)
	//json 序列化
	jsonByte, _ := json.Marshal(customers)
	fmt.Println(string(jsonByte))

	//然后写入到文件
	os.WriteFile("customers.json", jsonByte, 666)
}

//以上注意两个点
//1.对外的变量属性需要使用大写,否则获取不到
//2.如果对于json序列化的时候大写不友好，可以通过label标签来改写`json:"name"`
