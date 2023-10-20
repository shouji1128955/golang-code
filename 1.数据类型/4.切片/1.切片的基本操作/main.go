package main

import (
	"fmt"
	"sort"
)

func main() {
	//这一节介绍切片的构建,有两种方式
	//1.方式 1，通过数组切片来实现
	var arr = [5]int{11, 12, 13, 14, 15}
	//对数组切片得到切片对象
	s1 := arr[0:1]
	s2 := arr[1:4]
	fmt.Println(s1, s2)
	fmt.Printf("%T,%T", s1, s2)
	//修改值
	arr[2] = 130
	fmt.Println(s1, s2) // [12 130 14]
	//切片的存储原理,并非通过值拷贝实现，切片其实就是一个结构体，由三部分构成
	//指针(起始位置),长度, 容量

	//值类型(int,float,bool,string) 以及数组和struct都属于值类型，特点就是声明未赋值的时候都有默认值

	//引用类型- 指针,slice,map,chan,interface都是应勇类型，拷贝的时候对地址进行拷贝，指向公共的空间

	a := []int{10, 120, 30, 80, 50}
	sort.Ints(a)
	fmt.Println(a)

}
