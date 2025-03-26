package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loadV1(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")

	user := apiV1.Group("/user")
	{
		user.GET("/:address/login-message", func(ctx *gin.Context) {

			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"address": ctx.Param("address"),
				"message": "生成login签名信息",
			})

		}) // 生成login签名信息
		user.POST("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"message": "登陆",
			})
		}) // 登陆
		user.GET("/:address/sig-status", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"address": ctx.Param("address"),
				"message": "获取用户签名状态",
			})
		}) // 获取用户签名状态
	}

	collections := apiV1.Group("/collections")
	{
		// 接口定义： 路由 + 中间件 + 处理函数
		collections.GET("/:address", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"address": ctx.Param("address"),
				"message": "指定Collection详情",
			})
		}) // 指定Collection详情
		collections.GET("/:address/bids", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"address": ctx.Param("address"),
				"message": "指定Collection的bids信息",
			})
		}) // 指定Collection的bids信息
		collections.GET("/:address/:token_id/bids", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":      ctx.Request.URL,
				"address":  ctx.Param("address"),
				"token_id": ctx.Param("token_id"),
				"message":  "指定Item的bid信息",
			})
		}) // 指定Item的bid信息
		collections.GET("/:address/items", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"address": ctx.Param("address"),
				"message": "指定Collection的items信息",
			})
		}) // 指定Collection的items信息

		collections.GET("/:address/:token_id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":      ctx.Request.URL,
				"address":  ctx.Param("address"),
				"token_id": ctx.Param("token_id"),
				"message":  "获取NFT Item的详细信息",
			})
		}) // 获取NFT Item的详细信息
		collections.GET("/:address/:token_id/traits", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":      ctx.Request.URL,
				"address":  ctx.Param("address"),
				"token_id": ctx.Param("token_id"),
				"message":  "获取NFT Item的Attribute信息",
			})
		}) //获取NFT Item的Attribute信息
		collections.GET("/:address/top-trait", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"address": ctx.Param("address"),

				"message": "获取NFT Item的Trait的最高价格信息",
			})
		}) //获取NFT Item的Trait的最高价格信息
		collections.GET("/:address/:token_id/image", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":      ctx.Request.URL,
				"address":  ctx.Param("address"),
				"token_id": ctx.Param("token_id"),
				"message":  "获取NFT Item的图片信息",
			})
		}) // 获取NFT Item的图片信息
		collections.GET("/:address/history-sales", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"address": ctx.Param("address"),
				"message": "NFT销售历史价格信息",
			})
		}) // NFT销售历史价格信息
		collections.GET("/:address/:token_id/owner", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":      ctx.Request.URL,
				"address":  ctx.Param("address"),
				"token_id": ctx.Param("token_id"),
				"message":  "获取NFT Item的owner信息",
			})
		}) // 获取NFT Item的owner信息
		collections.POST("/:address/:token_id/metadata", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":      ctx.Request.URL,
				"address":  ctx.Param("address"),
				"token_id": ctx.Param("token_id"),
				"message":  "刷新NFT Item的metadata",
			})
		}) // 刷新NFT Item的metadata

		collections.GET("/ranking", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"message": "获取NFT集合排名信息",
			})
		}) // 获取NFT集合排名信息
	}

	activities := apiV1.Group("/activities")
	{
		activities.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"message": "批量获取activity信息",
			})
		}) // 批量获取activity信息
	}

	portfolio := apiV1.Group("/portfolio")
	{
		portfolio.GET("/collections", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"message": "获取用户拥有Collection信息",
			})
		}) // 获取用户拥有Collection信息
		portfolio.GET("/items", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"message": "查询用户拥有nft的Item基本信息",
			})
		}) // 查询用户拥有nft的Item基本信息
		portfolio.GET("/listings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"message": "查询用户挂单的Listing信息",
			})
		}) // 查询用户挂单的Listing信息
		portfolio.GET("/bids", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"message": "查询用户挂单的Bids信息",
			})
		}) // 查询用户挂单的Bids信息
	}

	orders := apiV1.Group("/bid-orders")
	{
		orders.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"url":     ctx.Request.URL,
				"message": "批量查询出价信息",
			})
		}) // 批量查询出价信息
	}
}
