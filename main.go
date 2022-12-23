package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 拦截器
func myhandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		//这里设置了别的地方都能拿到
		context.Set("usersession", "user-1")
		context.Next()  //放行
		context.Abort() //阻止
	}
}

func main() {
	//create a server
	ginServer := gin.Default()

	//request 使用拦截器
	ginServer.GET("/hello", myhandler(), func(context *gin.Context) {
		log.Println("---------logggggg")
		context.JSON(200, gin.H{"msg": "hello from go"})
	})

	//load the html
	ginServer.LoadHTMLGlob("templates/*")

	//load the static
	ginServer.Static("/static", "./static")

	//return a html
	ginServer.GET("/test", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"name": "BUBU",
		})
	})

	//url?name=xxx&age=xxx
	ginServer.GET("/test1", func(context *gin.Context) {
		name := context.Query("name")
		age := context.Query("age")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	//url/yhr/33
	ginServer.GET("/test2/:username/:age", func(context *gin.Context) {
		name := context.Param("username")
		age := context.Param("age")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	//json
	ginServer.POST("/json", func(context *gin.Context) {
		b, _ := context.GetRawData()
		var m map[string]interface{}
		json.Unmarshal(b, &m)
		context.JSON(http.StatusOK, m)
	})

	//form
	ginServer.POST("/form", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	//404
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(404, "404.html", gin.H{})
	})

	//route group /user/add
	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("/add")
	}

	//run
	ginServer.Run(":8080")
}
