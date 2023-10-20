package main

import "fmt"

func main() {

	//append的时候添加一个切片，需要使用到 ... ,需要打散添加
	var s []int
	t := []int{100, 200, 300}
	s1 := append(s, t...)
	fmt.Println(s1)

	//经典案例
	/*	a := []int{11, 22, 33}
		fmt.Println(len(a), cap(a))

		c := append(a, 55)
		a[0] = 100
		fmt.Println(a) //[100 22 33]
		fmt.Println(c) //[11 22 33 55]*/

	a := make([]int, 3, 10)
	fmt.Println(a)
	b := append(a, 11, 22)
	fmt.Println(a)
	fmt.Println(b)
	a[0] = 100
	fmt.Println(a)
	fmt.Println(b)

	//经典面试题，需要理解
	arr := [4]int{10, 20, 30, 40}
	s1 = arr[0:2] //[10 20]
	fmt.Println(s1)
	s2 := s1 // [10 20]
	s3 := append(append(append(s1, 1), 2), 3)
	s1[0] = 1000
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(arr)

}
