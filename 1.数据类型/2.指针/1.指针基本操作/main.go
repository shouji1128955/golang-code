package main

import "fmt"

func main() {

	/*	//练习1
		fmt.Println("hello,world")
		fmt.Println("hi，shijie")
		var x = 10
		var y = &x //y是指针变量，相当于 var y *int
		var z = *y // 发生的就是值拷贝
		x = 20
		fmt.Println(x)  //x=10
		fmt.Println(*y) //y指针变量对应的值
		fmt.Println(z)  //z=10  值拷贝
		//创建一个redis客户端
	*/

	//练习2
	var a = 100
	var b = &a // var b *int  --> b是指针变量地址 然后取a的地址赋值给b
	var c = &b // var c *int --> c是存储b的指针地址
	**c = 200  // var c **int (*int)表示找到c的存储值 (**int)表示找到周到b的存储值
	fmt.Println(a)

}
