package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4, 5, 6}
	var s1 = s
	fmt.Println(s1)
	fmt.Println("%p", &s1[0])
	fmt.Println("%p", &s[0])
	//%p 0xc00000e300
	//%p 0xc00000e300

	//如何改变以上两个地址都不是同一个地址呢 ?

	var s2 = []int{10, 20, 30, 40, 50}
	var s3 = make([]int, 5, 5)
	copy(s3, s2) //可以通过copy来实现拷贝，把右边的拷贝到左边
	fmt.Println(s3)
	fmt.Println("%p", &s2[0])
	fmt.Println("%p", &s3[0])
	//%p 0xc00000e330
	//%p 0xc00000e360

	//copy其他的用法
	var copy01 = []int{11, 22, 33, 44}
	var copy02 = []int{12, 13, 14, 15, 16, 19}
	copy(copy02, copy01)
	fmt.Println(copy02) //[11 22 33 44 16 19] 按照索引进行拷贝
}
