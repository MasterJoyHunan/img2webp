package main

import "img2webp"

func main() {
	img2webp.Setup()
	router := img2webp.InitRouter()
	panic(router.Run(":8888"))
}
