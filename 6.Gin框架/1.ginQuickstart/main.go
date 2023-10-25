package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func auth(c *gin.Context) {
	//实现登录判断
	username := c.PostForm("user")
	password := c.PostForm("pwd")
	if username == "zhang" && password == "123.com" {
		c.String(200, "登录成功!")
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

	fmt.Println(c.Request.Method)
	c.String(200, "DeleteBook books"+"\n"+"你的请方法是:")
	c.JSON(200, gin.H{
		"method":   c.Request.Method,
		"URL":      c.Request.URL,
		"ADDRESS":  c.Request.RemoteAddr,
		"ClientIP": c.ClientIP(),
		"AGENT":    c.Request.Header["User-Agent"],
		"ALL":      c.Request,
	})

}

func postcomit(ctx *gin.Context) {
	//当前端请求的数据通过JSON提交时，例如向/xxx发送一个POST请求，则获取请求参数的方式如下
	// 获取request.Body() 中的数据(这里没有进行错误处理)
	// 返回的是字节数组
	dataBytes, _ := ctx.GetRawData()
	fmt.Println("dataBytes", dataBytes)
	// 定义一个map
	var m map[string]interface{}

	// 反序列化 别忘了&
	_ = json.Unmarshal(dataBytes, &m)

	fmt.Println(m)
	// 数据返回
	ctx.JSON(http.StatusOK, m)

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
	router.POST("xxx", postcomit)
	//实现路由分组
	bookRoute := router.Group("/books")
	{
		bookRoute.GET("/", getBook)
		bookRoute.GET("/add", AddBook)
		bookRoute.GET("/edit", EditBook)
		bookRoute.GET("/delete", DeleteBook)
	}
	// GET

	router.Run(":8080")

	// 书籍相关的路由

}
