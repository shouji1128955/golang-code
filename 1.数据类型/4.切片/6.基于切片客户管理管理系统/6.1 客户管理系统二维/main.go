package main

import "fmt"

func main() {
	var customers [][]string
	for {
		fmt.Println(`
		1.查看客户
		2.添加客户
		3.删除客户
		4.修改客户
		5.退出程序
`)
		var choise int
		fmt.Println("请输入你的选择")
		fmt.Scan(&choise)

		switch choise {
		case 1:

			fmt.Println("[查看客户]")
			for i, customerslice := range customers {
				//fmt.Println(i+1, customerslice)
				fmt.Printf("%d. 姓名: %s 年龄: %s 身高: %s", i+1, customerslice[0], customerslice[1], customerslice[2])
				fmt.Println("\n")
			}

		case 2:
			var customer []string
			fmt.Println("[添加客户]")
			var userName string
			fmt.Println("请输入用户的姓名")
			fmt.Scan(&userName)

			customer = append(customer, userName)

			var userAge string
			fmt.Println("请输入用户的年龄")
			fmt.Scan(&userAge)
			customer = append(customer, userAge)

			var userHeigth string
			fmt.Println("请输入用户的身高")
			fmt.Scan(&userHeigth)
			customer = append(customer, userHeigth)
			customers = append(customers, customer)
			fmt.Println(customer)

		case 3:
			fmt.Println("[删除客户]")
		case 4:
			fmt.Println("[修改客户]")
		case 5:
			fmt.Println("[退出程序]")
		}
	}
}
