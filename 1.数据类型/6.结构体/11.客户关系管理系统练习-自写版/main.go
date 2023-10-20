package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 通过结构体实现客户关系管理系统并且存储到文件，实现序列化和饭序列化

type student struct {
	Name   string
	Age    int
	Height string
}

func main() {

	//var stuStruct = student{"zhangsan", 31, "172"}

	//刚开始进行反序列化
	var customers []student
	jsonBytes, _ := os.ReadFile("custermanager.json")
	fmt.Println(string(jsonBytes))
	var data []student
	json.Unmarshal(jsonBytes, &data)
	customers = data
	//customers = append(customers, stuStruct)
	for {

		fmt.Println("==========================")
		fmt.Println(`
		1.查看客户
		2.添加客户
		3.删除客户
		4.修改客户
		5.退出程序
`)

		var choice int
		fmt.Println("请输入您的选择:")
		fmt.Scan(&choice)
		fmt.Println()
		switch choice {
		case 1:
			fmt.Println("【查看用户】")
			fmt.Println()
			for i, customerSice := range customers {
				fmt.Printf("序号:%d 姓名: %s 年龄: %d, 身高: %scm\n", i+1, customerSice.Name, customerSice.Age, customerSice.Height)

			}

		case 2:
			var customerSice student
			fmt.Println("【添加用户】")
			var name string
			fmt.Println("请输入客户名字")
			fmt.Scan(&name)

			var age int
			fmt.Println("请输入客户年龄")
			fmt.Scan(&age)

			var height string
			fmt.Println("请输入客户身高")
			fmt.Scan(&height)

			customerSice = student{name, age, height}
			customers = append(customers, customerSice)

			//使用json序列化存入到json文件中
			jsonByte, err := json.Marshal(customers)
			//如果想查看内容
			fmt.Println(err)
			fmt.Println(string(jsonByte))
			//然后写入到文件
			os.WriteFile("custermanager.json", jsonByte, 666)
			fmt.Println(customers)
		case 3:
			fmt.Println("【删除客户】")
			fmt.Println("请输入需要删除的用户")
			var customername string
			fmt.Scan(&customername)
			for i, customerSice := range customers {
				fmt.Printf("姓名: %s 年龄: %s 身高: %s\n", customerSice.Name, customerSice.Age, customerSice.Height)
				if customerSice.Name == customername {
					customers = append(customers[:i], customers[i+1:]...)
					fmt.Println("已经删除用户:", customerSice.Name)
				}

			}

		case 4:
			fmt.Println("【修改程序】")
		case 5:
			fmt.Println("【退出程序】")
			os.Exit(200)
		default:
			fmt.Println("非法输入")
		}
		//通过切片中套用结构体来实现

	}

}
