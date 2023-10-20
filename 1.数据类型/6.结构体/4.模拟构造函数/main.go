package main

//就是用来模拟在java项目中如何实现类和对象的代码编写的风格 ,不重要
import "fmt"

type Persion struct {
	//以下的属性称为成员变量，java中称为类变量
	name      string
	age       int
	isMarried bool
	scores    []float64
}

func NewPerson(name string, age int, isMarried bool, scores []float64) *Persion {
	return &Persion{name, age, isMarried, scores}
}

func main() {

	p1 := NewPerson("rain", 23, true, []float64{100, 99, 89})
	fmt.Println(p1)
}
