package main

import "fmt"

func main() {
	var stuSlice = []string{"ilync", "32", "male"}
	var stuMap = map[string]string{"name": "wang", "age": "42", "gender": "male"}

	//查看键值
	fmt.Println(stuSlice[1])
	fmt.Println(stuMap["age"])
	fmt.Println(stuMap["gender"])

	//添加&更新键值对
	stuMap["weight"] = "70kg"
	stuMap["jiguan"] = "beijing"
	fmt.Println(stuMap)

	//删除键值
	delete(stuMap, "name")
	fmt.Println(stuMap)
}
