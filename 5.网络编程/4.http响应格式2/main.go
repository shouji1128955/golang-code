package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func getJson(conn net.Conn) {
	//net.Conn是数据类型
	conn.Write([]byte("HTTP/1.1 200 ok\r\ncontent-Type:application/json\r\n\r\n{\"name\":\"zhang\",\"age\":\"28\"}"))
}

func getMeinv(conn net.Conn) {
	data, _ := os.ReadFile("index.html")
	conn.Write([]byte("HTTP/1.1 200 ok\r\ncontent-Type:text/html\r\n\r\n" + string(data)))
}

func other(conn net.Conn) {
	conn.Write([]byte("HTTP/1.1 200 ok\r\ncontent-Type:text/html\r\n\r\n not Found"))
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err::::", err)
	}
	defer listen.Close()
	//基于web得开发

	for {
		conn, _ := listen.Accept()
		data := make([]byte, 1024)
		n, _ := conn.Read(data)

		requestStr := string(data[:n])

		ret := strings.Split(requestStr, "\r\n")
		path := strings.Split(ret[0], " ")[1]

		//打印访问路劲
		fmt.Println("访问路劲:", path)
		//conn.Write([]byte("HTTP/1.1 200 ok\r\n\r\n<div>hello, zhanglaiqiang</div> <img src='https://img2.baidu.com/it/u=2595535765,1140102480&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1697907600&t=ded41d14bf216fe04bc26194cf97e097'>"))

		if path == "/json" {
			//返回json数据
			getJson(conn)
		} else if path == "/meinv" {
			getMeinv(conn)
		} else {
			other(conn)
		}
		//将内容写入到一个文件

		fmt.Println("##########")
		fmt.Println(string(data))
		fmt.Println("*************")

		//在响应得时候，需要指明我们返回的数据是什么类型
		//conn.Write([]byte("HTTP/1.1 200 ok\r\ncontent-Type:text/plain\r\n\r\n" + string(data)))
		//说明: application/json 配置此项,使用postman的时候自动按照json格式能够自动展开，如果配置content-Type:text/plain项，怎以文本形式进行展示

		conn.Close()

	}

}
