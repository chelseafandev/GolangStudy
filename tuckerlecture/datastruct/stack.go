package datastruct

import "fmt"

type Stack struct {
	ll *LinkedList
}

func NewStack() *Stack {
	return &Stack{ll: NewLinkedList()}
}

func (s *Stack) empty() bool {
	return s.ll.Empty()
}

func (s *Stack) Push(val int) {
	s.ll.AddNode(val)
}

func (s *Stack) Pop() int {
	var popped int
	if !s.empty() {
		popped = s.ll.Back()
		if s.ll.Tail != nil {
			s.ll.RemoveNode(s.ll.Tail)
		}
	} else {
		fmt.Println("Stack is empty!")
	}
	return popped
}
