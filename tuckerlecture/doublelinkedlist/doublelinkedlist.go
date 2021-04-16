package main

import "GolangStudy/tuckerlecture/datastruct"

func main() {
	list := datastruct.NewLinkedList()
	list.AddNode(0)
	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()
	list.RemoveNode(list.Root)
	list.PrintNodes()
	list.RemoveNode(list.Root.Next)
	list.PrintNodes()
	list.RemoveNode(list.Tail)
	list.PrintNodes()

	list.PrintNodesReverse()
}
