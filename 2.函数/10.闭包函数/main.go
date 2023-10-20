package main

import "fmt"

// 高阶函数-比较难理解，需要把基础学的很好，但是在学习第一阶段不是很重要，在中高级很重要
// 学习一门语言，但是需要知道，这个在js中用的非常多

func getCounter() func() {
	var i = 0
	// 我们希望i跟counter能够绑定，不需要在全局中配置，避免其他的变量引用
	fmt.Println("正在打印")
	counter := func() {
		i++
		fmt.Println("i=", i)
	}
	fmt.Println(counter)
	return counter
}
func main() {
	counter := getCounter()
	counter()
	counter()
	counter()
}
