package main

import (
	"fmt"
	"os"
)

func main() {
	//基于切片的客户管理系统
	var customers []string
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
			for i, v := range customers {
				fmt.Println(i+1, v)
			}

		case 2:
			fmt.Println("【添加用户】")
			var addName string
			fmt.Println("请输入客户名字")
			fmt.Scan(&addName)
			customers = append(customers, addName)

		case 3:
			fmt.Println("【删除客户】")
			fmt.Println("请输入需要删除的用户")
			var delName string
			fmt.Scan(&delName)
			for i, v := range customers {
				if v == delName {
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
