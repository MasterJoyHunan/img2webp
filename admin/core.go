package img2webp

import (
	"archive/zip"
	"bytes"
	"errors"
	"github.com/chai2010/webp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// 将图片转换为 webp 格式  quality 默认 80
func WebpEncoder(f *multipart.FileHeader, quality float32, path string) (err error) {
	var buf bytes.Buffer
	var img image.Image

	file, err := f.Open()
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	contentType := http.DetectContentType(data[:512])
	var fileName string
	if strings.Contains(contentType, "jpeg") {
		fileName = strings.Trim(f.Filename, ".jpeg")
		img, err = jpeg.Decode(bytes.NewReader(data))
	} else if strings.Contains(contentType, "png") {
		fileName = strings.Trim(f.Filename, ".png")
		img, err = png.Decode(bytes.NewReader(data))
	} else if strings.Contains(contentType, "gif") {
		fileName = strings.Trim(f.Filename, ".gif")
		img, err = gif.Decode(bytes.NewReader(data))
	}
	if err != nil {
		return
	}
	if img == nil {
		msg := "image not supported"
		err = errors.New(msg)
		return
	}

	if err = webp.Encode(&buf, img, &webp.Options{Lossless: false, Quality: quality}); err != nil {
		return
	}

	toPath := path + "/" + fileName + ".webp"
	if _, err = os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, os.ModePerm); err != nil {
			return
		}
	}
	if err = ioutil.WriteFile(toPath, buf.Bytes(), 0644); err != nil {
		return
	}
	return nil
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
			return nil
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
