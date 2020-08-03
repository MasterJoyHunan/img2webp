package img2webp

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path/filepath"
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
	dir := fmt.Sprintf("./%d",time.Now().Unix())
	for _, file := range files {
		if err := WebpEncoder(file, QUALITY, dir); err != nil {
			c.Error(err)
			return
		}
	}

	// 压缩打包
	Zip(dir)

	c.JSON(200, gin.H{
		"code": 1,
		"data": "",
		"msg":  "",
	})
}


// 压缩打包
func Zip(dir string) {
	fz, err := os.Create(dir + ".zip")
	if err != nil {
		log.Fatalf("Create zip file failed: %s\n", err.Error())
	}
	defer fz.Close()

	w := zip.NewWriter(fz)
	defer w.Close()

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return  nil
		}
		fDest, err := w.Create(path[len(dir)+1:])
		if err != nil {
			log.Printf("Create failed: %s\n", err.Error())
			return nil
		}
		fSrc, err := os.Open(path)
		if err != nil {
			log.Printf("Open failed: %s\n", err.Error())
			return nil
		}
		defer fSrc.Close()
		_, err = io.Copy(fDest, fSrc)
		if err != nil {
			log.Printf("Copy failed: %s\n", err.Error())
			return nil
		}
		return nil
	})
}

