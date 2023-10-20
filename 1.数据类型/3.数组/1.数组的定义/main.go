package main

import "fmt"

//开发时候经常需要用到切片，熟悉数组，熟练切片
//var 数组名  [元素数量]元素类型

func main() {

	//数组中元素有顺序，长度一旦定下来，不可伸缩，类型不能混用
	//特点： 1. 一致性-元素类型必须相同， 有序性， 不可变性-长度定下来，不可变

	//(1)先声明，在赋值
	var student [5]string
	student[0] = "beijing"
	student[1] = "shanghai"
	student[2] = "guanzhou"
	student[3] = "shengzhen"
	student[4] = "xian"

	//(2)赋值第二种方式 声明并且赋值
	var stusNames = [3]string{"张三", "李四", "王五"}
	fmt.Println(student)
	fmt.Println(stusNames)

	//(3)省略长度，如果数组长度太长，可以直接使用...
	//arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr)

	//数组取值
	var couse = [3]string{"语文", "数学", "英语"}
	fmt.Println(&couse[0])
	fmt.Println(&couse[1])
	fmt.Println(&couse[2]) //一个字符串占16字节

	/*	//输出以下的结果 0x表示16进制
			0xc0000904e0
			0xc0000904f0
			0xc000090500
			因为是逢16进一，从1开始，16就是f*
		   以上因为用的是string类型，所以是16个字节，如果定义用的是占用8个字节。那定义如下
		   var stusAgets = [3]int64{1,2,3}
		/
	*/

	//演示定义使用8字节
	var StusAges = [3]int64{1, 2, 3}
	fmt.Println(&StusAges[0])
	fmt.Println(&StusAges[1])
	fmt.Println(&StusAges[2])

	/*	执行结果如下
			0xc0001104e0
			0xc0001104f0
			0xc000110500
			0xc000014138
			0xc000014140
			0xc000014148
		所以以上验证，16位的地址具体以多少位进制决定主要还是在于类型
	*/

}
