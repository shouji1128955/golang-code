package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Println("当前的连接通道:", conn)
	fmt.Println(conn)
	defer conn.Close()
	for {

		//开始输入数据
		var word string
		fmt.Println("请输入一个英文单词")
		fmt.Scan(&word)
		conn.Write([]byte(word))

		//接受响应的数据
		data := make([]byte, 1024)
		n, _ := conn.Read(data)
		fmt.Println("服务器返回的数据:", string(data[:n]))
	}
}
