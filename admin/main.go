package main

import (
	"flag"
	"fmt"

	"img2webp/internal"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host", "0.0.0.0", "IP地址")
	flag.IntVar(&port, "port", 8865, "端口号")
}

func main() {
	flag.Parse()
	internal.Setup()
	router := internal.InitRouter()
	panic(router.Run(fmt.Sprintf("%s:%d", host, port)))
}
