package main

import (
	"fmt"
)

// 注意: 在生产中，注意三个点，第一看函数名，见名知意，第二看参数，第三看返回值，如果还是看不懂，就去看需求文档
// func add(x, y int) [这里声明返回值]
// func add(x, y int) [int][string] 返回两个返回值
// 这种编译型语言会比较强调，输入的是什么类型，输出的是什么类型，必须得声明清楚，跟python解释性不一样
func add(x, y int) int {
	var ret = x + y
	return ret
}

// 计算器案例
func add2(oper string, nums ...int) int {
	var ret int
	switch oper {
	case "+":
		ret = 0
		for _, value := range nums {
			ret += value
		}
		//return (ret)
	case "*":
		ret = 1
		for _, value := range nums {
			ret *= value
		}
		//return (ret)
	}
	return (ret)
}

// 返回值命名
func bar() int {
	ret := 100
	return ret
}

func bar2() (ret int) {
	ret = 1000
	//return ret
	return
}

func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return //return  sum, sub
}
func main() {
	result := add(100, 200)
	fmt.Println(result)

	result2 := add2("*", 1, 2, 4, 6)
	fmt.Println(result2)
}

//面试题: 返回值命名
