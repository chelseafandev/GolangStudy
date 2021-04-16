package main

import (
	"GolangStudy/tuckerlecture/datastruct"
	"fmt"
)

func main() {
	val := 1
	tree := datastruct.NewTree(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(tree.Root.Children); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Children[i].AddNode(val)
			val++
		}
	}

	tree.DFS_recursive()
	fmt.Println()

	tree.DFS_stack()
	fmt.Println()

	tree.BFS()
}
