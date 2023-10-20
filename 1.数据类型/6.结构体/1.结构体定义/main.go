package main

import "fmt"

//声明Persion结构体，或者Person类型

type Persion struct {
	//以下的属性称为成员变量，java中称为类变量
	name      string
	age       int
	isMarried bool
	//scores    []float64
	scores map[string]float64
}

func main() {

	var p1 = Persion{}
	fmt.Println(p1) //{ 0 false []}
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	//验证存储地址的连续性
	fmt.Printf("%p\n", &p1.name)
	fmt.Printf("%p\n", &p1.age)
	fmt.Printf("%p\n", &p1.isMarried)
	fmt.Printf("%p\n", &p1.scores)
	/*	0xc00006c040 //string 占16个字节，逢16进一
		0xc00006c050
		0xc00006c058
		0xc00006c060
	*/
	//给成员变量赋值
	p1.name = "ilync"
	p1.age = 22
	p1.isMarried = false
	p1.scores = map[string]float64{"chinese": 100, "english": 100, "match": 98}
	fmt.Println(p1)

}
