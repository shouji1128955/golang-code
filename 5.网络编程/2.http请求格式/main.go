package main

import (
	"fmt"
	"net"
	"os"
)

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
		fmt.Println(string(data[:n]))
		//conn.Write([]byte("HTTP/1.1 200 ok\r\n\r\n<div>hello, zhanglaiqiang</div> <img src='https://img2.baidu.com/it/u=2595535765,1140102480&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1697907600&t=ded41d14bf216fe04bc26194cf97e097'>"))

		//将内容写入到一个文件
		data, _ = os.ReadFile("index.html")

		fmt.Println("##########")
		fmt.Println(string(data))
		fmt.Println("*************")
		conn.Write([]byte("HTTP/1.1 200 ok\r\n\r\n" + string(data)))
		conn.Close()

	}

}
