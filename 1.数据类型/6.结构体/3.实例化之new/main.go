package main

import (
	"fmt"
)

// new方式也可以实例化结构体
type person struct {
	Name string
	Age  int
}

func main() {

	var p2 = new(person)
	p2.Name = "ilync"
	p2.Age = 28
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}

}
