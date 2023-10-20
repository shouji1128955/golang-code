package main

import "fmt"

func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
	fmt.Println(*p1, *p2)
}
func main() {
	//实现对两个数的互换
	var x = 10
	var y = 20

	swap(&x, &y)
	//fmt.Println("x", x)
	//fmt.Println("y", y)

}
