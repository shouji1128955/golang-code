package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var Customers []Customer

// 实现一个客户关系管理系统
// 实现客户信息的增删改查以及保存
type Customer struct {
	ID    int
	Name  string
	Age   int
	Email string
	Phone string
}

func customerExistByID() (index, Id int) {

	fmt.Printf("请输入客户的ID:")
	fmt.Scan(&Id)
	index = -1
	for i, c := range Customers {
		if c.ID == Id {
			index = i
		}
	}
	return
}

// 1)
func addCustomer() {
	var name string
	var age int
	var email string
	var phone string
	fmt.Printf("请输入客户姓名:")
	fmt.Scan(&name)
	fmt.Printf("请输入客户年龄:")
	fmt.Scan(&age)
	fmt.Printf("请输入客户邮箱:")
	fmt.Scan(&email)
	fmt.Printf("请输入客户电话号码:")
	fmt.Scan(&phone)
	//实例Customer对象
	newCustomer := Customer{len(Customers), name, age, email, phone}
	Customers = append(Customers, newCustomer)
	fmt.Printf("[INFO]添加对象%s成功!", name)
}

// 2)
func listCustomer() {
	//fmt.Println(Customers)
	for _, customer := range Customers {
		fmt.Printf("客户ID:%d 姓名: %s 年龄:%d 邮箱:%s 手机号码:%s\n", customer.ID, customer.Name, customer.Age, customer.Email, customer.Phone)
	}
}

// 3)
func GetOneCustomer() {
	index, Id := customerExistByID()
	if index == -1 {
		fmt.Printf("没有查询到%d用户", Id)
	} else {
		fmt.Printf("ID:%d 客户姓名：%s 年龄：%d，邮箱：%s，电话：%s\n", Customers[index].ID, Customers[index].Name, Customers[index].Age, Customers[index].Email, Customers[index].Phone)
	}
}

// 4)
func updateCustomer() {
	index, Id := customerExistByID()
	if index == -1 {
		fmt.Printf("没有查询到%d用户", Id)
	} else {
		fmt.Printf("ID:%d 客户姓名：%s 年龄：%d，邮箱：%s，电话：%s\n", Customers[index].ID, Customers[index].Name, Customers[index].Age, Customers[index].Email, Customers[index].Phone)

		fmt.Println(`
             修改属性:
                  1. 姓名
                  2. 年龄
                  3. email
                  4. phone`)
		var attr int
		fmt.Printf("请输入修改属性的序号:")
		fmt.Scan(&attr)
		switch attr {
		case 1:
			var name string
			fmt.Printf("请输入姓名:")
			fmt.Scan(&name)
			Customers[index].Name = name
		case 2:
			//修改年龄
			var age int
			fmt.Printf("请输入年龄:")
			fmt.Scan(&age)
			Customers[index].Age = age
		case 3:
			var email string
			fmt.Printf("请输入邮件:")
			fmt.Scan(&email)
			Customers[index].Email = email
		case 4:
			var phone string
			fmt.Printf("请输入电话号码:")
			fmt.Scan(&phone)
			Customers[index].Phone = phone
		}
	}

}

// 5)
func deleteCustomer() {
	index, Id := customerExistByID()
	if index == -1 {
		fmt.Printf("没有查询到%d用户", Id)
	} else {
		Customers = append(Customers[:index], Customers[index+1:]...)
		fmt.Printf("成功删除%d用户!", Id)
	}

}

// 6)
func save() {
	//保存写入到文件
	jsonBytes, _ := json.Marshal(Customers)
	os.WriteFile("customers.json", jsonBytes, 666)
	fmt.Println("[INFO]文件保存成功!")
}

func initDB() {
	jsonByte, _ := os.ReadFile("customers.json")
	//进行序列化
	var data []Customer
	json.Unmarshal(jsonByte, &data)
	Customers = data
	//jsonWrite, _ := json.Marshal(Customers)
	//os.WriteFile("customers.json", jsonWrite, 666)
}
func main() {
	initDB()
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
			addCustomer()
		case 2:
			listCustomer()
		case 3:
			GetOneCustomer()
		case 4:
			updateCustomer()
		case 5:
			deleteCustomer()
		case 6:
			save()
		case 7:
			os.Exit(200)
		default:
			fmt.Println("非法输入")
		}

	}

}
