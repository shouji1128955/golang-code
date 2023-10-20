package main

import "fmt"

func main() {
	//(1)首位添加一个值
	var s = []int{1, 2, 3, 4, 5, 6}

	//比如在s 切片最前面添加1000, append传入的第一位必须是切片
	s1 := append([]int{1000}, s...)
	fmt.Println(s1)

	//(2)任意位添加
	var a = []int{10, 20, 30, 40, 50}
	temp := append([]int{1000}, a[2:]...)
	fmt.Println(temp)
	fmt.Println(append(a[2:], temp...))

	//(3)删除某索引值-比如删除300
	//第一中方式，对元素据进行了修改
	/*	var d = []int{100, 200, 300, 400, 500}
		ret := append(d[0:2], d[3:]...)
		fmt.Println(ret)*/

	/*	//第二种方式，就是构建一个新数组，把不满足条件的进行添加到新数组，然后达到删除数组元素的目标
		var d = []int{10, 20, 30, 40, 50, 60}
		var deleteValue = 30
		var d2 []int
		for _, v := range d {
			if v != deleteValue {
				d2 = append(d2, v)
			}
		}
		fmt.Println(d2)
	*/

	//第三种方式

	var d = []int{10, 20, 30, 40, 50, 60}
	var deleteValue = 30
	for i, v := range d {
		if v == deleteValue {
			d = append(d[:i], d[i+1:]...)
		}
	}
	fmt.Println(d)

}
