package datastruct

import "fmt"

type TreeNode struct {
	Val      int
	Children []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func NewTree(val int) *Tree {
	return &Tree{Root: &TreeNode{Val: val}}
}

func (t *TreeNode) AddNode(val int) {
	t.Children = append(t.Children, &TreeNode{Val: val})
}

func dFS_recursive(node *TreeNode) {
	fmt.Printf("%d->", node.Val)

	for i := 0; i < len(node.Children); i++ {
		dFS_recursive(node.Children[i])
	}
}

func (t *Tree) DFS_recursive() {
	dFS_recursive(t.Root)
}

func (t *Tree) DFS_stack() {
	// 우리 이전 강좌에서 구현해둔 stack은 int형 타입만을 val값으로 받을 수 있기때문에
	// DFS_stack 함수에서 사용하려는 TreeNode의 포인터 타입은 구현해둔 stack으로 사용할 수 없다.
	// interface에 대해서는 추후 강좌에서 설명할 예정이다.
	// 그래서 일단 여기에서는 stack을 slice로 구현해서 사용하겠다
	s := []*TreeNode{}
	s = append(s, t.Root)

	for len(s) > 0 {
		var top *TreeNode
		top, s = s[len(s)-1], s[:len(s)-1]
		fmt.Printf("%d->", top.Val)
		for i := len(top.Children) - 1; i >= 0; i-- {
			s = append(s, top.Children[i])
		}
	}
}
