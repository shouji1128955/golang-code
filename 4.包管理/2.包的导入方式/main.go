package main

import (
	api2 "4.包管理/1.包的使用/api" //如果包名冲突,可以通过这种方式重命名
	. "4.包管理/2.包的导入方式/db"   //直接导入所有的变量，不需要在通过db.这种方式来调用
	"fmt"
)

func main() {
	api2.RestAPI01()
	api2.RestAPI02()
	api2.RpcAPI01()
	api2.RpcAPI02()
	MySQLHandler()
	fmt.Println("hello")
}
