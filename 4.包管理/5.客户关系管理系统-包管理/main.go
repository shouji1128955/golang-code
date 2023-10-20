package main

import (
	. "cms/service"
	"fmt"
	"os"
)

// 实现一个客户关系管理系统
// 实现客户信息的增删改查以及保存

func main() {
	cs := CustomerService{}
	fs := FsDBservice{}

	//读取文件
	cs.Customers = fs.Read()
	for {
		//fmt.Println(Customers)
		fmt.Println(`
        1. 添加客户
        2. 查询所有客户
        3. 查询客户
        4. 修改客户信息
        5. 删除客户
        6. 保存
        7. 退出程序`)

		var choice int
		fmt.Print("请输入选择:")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			cs.AddCustomer()
		case 2:
			cs.ListCustomer()
		case 3:
			cs.GetOneCustomer()
		case 4:
			cs.UpdateCustomer()
		case 5:
			cs.DeleteCustomer()
		case 6:
			fs.Write(cs.Customers)
		case 7:
			os.Exit(200)
		default:
			fmt.Println("非法输入")
		}

	}

}
