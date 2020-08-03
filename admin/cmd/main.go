package main

import "img2webp"

func main() {
	router := img2webp.InitRouter()
	panic(router.Run(":8888"))
}
