package tree

import "testing"

func newBsTree() *bsTree {
	tree := NewBsTree(NewNode(5))
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(11)

	/*
				5
			2		9
		1	   3		11
	*/
	return tree
}

func TestBsTree_Search(t *testing.T) {
	tree := newBsTree()

	println(tree.Search(3))
	println(tree.Search(4))
}

func TestBsTree_Insert(t *testing.T) {
	tree := newBsTree()
	tree.Insert(4)
	tree.Insert(10)

	dfs(tree.root, 1)
}

func TestBsTree_Remove(t *testing.T) {
	t.Run("", func(t *testing.T) {
		tree := newBsTree()
		tree.Remove(1)
		dfs(tree.root, 1)
	})

	t.Run("", func(t *testing.T) {
		tree := newBsTree()
		tree.Remove(9)
		dfs(tree.root, 1)
	})

	t.Run("", func(t *testing.T) {
		tree := newBsTree()
		tree.Remove(2)
		dfs(tree.root, 0)
	})

	t.Run("", func(t *testing.T) {
		tree := newBsTree()
		tree.Remove(5)
		dfs(tree.root, 0)
	})

	t.Run("", func(t *testing.T) {
		tree := newBsTree()
		tree.Remove(9)
		tree.Remove(11)
		tree.Remove(5)
		dfs(tree.root, 0)
	})
}
