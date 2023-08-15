package tree

import "fmt"

type Comparable interface {
	// Compare compares self with another
	// return -1|0|1: smaller than another, equal, larger
	Compare(another Comparable) int
}

type TreeNode interface {
	isEmpty() bool
	Left() TreeNode
	Right() TreeNode

	String() string
}

// -------------------
// hnode 带高度属性的节点
type hnode struct {
	height      int
	value       int
	left, right *hnode
}

func (n *hnode) Left() TreeNode {
	return n.left
}

func (n *hnode) Right() TreeNode {
	return n.right
}

func (n *hnode) isEmpty() bool {
	return n == nil
}

func (n *hnode) String() string {
	return fmt.Sprintf("[value:%v, height:%v]\n", n.value, n.height)
}

// -------------------
type node struct {
	value       int
	left, right *node
}

func NewNode(value int) *node {
	return &node{
		value: value,
		left:  nil,
		right: nil,
	}
}

func (n *node) Left() TreeNode {
	return n.left
}

func (n *node) Right() TreeNode {
	return n.right
}

func (n *node) String() string {
	return fmt.Sprintf("%v", n.value)
}

func (n *node) isEmpty() bool {
	return n == nil
}

func (n *node) nonLeft() bool {
	return n.left == nil
}

func (n *node) nonRight() bool {
	return n.right == nil
}

func println(a ...interface{}) {
	fmt.Println(a...)
}

func print(a ...interface{}) {
	fmt.Print(a...)
}
