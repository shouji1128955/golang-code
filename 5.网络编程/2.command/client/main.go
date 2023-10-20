package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Println("conn:::", conn)

	defer conn.Close()

	for true {
		// conn.Write  conn.Read()
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("请输入命令和参数: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // 去掉前后空格

		conn.Write([]byte(input))

		// 接受服务端响应数据
		res := make([]byte, 10240)
		n, _ := conn.Read(res)
		fmt.Println("命令结果长度,", n)
		fmt.Println("服务端响应数据：", string(res[:n]))
	}

}
