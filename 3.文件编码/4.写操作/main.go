package main

// 写文件需要使用一个Openfile的方法 读和写底层都是使用openfile方法，使用的参数不同
import (
	"fmt"
	"os"
)

func writeBytesOrStr(file *os.File) {
	s := "怒发冲冠，凭阑处，㴋㴋雨歇。\n"

	//下面这两种方式非常重要，在网络编程也会用到
	file.Write([]byte(s)) //第一种方式 需要把字符串转为字节转才能写进去
	file.WriteString(s)   //第二种方式， 直接把字符串写入到文件
}

// writeFile用法，直接把字符串写入到文件
func writeFile(file *os.File) {
	data := `
		怒发冲冠，凭阑处，㴋㴋雨歇。
		抬望眼，仰天长啸，壮怀激烈。
		三十功名尘与土，八千里路云和月。
		莫等闲，白了少年头，空悲切。
		靖康耻，犹未雪；臣子恨，何时灭？
		驾长车，踏破贺兰山缺。
		壮志饥餐胡虏肉，笑谈渴饮匈奴血。
		待从头，收拾旧山河，朝天阙!
`
	os.WriteFile("满江红3.txt", []byte(data), 666)
}

// 实现读取满江红3.txt 写入到满江红4.txt
func readWritefile(file *os.File) {
	data, _ := os.ReadFile("满江红3.txt")
	os.WriteFile("满江红4.txt", []byte(data), 666)

}
func main() {

	//打开文件
	file, err := os.OpenFile("满江红.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 666)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
	//writeBytesOrStr(file)
	//writeFile(file)
	readWritefile(file)
}
