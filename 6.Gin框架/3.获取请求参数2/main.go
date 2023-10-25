package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Name  string `json:"name" form:"name"`
	Age   int    `json:"age" form:"age"`
	Email string `json:"email" form:"email"`
}

// 获取PostFrom的数据
func postform(c *gin.Context) {
	username := c.PostForm("name")
	age := c.PostForm("age")
	fmt.Println(username)
	fmt.Println(age)
}

func ShouldBind(c *gin.Context) {

	var users = User{} //结构体实例化,相当于 user := User{}
	if c.ShouldBind(&users) == nil {
		//绑定成功，打印请求参数
		log.Println("user", users)
	}
}
func main() {

	//此处主要验证获取PostForm和shouldBind函数
	router := gin.Default()
	router.POST("/postform", postform)

	router.POST("/ShouldBind", ShouldBind)

	router.Run(":8080")

}
