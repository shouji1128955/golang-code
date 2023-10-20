package main

import "fmt"

func main() {

	//切片实现
	stu01 := []interface{}{"zhang", 22, "180"}
	//重点 interface 通过空接口，实现对任务类型的匹配
	fmt.Println(stu01)

	//map实现
	/*	map的定义
		定义： map[KeyType]ValueType
		其中，KeyType表示键的类型，ValueType表示对应值的类型。

		map类型变量默认初始值为 nil，要使用make()函数分配内存。
		语法 ： make(map[KeyType]ValueType, [cap])
		其中，cap为map的容量，参数不是必须给的
	*/

	//第一种方式，先声明并初始化
	stu02 := make(map[string]string)
	stu02["name"] = "yuan"
	stu02["age"] = "32"
	fmt.Println(stu02)

	//第二种方式: 声明并初始化
	var stu03 = map[string]string{"name": "ilync", "age": "32", "gender": "male"}
	fmt.Println(stu03)
}
