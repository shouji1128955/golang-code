package main

import (
	"fmt"
)

func main() {
	//(1)切片可以嵌套map,map也可以嵌套切片

	var stu01 = map[string]string{"name": "zhangsan", "age": "30", "gender": "male"}
	var stu02 = map[string]string{"name": "wangwu", "age": "22", "gender": "male"}
	var stu03 = map[string]string{"name": "zhaoliu", "age": "23", "gender": "male"}

	//切片管理
	stus := []map[string]string{stu01, stu02, stu03}
	fmt.Println(stus)
	//查询第二个学生的你年龄
	fmt.Println(stus[1]["name"])

	//(2)循环打印每一个学生的信息
	for i, stuMap := range stus {
		fmt.Printf("%d- 姓名:%s 年龄:%s 性别:%s \n", i, stuMap["name"], stuMap["age"], stuMap["gender"])
	}

	//(3)添加学生
	//思路，先把输入的放到一个新的map中，然后在把这个新的map放到大的map中
	var name string
	fmt.Println("请输入学生名字")
	fmt.Scan(&name)

	var age string
	fmt.Println("请输入学生年龄")
	fmt.Scan(&age)

	var gender string
	fmt.Println("请输入学生性别")
	fmt.Scan(&gender)

	var stunew = map[string]string{"name": name, "age": age, "gender": gender}
	stus = append(stus, stunew)
	fmt.Println(stus)
	//以上是切片套map

	/*	//map嵌套map
		var stu011 = map[string]string{"name": "zhangsan", "age": "30", "gender": "male"}
		var stu022 = map[string]string{"name": "wangwu", "age": "22", "gender": "male"}
		var stu033 = map[string]string{"name": "zhaoliu", "age": "23", "gender": "male"}

		var stuss = map[int]map[string]string{1001: stu011, 1002: stu022, 1003: stu033}
		//但是通过key来删除比较方便

		fmt.Println(stuss[1001]["name"])*/

	//map嵌套切片
	//片用的场景比较少
	var stu011 = map[string]interface{}{"name": "zhangsan", "age": "30", "gender": "male", "hobby": []string{"羽毛球", "篮球", "书法"}}
	var stuss = []map[string]interface{}{stu011}
	fmt.Println(stuss[0]["hobby"].([]string)[2]) //这里使用到了断言，因为用了空接口，在后续章节了解断言

}
