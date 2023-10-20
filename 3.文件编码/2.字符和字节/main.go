package main

import "fmt"

func main() {
	//存储字符， 使用单引号，但是需要注意类型 byte
	var s byte
	s = 'a'
	fmt.Println(s) //这个值就是ascci表对应的值
	//声明byte跟uint8是一样的,用uint8是当作数字来看待
	//转换为字符串
	fmt.Println(string(s))

	//第二种方式,使用uint8
	var s2 uint8
	s2 = 65
	fmt.Println(s2)
	fmt.Println(string(s2))

	//如何存储中文
	//var b3 byte //定义存储为一个字节
	//b3 = "张"   //默认是不在ascci表里
	var s3 rune //rune == int32 四个字节
	s3 = '张'
	fmt.Println(s3)            //24352
	fmt.Println(string(24352)) //张

}
