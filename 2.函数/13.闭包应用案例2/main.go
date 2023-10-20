package main

import (
	"fmt"
	"time"
)

func add(x, y int) int {
	return x + y
}

// f版本1
//foo是高阶函数，它以函数作为参数
/*func foo() {
	start := time.Now().Unix()
	fmt.Println("foo功能开始")
	time.Sleep(3 * time.Second)
	fmt.Println("foo功能结束")
	end := time.Now().Unix()
	fmt.Println("cost timer:", end-start)

}

func bar() {
	start := time.Now().Unix()
	fmt.Println("bar功能开始")
	time.Sleep(2 * time.Second)
	fmt.Println("bar功能结束")
	end := time.Now().Unix()
	fmt.Println("cost timer:", end-start)

}
*/

func foo() {

	fmt.Println("foo功能开始")
	time.Sleep(3 * time.Second)
	fmt.Println("foo功能结束")

}

func bar() {

	fmt.Println("bar功能开始")
	time.Sleep(2 * time.Second)
	fmt.Println("bar功能结束")

}

func getTimer(f func()) func() {
	return func() {
		start := time.Now().Unix()
		f()
		end := time.Now().Unix()
		fmt.Println("cost timer:", end-start)
	}
}

func main() {

	//timer(foo)
	//timer(bar)
	bar := getTimer(bar)
	bar()

	foo := getTimer(foo)
	foo()
}
