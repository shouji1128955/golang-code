package main

import "fmt"

func main() {
	//基本数据类型都是值类型
	var x int
	fmt.Println(x)
	x = 10
	//整型的默认值就是0
	var b bool
	fmt.Println("b", b)

	var y = 10
	//声明指针变量
	var p *int
	p = &y
	fmt.Println(p)

	var c *int     //也算是开了空间，就是空间中没有创建地址
	fmt.Println(c) // <nil> 空对象
	c = new(int)   //在空间中创建地址，以及默认值为0
	*c = 120
	fmt.Println(c)
}
