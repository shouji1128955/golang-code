package service

import (
	. "cms/model"
	"fmt"
)

type CustomerService struct {
	Customers []Customer
}

func (cs *CustomerService) CustomerExistByID() (index, Id int) {

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
func (cs *CustomerService) AddCustomer() {
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
func (cs *CustomerService) ListCustomer() {
	//fmt.Println(Customers)
	for _, customer := range cs.Customers {
		fmt.Printf("客户ID:%d 姓名: %s 年龄:%d 邮箱:%s 手机号码:%s\n", customer.ID, customer.Name, customer.Age, customer.Email, customer.Phone)
	}
}

// 3)
func (cs *CustomerService) GetOneCustomer() {
	index, Id := cs.CustomerExistByID()
	if index == -1 {
		fmt.Printf("没有查询到%d用户", Id)
	} else {
		fmt.Printf("ID:%d 客户姓名：%s 年龄：%d，邮箱：%s，电话：%s\n", cs.Customers[index].ID, cs.Customers[index].Name, cs.Customers[index].Age, cs.Customers[index].Email, cs.Customers[index].Phone)
	}
}

// 4)
func (cs *CustomerService) UpdateCustomer() {
	index, Id := cs.CustomerExistByID()
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
func (cs *CustomerService) DeleteCustomer() {
	index, Id := cs.CustomerExistByID()
	if index == -1 {
		fmt.Printf("没有查询到%d用户", Id)
	} else {
		cs.Customers = append(cs.Customers[:index], cs.Customers[index+1:]...)
		fmt.Printf("成功删除%d用户!", Id)
	}

}
