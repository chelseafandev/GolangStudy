package datastruct

import "fmt"

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree(val int) *BinaryTree {
	return &BinaryTree{Root: &BinaryTreeNode{Val: val}}
}

func (b *BinaryTreeNode) AddNode(val int) {
	if b.Val > val {
		if b.Left == nil {
			b.Left = &BinaryTreeNode{Val: val}
			return
		} else {
			b.Left.AddNode(val)
		}
	} else {
		if b.Right == nil {
			b.Right = &BinaryTreeNode{Val: val}
			return
		} else {
			b.Right.AddNode(val)
		}
	}
}

type depthNode struct {
	depth int
	node  *BinaryTreeNode
}

func (b *BinaryTree) PrintBinaryTree() {
	// BFS로 Binary Tree 출력하기
	q := []depthNode{}
	q = append(q, depthNode{depth: 0, node: b.Root})
	currentDepth := 0

	for len(q) > 0 {
		var front depthNode
		front, q = q[0], q[1:]

		if front.depth != currentDepth {
			fmt.Println()
			currentDepth = front.depth
		}
		fmt.Print(front.node.Val, " ")

		if front.node.Left != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: front.node.Left})
		}

		if front.node.Right != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: front.node.Right})
		}
	}
}
