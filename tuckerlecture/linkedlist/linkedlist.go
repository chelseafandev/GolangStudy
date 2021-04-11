package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}

	l.tail.next = &Node{val: val}
	l.tail = l.tail.next
}

func (l *LinkedList) RemoveNode(target *Node) {
	if target == l.root {
		l.root = l.root.next
		return
	}

	prev := l.root
	for prev.next != target {
		prev = prev.next
	}

	if target == l.tail {
		prev.next = nil
		l.tail = prev
	} else {
		prev.next = target.next
	}

	target.next = nil
}

func (l *LinkedList) PrintNodes() {
	node := l.root
	for node.next != nil {
		fmt.Print(node.val, " -> ")
		node = node.next
	}
	fmt.Println(node.val)
}

func main() {
	list := &LinkedList{}
	list.AddNode(0)
	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()
	list.RemoveNode(list.root)
	list.PrintNodes()
	list.RemoveNode(list.root.next)
	list.PrintNodes()
	list.RemoveNode(list.tail)
	list.PrintNodes()
}
