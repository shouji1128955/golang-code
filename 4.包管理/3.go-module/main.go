package main

import (
	db "api222/db"
	"fmt"
)
import (
	"api222/api"
)

func main() {
	api.RestAPI01()
	api.RestAPI02()
	api.RpcAPI01()
	api.RpcAPI02()
	db.MySQLHandler()
	fmt.Println("hello")
}
