package api //如果api问价夹下有很多个go文件,这里的包名建议都是跟问价夹名一样
import "fmt"

func RestAPI01() {
	fmt.Println("restAPI01....")
	foo()
}
func RestAPI02() {
	fmt.Println("restAPI01....")
}
func foo() {
	fmt.Println("foo")
}
func init() {
	fmt.Println("http初始化")
}
