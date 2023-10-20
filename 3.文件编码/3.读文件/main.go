package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 第一种方式，按照字节
// 读文件，需要掌握，比较重要，前面2节都是做辅助
// 读取文件需要使用readBytes，还有一个为readfile
func readBytes(file *os.File) {
	data := make([]byte, 6) //取出两个汉字，一个汉字占3个字节
	file.Read(data)
	fmt.Println(data)
	fmt.Println(string(data)) //怒发
}

// 可以一行一行去读
func readByline(file *os.File) {
	reader := bufio.NewReader(file) //阅读器
	data, _, _ := reader.ReadLine()
	data2, _, _ := reader.ReadLine()
	fmt.Println(string(data))
	fmt.Println(string(data2))
	//fmt.Println(reader.ReadLine())
}

func readString(file *os.File) {
	reader := bufio.NewReader(file)
	//1)下面实现的方式就是一行一行去读
	//line, _ := reader.ReadString('\n')
	//line2, _ := reader.ReadString('\n')
	//fmt.Print(line)
	//fmt.Print(line2)

	//2)如何实现一行一行全部读出来
	for true {
		line, err := reader.ReadString('\n')
		fmt.Print(line)
		if err == io.EOF {
			break
		}
	}
}

func readfile(file *os.File) {
	data, _ := os.ReadFile("满江红.txt")
	fmt.Print(string(data))
}

func main() {
	//打开文件,需要用到os 满江红
	file, err := os.Open("满江红.txt")
	fmt.Println(file) //如果返回nil ,说明报错了
	if err != nil {
		fmt.Println("err: ", err)
	}
	//&{0xc000004a00} 正常返回类似这样的文件句柄

	//第一种读取方式，按照字节方式,使用的比较少
	//readBytes(file)
	//readByline(file)
	//readString(file)
	readfile(file) //读取整个文件,不适合大文件
}

/*从 Go 语言 1.16 开始，ioutil.ReadAll、ioutil.ReadFile 和 ioutil.ReadDir 被弃用，因为 io/ioutil 包被弃用。

解决方法如下，使用 io 或 os 包中相同的方法替换，即修改自己按如下参照修改包名即可

ioutil.ReadAll -> io.ReadAll
ioutil.ReadFile -> os.ReadFile
ioutil.ReadDir -> os.ReadDir
// others
ioutil.NopCloser -> io.NopCloser
ioutil.ReadDir -> os.ReadDir
ioutil.TempDir -> os.MkdirTemp
ioutil.TempFile -> os.CreateTemp
ioutil.WriteFile -> os.WriteFile
*/
