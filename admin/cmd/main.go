package main

import (
	"flag"
	"fmt"
	"img2webp"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8865, "设置端口")
}

func main() {
	flag.Parse()
	img2webp.Setup()
	router := img2webp.InitRouter()
	panic(router.Run(fmt.Sprintf(":%d", port)))
}
