package main

import (
	db "4.包管理/1.包的使用/db"
	"fmt"
)
import (
	"4.包管理/1.包的使用/api"
)

func main() {
	api.RestAPI01()
	api.RestAPI02()
	api.RpcAPI01()
	api.RpcAPI02()
	db.MySQLHandler()
	fmt.Println("hello")
}
