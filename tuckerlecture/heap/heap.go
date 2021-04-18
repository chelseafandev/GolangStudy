package main

import (
	"GolangStudy/tuckerlecture/datastruct"
	"fmt"
)

func main() {
	heap := datastruct.NewHeap()
	heap.Push(9)
	heap.Push(4)
	heap.Push(6)
	heap.Push(3)
	heap.Push(2)
	heap.PrintHeap()
	fmt.Println()

	fmt.Printf("Popped: %d\n", heap.Pop())
	heap.PrintHeap()
}
