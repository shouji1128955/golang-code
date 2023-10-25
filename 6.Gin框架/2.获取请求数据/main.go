package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func auth(c *gin.Context) {
	//实现登录判断-此处根据情况修改index中的请求方式，

	//post验证部分
	username := c.PostForm("user")
	password := c.PostForm("pwd")
	if username == "zhang" && password == "123.com" {
		c.String(200, "登录成功!")
	}
	fmt.Println(c.Request.Method)

	//get验证获取参数部分-必须通过url携带参数验证,postman不支持
	if c.Request.Method == "GET" {
		c.JSON(200, gin.H{
			"username": c.Query("user"),
			"pwd":      c.Query("pwd"),
		})
	}
}

func getBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"books": []string{"西游记", "三国演义", "金瓶梅", "国色天香"}, //json一个key,多个值
	})
}

func AddBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "添加书本",
	})
}

func EditBook(c *gin.Context) {
	c.String(200, "EditBook books")

}

func DeleteBook(c *gin.Context) {

	//获取Header信息
	c.String(200, "DeleteBook books")
	c.JSON(200, gin.H{
		"Method":     c.Request.Method,
		"URL":        c.Request.URL,
		"RemoteAddr": c.ClientIP(),
		"Header":     c.Request.Header,
	})

	//获取路劲参数param
	id := c.Param("id")
	fmt.Println("获取的参数id:", id)

	//获取查询参数query

}

func main() {
	//获取路由对象
	router := gin.Default()

	//加载相应得HTML文件
	router.LoadHTMLGlob("./templates/*")
	//LoadHTMLGlob 可以直接加载 目录下面所有得文件
	//返回一个json数据

	//返回一个页面文件
	router.GET("getBook", getBook)
	router.GET("index", index)
	router.POST("auth", auth)
	//获取参数验证
	router.GET("auth", auth)
	//实现路由分组
	bookRoute := router.Group("/books")
	{
		bookRoute.GET("/", getBook)
		bookRoute.GET("/add", AddBook)
		bookRoute.GET("/edit", EditBook)
		bookRoute.GET("/delete", DeleteBook)
		bookRoute.GET("/delete/:id", DeleteBook)
	}
	router.Run(":8080")

	// 书籍相关的路由
}
