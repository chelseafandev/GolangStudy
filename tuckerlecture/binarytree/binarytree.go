package main

import "GolangStudy/tuckerlecture/datastruct"

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
}
