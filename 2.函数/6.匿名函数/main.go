package main

//匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:
//在Go语言中，不需要函数名的函数，称为匿名函数
//备注： 这部分内容非常重要，在go的并发，协程应用部分都会用到匿名函数来使用
//Go支持匿名函数，如果我们某个函数只是希望使用一次，可以考虑使用匿名函数，匿名函数也可以实现多次调用。
import (
	"fmt"
)

func foo() {
	fmt.Println("hello")

}
func do(action func()) {
	action()
}
func get_action() func() {
	return func() {
		fmt.Println("result func")
	}
}
func main() {
	//foo()

	/*	//方式1 声明新函数
		(func(x, y int) {
			fmt.Println(x + y)
		})(10, 20)
		//方式2
		var add = func(x, y int) int {
			c := (x + y)
			return c
		}
		aa := add(101, 202)
		fmt.Println("res=", aa)
	*/

	/*	//匿名函数可以有入参
		func(str string) {
			fmt.Println(str)
		}("hello")*/

	//匿名函数可以作为值赋给一个变量
	/*
		value := func() {
			fmt.Println("hello")
		}
		value()*/

	/*	//匿名函数可以作为函数的入参和返回值

		say_hello := func() {
			fmt.Println("sayhello")
		}
		do(say_hello)*/

	//匿名函数可以称为另一个函数的返回值
	action := get_action()
	action()
}
