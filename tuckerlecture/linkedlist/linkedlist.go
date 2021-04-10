package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func main() {

	var root *Node = &Node{val: 0, next: nil}
	var tail *Node = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}
	PrintNodes(root)
	root, tail = RemoveNode(root, root, tail)
	PrintNodes(root)
	root, tail = RemoveNode(root.next.next, root, tail)
	PrintNodes(root)
	root, tail = RemoveNode(tail, root, tail)
	PrintNodes(root)
}

func AddNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

func RemoveNode(target *Node, root *Node, tail *Node) (*Node, *Node) {
	if target == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	prev := root
	for prev.next != target {
		prev = prev.next
	}

	if target == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = (prev.next).next
	}

	return root, tail
}

func PrintNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Print(node.val, "->")
		node = node.next
	}
	fmt.Println(node.val)
}
