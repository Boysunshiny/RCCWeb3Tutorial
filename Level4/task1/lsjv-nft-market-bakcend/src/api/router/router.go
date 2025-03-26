package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// r.Use(middleware.RecoverMiddleware()) // 使用恢复中间件
	// r.Use(middleware.RLog())              // 使用日志中间件

	r.Use(cors.New(cors.Config{ // 使用cors中间件
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-CSRF-Token", "Authorization", "AccessToken", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "X-GW-Error-Code", "X-GW-Error-Message"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}))
	loadV1(r)
	return r
}
