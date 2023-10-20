package main

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func main() {

	// 声明 IP 端口和协议
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err:::", err)
	}
	defer listen.Close()

	for true {
		// 操作资源
		fmt.Println("server is waiting...等待客户端连接")
		conn, _ := listen.Accept()
		fmt.Println("conn:", conn)
		// conn.Write  conn.Read()

		for true {
			// 接受客户端消息
			data := make([]byte, 10240)
			fmt.Println("等待客户端输入数据 read等待")
			n, _ := conn.Read(data)
			fmt.Println("data:", data[:n])

			if n == 0 {
				// 客户端意外退出
				conn.Write([]byte(string("0")))
				break

			}

			cmdStr := string(data[:n])
			fmt.Println("cmdStr", string(data))
			//cmd := exec.Command(cmdName, cmdArgs...)

			// 解析用户输入
			parts := strings.Split(cmdStr, " ")
			if len(parts) == 0 {
				fmt.Println("没有输入命令。")
				return
			}

			cmdName := parts[0]
			cmdArgs := parts[1:]

			fmt.Println("解析用户输入:", cmdName, cmdArgs)
			cmd := exec.Command(cmdName, cmdArgs...)
			ret, error := cmd.Output()
			if error != nil {
				fmt.Println(error)
				conn.Write([]byte("输入命令错误"))
			}
			fmt.Println("ret", ret)
			fmt.Println("命令结果长度,", len(ret))
			conn.Write([]byte(ret))

		}
	}

}
