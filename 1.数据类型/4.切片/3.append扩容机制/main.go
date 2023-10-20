package main

import "fmt"

func main() {

	//(1)有容量的 append
	emp1 := make([]string, 3, 6)
	emp1[0] = "beijing"
	emp1[1] = "shanghai"
	emp1[2] = "shengzhen"
	fmt.Println(emp1)

	emp2 := append(emp1, "hangzhou")
	fmt.Println(emp2)

	//(2)没有容量的append，如果空间不够，会进行两倍扩容
	var s = []int{1, 2, 3}
	fmt.Println(len(s), cap(s)) //3 3
	s2 := append(s, 4)
	fmt.Println(s2)
	fmt.Println(len(s2), cap(s2)) //4 6

	var a = make([]int, 3, 6)
	a2 := append(a, 8)
	fmt.Println(a)
	a[0] = 120
	fmt.Println(a2) //[120 0 0 8]

}
