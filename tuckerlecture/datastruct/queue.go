package datastruct

import "fmt"

type Queue struct {
	ll *LinkedList
}

func NewQueue() *Queue {
	return &Queue{ll: NewLinkedList()}
}

func (q *Queue) empty() bool {
	return q.ll.Empty()
}

func (q *Queue) Push(val int) {
	q.ll.AddNode(val)
}

func (q *Queue) Pop() int {
	var popped int
	if !q.empty() {
		popped = q.ll.Front()
		if q.ll.Root != nil {
			q.ll.RemoveNode(q.ll.Root)
		}
	} else {
		fmt.Println("Queue is empty!")
	}
	return popped
}
