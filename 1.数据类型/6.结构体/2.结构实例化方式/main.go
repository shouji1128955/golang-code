package main

import (
	"fmt"
	"reflect"
)

//声明Persion结构体，或者Person类型

// 结构体的声明
type Persion struct {
	//以下的属性称为成员变量，java中称为类变量
	name      string
	age       int
	isMarried bool
	scores    []float64
}

// 验证，通过下面这种方式修改值不生效，因为是值拷贝
func ageInit(p Persion) {
	p.age = 0
}

func ageInit2(p *Persion) {
	p.age = 0
}

func main() {

	//结构体实例化方式2, 直接实例化-按顺序来
	var p2 = Persion{"rain", 23, true, []float64{100, 99, 89}}
	fmt.Println(p2)

	//只想赋值个别属性
	var p3 = Persion{name: "ilync", age: 28, isMarried: false}
	fmt.Println(p3)

	//第三种实例化方式-实例化取地址 就是第二种的变种
	var p4 = &Persion{name: "zhaoliu", age: 33, isMarried: false}
	fmt.Println(p4)
	ageInit(p2)
	fmt.Println(p2.age)
	fmt.Println(reflect.TypeOf(p4))
	fmt.Println(*p4)
	fmt.Println((*p4).name)
	fmt.Println((*p4).age)

	//语法糖
	fmt.Println(p4.name) //这里go的编译器首先会判断p4是不是一个指针变量,如果判断什么类型的变量，然后自动添加* 单仅限于结构体变量

	ageInit2(p4)
	fmt.Println(p4.age)
}
