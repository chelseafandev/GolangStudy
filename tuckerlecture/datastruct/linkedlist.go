package datastruct

import "fmt"

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type LinkedList struct {
	Root *Node
	Tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.Root == nil {
		l.Root = &Node{Val: val}
		l.Tail = l.Root
		return
	}

	l.Tail.Next = &Node{Val: val}
	prev := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.Prev = prev
}

func (l *LinkedList) RemoveNode(target *Node) {
	if target == l.Root {
		l.Root = l.Root.Next
		l.Root.Prev = nil
		target.Next = nil
		return
	}

	prev := target.Prev

	if target == l.Tail {
		prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = prev
	} else {
		target.Prev = nil
		prev.Next = target.Next
		(prev.Next).Prev = prev
	}

	target.Next = nil
}

func (l *LinkedList) PrintNodes() {
	node := l.Root
	for node.Next != nil {
		fmt.Print(node.Val, " <-> ")
		node = node.Next
	}
	fmt.Println(node.Val)
}

func (l *LinkedList) PrintNodesReverse() {
	node := l.Tail
	for node.Prev != nil {
		fmt.Print(node.Val, " <-> ")
		node = node.Prev
	}
	fmt.Println(node.Val)
}
