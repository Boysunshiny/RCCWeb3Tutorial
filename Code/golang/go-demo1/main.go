package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	Name string `form:"name" binding:"required"`
	Age  int    `form:"age" binding:"required"`
}

func main() {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {

	})

	r.LoadHTMLFiles("templates/index.html")

	v1 := r.Group("/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			// 渲染 HTML 模板，并传递数据
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Home Page",
			})
		})
		v1.POST("/test", func(c *gin.Context) {
			//localhost:8080/v1/test?ids=["111111","2222"]&ids2=1&ids2=2&ids2=3
			ids := c.QueryArray("ids")
			ids2 := c.QueryArray("ids2")
			c.JSON(http.StatusOK, gin.H{
				"message": "post request",
				"ids":     ids,
				"ids2":    ids2,
			})
		})
		v1.DELETE("/test", func(ctx *gin.Context) {
			//http://localhost:8080/v1/test?name=John&age=30
			var queryParams QueryParams
			if err := ctx.BindQuery(&queryParams); err != nil {
				// 在绑定失败时自动返回 HTTP 400 状态码和错误信息
				return
			}
			ctx.JSON(http.StatusOK, queryParams)
		})
		v1.PUT("/test", func(ctx *gin.Context) {

		})
	}

	r.Run()
}
