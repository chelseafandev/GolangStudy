package main

import "fmt"

type Node struct {
	val  int
	prev *Node
	next *Node
}

type DoubleLinkedList struct {
	root *Node
	tail *Node
}

func (l *DoubleLinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}

	l.tail.next = &Node{val: val}
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
}

func (l *DoubleLinkedList) RemoveNode(target *Node) {
	if target == l.root {
		l.root = l.root.next
		l.root.prev = nil
		target.next = nil
		return
	}

	prev := target.prev

	if target == l.tail {
		prev.next = nil
		l.tail.prev = nil
		l.tail = prev
	} else {
		target.prev = nil
		prev.next = target.next
		(prev.next).prev = prev
	}

	target.next = nil
}

func (l *DoubleLinkedList) PrintNodes() {
	node := l.root
	for node.next != nil {
		fmt.Print(node.val, " <-> ")
		node = node.next
	}
	fmt.Println(node.val)
}

func (l *DoubleLinkedList) PrintNodesReverse() {
	node := l.tail
	for node.prev != nil {
		fmt.Print(node.val, " <-> ")
		node = node.prev
	}
	fmt.Println(node.val)
}

func main() {
	list := &DoubleLinkedList{}
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

	list.PrintNodesReverse()
}
