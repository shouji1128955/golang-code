package main

import "fmt"

var x = 100 //这个地方绝对不能使用 x := 100 ,这种用法只能在局部做变量

func foo() {
	var x = 1000
	fmt.Println(x)
}
func main() {

	//理解,变量得作用域
	//变量得生命周期
	foo()
	//x = 200
	fmt.Println(x) //优先级:  局部 >  全局

}

//定义全局定义域,函数之外的
