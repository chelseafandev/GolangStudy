package main

import (
	"GolangStudy/tuckerlecture/datastruct"
	"fmt"
)

func main() {
	binaryTree := datastruct.NewBinaryTree(5)
	binaryTree.Root.AddNode(3)
	binaryTree.Root.AddNode(2)
	binaryTree.Root.AddNode(4)
	binaryTree.Root.AddNode(8)
	binaryTree.Root.AddNode(7)
	binaryTree.Root.AddNode(6)
	binaryTree.Root.AddNode(10)
	binaryTree.Root.AddNode(9)

	binaryTree.PrintBinaryTree()
	fmt.Println()

	val := 5
	isFound, count := binaryTree.Search(val)
	if isFound {
		fmt.Println("found", val, ", count:", count)
	} else {
		fmt.Println("not found", val, ", count:", count)
	}
}
