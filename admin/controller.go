package img2webp

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"time"
)

// 历史记录
func History(c *gin.Context) {
	res := RedisClient.ZRevRange(context.Background(), Historys, 0, 99).Val()
	c.JSON(200, gin.H{
		"code": 1,
		"data": res,
		"msg":  "",
	})
}

// 上传
func Upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.Error(err)
		return
	}

	files := form.File["upload"]
	if files == nil {
		c.Error(errors.New("文件不能为空"))
		return
	}

	// 图片转换
	dir := fmt.Sprintf("./%d", time.Now().Unix())
	for _, file := range files {
		if err := WebpEncoder(file, QUALITY, dir); err != nil {
			c.Error(err)
			return
		}
	}

	// 压缩打包
	err = Zip(dir)
	if err != nil {
		c.Error(err)
		return
	}

	//

	zip, err := ioutil.ReadFile(dir + ".zip")
	c.Header("content-type", "application/zip")
	c.Header("content-disposition", fmt.Sprintf("attachment; filename=\"%s_%d.zip\"", time.Now().Format("2006-01-02 15:04:05"), rand.Int()))
	c.Writer.Write(zip)
}
