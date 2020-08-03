package img2webp

import (
	"github.com/gin-gonic/gin"
	//"tc/controller"
	//"tc/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// 加入通用中间件
	r.Use(
		gin.Recovery(),             // recovery 防止程序奔溃
		NoFound(),					// 404
		ErrorHandle(),              // 错误处理
	)

	r.GET("/history", History)
	r.POST("/upload", Upload)
	return r
}
