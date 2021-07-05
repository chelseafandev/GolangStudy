package main

import (
	https_server "GolangStudy/SideProject/https"
	"fmt"
)

func main() {
	fmt.Println("start https server!")
	https_server.StartServer("127.0.0.1", 18443, "/test")

	fmt.Println("hello world")
}
