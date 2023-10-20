package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 实现一个客户关系管理系统
// 实现客户信息的增删改查以及保存
type Customer struct {
	ID    int
	Name  string
	Age   int
	Email string
	Phone string
}
type CustomerService struct {
	Customers []Customer
}

func (cs *CustomerService) customerExistByID() (index, Id int) {

	fmt.Printf("请输入客户的ID:")
	fmt.Scan(&Id)
	index = -1
	for i, c := range cs.Customers {
		if c.ID == Id {
			index = i
		}
	}
	return
}

// 1)
func (cs *CustomerService) addCustomer() {
	fmt.Println(cs)
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
	newCustomer := Customer{len(cs.Customers), name, age, email, phone}
	cs.Customers = append(cs.Customers, newCustomer)
	fmt.Printf("[INFO]添加对象%s成功!", name)
}

// 2)
func (cs *CustomerService) listCustomer() {
	//fmt.Println(Customers)
	for _, customer := range cs.Customers {
		fmt.Printf("客户ID:%d 姓名: %s 年龄:%d 邮箱:%s 手机号码:%s\n", customer.ID, customer.Name, customer.Age, customer.Email, customer.Phone)
	}
}

// 3)
func (cs *CustomerService) GetOneCustomer() {
	index, Id := cs.customerExistByID()
	if index == -1 {
		fmt.Printf("没有查询到%d用户", Id)
	} else {
		fmt.Printf("ID:%d 客户姓名：%s 年龄：%d，邮箱：%s，电话：%s\n", cs.Customers[index].ID, cs.Customers[index].Name, cs.Customers[index].Age, cs.Customers[index].Email, cs.Customers[index].Phone)
	}
}

// 4)
func (cs *CustomerService) updateCustomer() {
	index, Id := cs.customerExistByID()
	if index == -1 {
		fmt.Printf("没有查询到%d用户", Id)
	} else {
		fmt.Printf("ID:%d 客户姓名：%s 年龄：%d，邮箱：%s，电话：%s\n", cs.Customers[index].ID, cs.Customers[index].Name, cs.Customers[index].Age, cs.Customers[index].Email, cs.Customers[index].Phone)

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
			cs.Customers[index].Name = name
		case 2:
			//修改年龄
			var age int
			fmt.Printf("请输入年龄:")
			fmt.Scan(&age)
			cs.Customers[index].Age = age
		case 3:
			var email string
			fmt.Printf("请输入邮件:")
			fmt.Scan(&email)
			cs.Customers[index].Email = email
		case 4:
			var phone string
			fmt.Printf("请输入电话号码:")
			fmt.Scan(&phone)
			cs.Customers[index].Phone = phone
		}
	}

}

// 5)
func (cs *CustomerService) deleteCustomer() {
	index, Id := cs.customerExistByID()
	if index == -1 {
		fmt.Printf("没有查询到%d用户", Id)
	} else {
		cs.Customers = append(cs.Customers[:index], cs.Customers[index+1:]...)
		fmt.Printf("成功删除%d用户!", Id)
	}

}

type FsDBservice struct {
}

// 6)
func (cs *FsDBservice) write(data []Customer) {
	//保存写入到文件
	jsonBytes, _ := json.Marshal(data)
	os.WriteFile("customers.json", jsonBytes, 666)
	fmt.Println("[INFO]文件保存成功!")
}

func (cs *FsDBservice) Read() []Customer {
	jsonByte, _ := os.ReadFile("customers.json")
	//进行序列化
	var data []Customer
	json.Unmarshal(jsonByte, &data)
	return data
	//jsonWrite, _ := json.Marshal(Customers)
	//os.WriteFile("customers.json", jsonWrite, 666)
}
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
			cs.addCustomer()
		case 2:
			cs.listCustomer()
		case 3:
			cs.GetOneCustomer()
		case 4:
			cs.updateCustomer()
		case 5:
			cs.deleteCustomer()
		case 6:
			fs.write(cs.Customers)
		case 7:
			os.Exit(200)
		default:
			fmt.Println("非法输入")
		}

	}

}
