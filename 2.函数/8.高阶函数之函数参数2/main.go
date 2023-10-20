package main

import (
	"fmt"
	"time"
)

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

func timer(f func()) {
	start := time.Now().Unix()
	f()
	end := time.Now().Unix()
	fmt.Println("cost time", end-start)
}

func main() {
	timer(foo)
	timer(bar)
}
