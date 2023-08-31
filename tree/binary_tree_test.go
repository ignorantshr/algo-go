package tree

import (
	"fmt"
	"testing"
)

func newBinaryTree(t *testing.T) *node {
	t.Helper()

	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	n4 := NewNode(4)
	n5 := NewNode(5)
	n6 := NewNode(6)
	n7 := NewNode(7)

	n1.left = n2
	n1.right = n3
	n2.left = n4
	n2.right = n5
	n3.left = n6
	n3.right = n7

	/*
			   1
		     2   3
			4 5 6 7
	*/
	return n1
}

func TestBFS(t *testing.T) {
	bfs(newBinaryTree(t))
}

func TestDFS(t *testing.T) {
	dfs(newBinaryTree(t), 0)
	println()

	dfs(newBinaryTree(t), 1)
	println()
	res := mirrorsDfsInOrder(newBinaryTree(t))
	fmt.Println(res)

	dfs(newBinaryTree(t), 2)
	println()
	postres := postorderTraversal1(newBinaryTree(t))
	fmt.Println(postres)
}

func Test_preorderIterate(t *testing.T) {
	fmt.Println("-------- v1 ---------")
	fmt.Println(preorderTraversal1(newBinaryTree(t)))
	fmt.Println(preorderTraversal2(newBinaryTree(t)))
	fmt.Println(preorderTraversal3(newBinaryTree(t)))

	fmt.Println("-------- v2 ---------")
	fmt.Println(inorderTraversal1(newBinaryTree(t)))
	fmt.Println(mirrorsDfsInOrder(newBinaryTree(t)))
	fmt.Println(inorderTraversal3(newBinaryTree(t)))

	fmt.Println("-------- v3 ---------")
	fmt.Println(postorderTraversal1(newBinaryTree(t)))
	fmt.Println(postorderTraversal2(newBinaryTree(t)))
	fmt.Println(postorderTraversal3(newBinaryTree(t)))
}
