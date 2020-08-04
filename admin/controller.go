package img2webp

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// 历史记录
func History(c *gin.Context) {
	res := RedisClient.ZRevRange(context.Background(), HISTORY, 0, 99).Val()
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

	files := form.File["file"]
	if files == nil {
		c.Error(errors.New("文件不能为空"))
		return
	}

	// 图片转换
	bashDir := fmt.Sprintf("%s_%d", time.Now().Format("20060102150405"), rand.Int())
	for _, file := range files {
		if err := WebpEncoder(file, QUALITY, bashDir); err != nil {
			c.Error(err)
			return
		}
	}

	// 压缩打包
	if err = Zip(bashDir); err != nil {
		c.Error(err)
		return
	}

	// 删除文件夹及内的文件
	if err = os.RemoveAll(bashDir); err != nil {
		c.Error(err)
		return
	}

	// 加入 redis
	pip := RedisClient.Pipeline()
	ctx := context.Background()
	pip.ZAdd(ctx, HISTORY, &redis.Z{
		Score:  float64(time.Now().UnixNano()),
		Member: bashDir + ".zip",
	})
	pip.ZRemRangeByRank(ctx, HISTORY, 0, -101)
	if _, err := pip.Exec(ctx); err != nil {
		c.Error(err)
		return
	}

	// 返回地址
	c.JSON(200, gin.H{
		"code": 1,
		"data": bashDir + ".zip",
		"msg":  "success",
	})
}

// 下载
func Download(c *gin.Context) {
	filename := c.Query("filename")
	zip, err := ioutil.ReadFile("webp_zip/" + filename)
	if err != nil {
		c.Error(err)
		return
	}
	c.Header("content-type", "application/zip")
	c.Header("content-disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", filename))
	c.Writer.Write(zip)
}
