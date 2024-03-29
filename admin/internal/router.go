package internal

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// 加入通用中间件
	r.Use(
		gin.Recovery(), // recovery 防止程序奔溃
		NoFound(),      // 404
		ErrorHandle(),  // 错误处理
	)

	r.Static("/download", "./webp_zip")
	r.GET("/history", History)
	r.POST("/upload", Upload)
	return r
}
