package main

import "fmt"

// 函数是一种代码的组织形式
// 函数声明: 不会执行内部的代码
func printLoop() {
	for i := 1; i <= 200; i++ {
		fmt.Println(i)
	}

}
func main() {
	//函数调用dd
	printLoop()
}
