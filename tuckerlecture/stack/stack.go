package main

import (
	"GolangStudy/tuckerlecture/datastruct"
	"fmt"
)

func main() {
	var stack = datastruct.NewStack()
	stack.Push(100)
	stack.Push(200)
	stack.Push(300)
	stack.Push(400)
	stack.Push(500)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
