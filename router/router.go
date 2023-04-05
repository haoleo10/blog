package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	//e.GET("/index", controller.ListUser)
	e.POST("/register", controller.Register)
	e.GET("/register", controller.GoRegister)
	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)
	//跳转到首页，首页就是这个“/”
	e.GET("/", controller.Index)

	//跳转到所有博客的一个列表
	e.GET("/post_index", controller.GetPostIndex)
	//这个post是真正意义上的一个添加
	e.POST("/post", controller.AddPost)
	//跳转到添加的那个页面
	e.GET("/post", controller.GoAddPost)
	e.Run()
}
