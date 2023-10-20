package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//网络主要涉及三个方面, 协议,IP,端口
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("err:::", err)
	}
	fmt.Println("正在等待连接....")
	conn, _ := listen.Accept()
	fmt.Println("conn:", conn)
	defer listen.Close()
	for {

		fmt.Println(conn)
		//接受客户端的请求
		data := make([]byte, 1024)
		fmt.Println("等待客户端输入数据 read等待")
		//读取数据时,如果客户端只发送了一条数据，conn.Read读取完之后，不能重复读，消息第一次取走后，需要重新发
		n, _ := conn.Read(data)
		if n == 0 {
			//判断如果返回的数据为空 这里的n表示索引的位置
			break
		}
		fmt.Println(string(data))
		//小写转大写服务
		uppperStr := strings.ToUpper(string(data[:n]))
		conn.Write([]byte(uppperStr))

	}

}
