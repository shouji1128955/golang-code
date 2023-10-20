package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) eat() {
	fmt.Println(a.name + "吃")
}

func (a Animal) sleep() {
	fmt.Println(a.name + "睡")
}

type Cat struct {
	Animal
}
type Dog struct {
	kind string
	Animal
}

func (d Dog) spark() {
	fmt.Println(d.name + "狂吠")
}
func main() {
	d1 := Dog{"金毛", Animal{"alex"}}
	fmt.Println(d1)
	fmt.Println(d1.name)
	fmt.Println(d1.kind)
	d1.eat()
	d1.spark()

	c1 := Cat{Animal{"喵喵"}}
	fmt.Println(c1.name)
	c1.eat()

}
