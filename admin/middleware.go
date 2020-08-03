package img2webp

import (
	"github.com/gin-gonic/gin"
	"log"
)

func NoFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Writer.Status() == 404 {
			c.AbortWithStatus(404)
			return
		}
	}
}


func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
		if len(c.Errors) == 0 {
			return
		}
		log.Print(c.Errors.Last().Err)
		// 其他未知错误
		c.AbortWithStatusJSON(500, gin.H{
			"code": 500,
			"data": "",
			"msg":  "系统错误",
		})
		return
	}
}

