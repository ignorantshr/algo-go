package tree

import "testing"

func newAvlTree() *avlTree {
	tree := NewAvlTree(&hnode{
		height: 0,
		value:  6,
		left:   nil,
		right:  nil,
	})
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(9)

	/*
				6
			2		7
		1	   3 		9
	*/

	return tree
}

func TestAvlTree_Insert(t *testing.T) {
	tree := newAvlTree()
	dfs(tree.root, 1)
}

func TestAvlTree_Search(t *testing.T) {
	tree := newAvlTree()
	bfs(tree.root)
}

func TestAvlTree_Remove(t *testing.T) {
	t.Run("", func(t *testing.T) {
		tree := newAvlTree()
		tree.Remove(1)
		dfs(tree.root, 1)
	})

	t.Run("", func(t *testing.T) {
		tree := newAvlTree()
		tree.Remove(7)
		dfs(tree.root, 1)
	})

	t.Run("", func(t *testing.T) {
		tree := newAvlTree()
		tree.Remove(3)
		dfs(tree.root, 1)
	})

	t.Run("", func(t *testing.T) {
		tree := newAvlTree()
		tree.Remove(6)
		dfs(tree.root, 1)
	})

}
