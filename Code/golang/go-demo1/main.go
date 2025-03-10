package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {

	})

	r.LoadHTMLFiles("templates/index.html")
	r.GET("/", func(c *gin.Context) {
		// 渲染 HTML 模板，并传递数据
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home Page",
		})
	})
	r.Run()
}
