package tree

// https://www.hello-algo.com/chapter_tree/binary_search_tree/

type bsTree struct {
	root *node
}

func NewBsTree(r *node) *bsTree {
	return &bsTree{root: r}
}

func (t *bsTree) Search(v int) *node {
	r := t.root

	for !r.isEmpty() {
		if r.value > v {
			r = r.left
		} else if r.value < v {
			r = r.right
		} else {
			// got it
			break
		}
	}
	return r
}

func (t *bsTree) Insert(v int) {
	r := t.root
	var pre *node

	for !r.isEmpty() {
		if r.value > v {
			pre = r
			r = r.left
		} else if r.value < v {
			pre = r
			r = r.right
		} else {
			// equal, do nothing
			return
		}
	}

	if pre == nil {
		return
	}

	if pre.value > v {
		pre.left = NewNode(v)
	} else {
		pre.right = NewNode(v)
	}
}

func (t *bsTree) Remove(v int) *node {
	r := t.root
	var pre *node

	for !r.isEmpty() {
		if r.value == v {
			break
		}

		pre = r
		if r.value > v {
			r = r.left
		} else {
			r = r.right
		}
	}

	if r == nil {
		return nil
	}

	// 子节点数是 0 或 1
	if r.nonLeft() || r.nonRight() {
		var child *node
		if r.left.isEmpty() {
			child = r.right
		} else {
			child = r.left
		}

		if pre == nil {
			t.root = child
			return r
		}

		if pre.left == r {
			pre.left = child
		} else {
			pre.right = child
		}
		return r
	}

	// 两个子节点，找到中序遍历的下一个节点， 递归删除之后替换 r
	next := t.getInOrderNext(r.right)
	t.Remove(next.value)
	r.value = next.value
	return r
}

// 获取中序遍历的下一个结点，因为中序遍历是顺序的，所以下一个结点一定是在右子树最左侧的位置
func (t *bsTree) getInOrderNext(root *node) *node {
	if root == nil {
		return root
	}
	for root.left != nil {
		root = root.left
	}
	return root
}

func deleteNode(root *node, val int) *node {
	if root == nil {
		return nil
	}
	if root.value == val {
		// case 1, 与 case 2 合并处理
		// if root.left == nil && root.right == nil {
		// 	return nil
		// }

		// case 2
		if root.left == nil {
			return root.right
		}
		if root.right == nil {
			return root.left
		}

		// case 3
		successor := root.right
		for successor != nil && successor.left != nil {
			successor = successor.left
		}

		root.right = deleteNode(root.right, successor.value)
		successor.left = root.left
		successor.right = root.right
		root = successor
	} else if root.value < val {
		root.right = deleteNode(root.right, val)
	} else {
		root.left = deleteNode(root.left, val)
	}

	return root
}
