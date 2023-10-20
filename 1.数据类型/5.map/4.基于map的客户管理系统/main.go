package main

import (
	"fmt"
	"os"
)

func main() {
	//切片中嵌套map来实现
	var customers []map[string]string
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
			for i, customerMap := range customers {
				fmt.Println(i+1, customerMap)
				fmt.Printf("姓名: %s 年龄: %s 身高: %s\n", customerMap["name"], customerMap["age"], customerMap["height"])
			}

		case 2:
			var customerMap map[string]string
			fmt.Println("【添加用户】")
			var name string
			fmt.Println("请输入客户名字")
			fmt.Scan(&name)

			var age string
			fmt.Println("请输入客户年龄")
			fmt.Scan(&age)

			var height string
			fmt.Println("请输入客户身高")
			fmt.Scan(&height)

			customerMap = map[string]string{"name": name, "age": age, "height": height}
			customers = append(customers, customerMap)

		case 3:
			fmt.Println("【删除客户】")
			fmt.Println("请输入需要删除的用户")
			var customerMapuser string
			fmt.Scan(&customerMapuser)
			for i, customerMap := range customers {
				fmt.Println(i+1, customerMap)
				fmt.Printf("姓名: %s 年龄: %s 身高: %s\n", customerMap["name"], customerMap["age"], customerMap["height"])
				if customerMap["name"] == customerMapuser {
					customers = append(customers[:i], customers[i+1:]...)
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
	}

}
