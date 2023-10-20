package main

import "fmt"

// 方式1
func add(x, y int) {
	fmt.Println(x + y)
}

// 方式2
func add2(x int, y int) {
	fmt.Println(x + y)
}

// 方式3,不定长,不固定,可伸缩
func add3(nums ...int) {
	ret := 0
	for _, num := range nums {
		ret += num
	}
	fmt.Println(ret)
}

// // 不定长变量一定要放在单个变量后面
func cal(oper string, nums ...int) {
	//上面oper不能在nums后面，否则会把所有的数据传递给nums
	//go语法没有默认参数，涉及理念，不支持
	ret := 0
	switch oper {

	case "+":
		//累计逻辑
		for _, num := range nums {
			ret += num
		}
	case "*":
		//累乘逻辑
		ret := 1
		for _, num := range nums {
			ret *= num
		}
	}
	fmt.Println(ret)
}

func main() {
	add(100, 200)
	add2(120, 120)
	add3(1, 12, 99)
	cal("+", 1, 2, 3, 5)
}
