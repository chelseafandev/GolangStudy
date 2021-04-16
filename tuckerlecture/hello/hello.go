package main

import (
	"GolangStudy/tuckerlecture/accounts"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	account := accounts.NewAccount("junhaeng")
	fmt.Println(account)
}
