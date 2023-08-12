package main

import (
	"fmt"

	"sample/go-basic-example/assets"
	"sample/go-basic-example/service"
)

func main() {
	fmt.Println(service.Hello("World"))

	fmt.Println(string(assets.HelloTextBytes))

	data, err := assets.EmbedFile.ReadFile("hello.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
