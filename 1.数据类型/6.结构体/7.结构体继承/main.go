package main

import "fmt"

type Cat struct {
	name string
}
type Dog struct {
	name string
	kind string
}

func (c Cat) eat() {
	fmt.Println(c.name + "吃")
}
func (c Cat) sleep() {
	fmt.Println(c.name + "睡")
}

func (d Dog) eat() {
	fmt.Println(d.name + "吃")
}
func (d Dog) sleep() {
	fmt.Println(d.name + "睡")
}

func main() {

	//说明，在以上的代码中，明显有很多的冗余，所以在结构体继承2中使用了优化

	Alex := Dog{"alex", "金毛"}
	fmt.Println(Alex.name)
	Alex.eat()
	Alex.sleep()

	c := Cat{"喵喵"}
	fmt.Println(c)
	fmt.Println(c.name)
	c.eat()
	c.sleep()
}
