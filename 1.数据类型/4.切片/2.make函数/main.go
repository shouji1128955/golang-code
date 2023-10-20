package main

import "fmt"

func main() {
	var x [3]int
	fmt.Println(x)

	var y []int //这样定义没有底层数组，所以没有索引，没法赋值,但是对于切片，可以追加方式添加，或者通过make初始化
	fmt.Println(y)
	yy := append(y, 123)
	fmt.Println("yy:", yy)

	//应该定义如何
	var z = make([]int, 3, 6) //三个值，但是有6个空间
	fmt.Println(z)            //打印后应该有三个0

	//案例-
	a := make([]int, 5)
	b := a[0:3] //b的初始值为0，长度为3，容量从初始化算，为6
	a[0] = 100
	fmt.Println(a)
	fmt.Println(b)

	//通过make进行初始化
	var test = make([]int, 2)
	test[0] = 1112
	fmt.Println(test)

}
