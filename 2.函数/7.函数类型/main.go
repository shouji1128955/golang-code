package main

//备注： 这部分内容非常重要，在go的并发，协程应用部分都会用到匿名函数来使用
import (
	"fmt"
	"reflect"
)

func foo(x, y int) {
	fmt.Println(x + y)
}

func bar(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(reflect.TypeOf(foo)) //拥有两个int形参的没有返回值的函数类型
	fmt.Println(reflect.TypeOf(bar)) //拥有两个int形参的有一个返回值的的函数类型

	//所以，[3]int{1,2,3} 跟 [4]int{1,2,3,4} 时不同的两个数据类型
	//一个是[3]int  一个是[4]int 完全是两个不同的数据类型

}
