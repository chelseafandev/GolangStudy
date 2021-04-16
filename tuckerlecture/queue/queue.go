package main

import (
	"GolangStudy/tuckerlecture/datastruct"
	"fmt"
)

func main() {
	queue := datastruct.NewQueue()

	queue.Push(100)
	queue.Push(200)
	queue.Push(300)
	queue.Push(400)
	queue.Push(500)

	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}
