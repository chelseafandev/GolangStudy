package main

import (
	"GolangStudy/accounts"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	account := accounts.NewAccount("junhaeng")
	fmt.Println(account)
}
