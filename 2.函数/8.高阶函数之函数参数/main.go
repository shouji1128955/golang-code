package main

import "fmt"

// 相对重要，比比较难的一部分，python里称作闭包函数

//特点
// 1.将一个或者多个函数作为形参
// 2.返回一个函数作为其结果

func add(x, y int) int {
	return x + y
}

// foo是高阶函数，它以函数作为参数
func foo(f func(int, int) int) {
	ret := f(1, 2)
	fmt.Println(ret)
}
func main() {
	foo(add)
}
