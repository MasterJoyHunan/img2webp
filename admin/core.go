package img2webp

import (
	"bytes"
	"errors"
	"github.com/chai2010/webp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
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


