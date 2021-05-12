package internal

import (
	"archive/zip"
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
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
		return errors.New("image not supported")
	}

	if err = webp.Encode(&buf, img, &webp.Options{Lossless: false, Quality: quality}); err != nil {
		return
	}

	toPath := path + "/" + fileName + ".webp"
	if _, err = os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			return
		}
	}
	if err = ioutil.WriteFile(toPath, buf.Bytes(), 0644); err != nil {
		return
	}
	return nil
}

// 压缩打包
func Zip(dir string) (err error) {
	if _, err = os.Stat("webp_zip"); os.IsNotExist(err) {
		if err = os.MkdirAll("webp_zip", os.ModePerm); err != nil {
			return
		}
	}
	fz, err := os.Create("webp_zip/" + dir + ".zip")
	if err != nil {
		return
	}
	defer fz.Close()

	w := zip.NewWriter(fz)
	defer w.Close()

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fDest, err := w.Create(path[len(dir)+1:])
		if err != nil {
			return err
		}
		fSrc, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fSrc.Close()

		_, err = io.Copy(fDest, fSrc)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}

	return nil
}
