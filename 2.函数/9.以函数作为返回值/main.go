package main

import "fmt"

func bar() func(x, y int) int {
	var add = func(x, y int) int {
		return x + y
	}
	return add
}
func main() {

	f := bar()
	ret := f(1, 2)
	fmt.Println(ret)
}
