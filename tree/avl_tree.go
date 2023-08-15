package tree

// 平衡二叉树 https://www.hello-algo.com/chapter_tree/avl_tree/

type avlTree struct {
	root *hnode
}

func NewAvlTree(r *hnode) *avlTree {
	return &avlTree{root: r}
}

// 树高就是该从节点到最远节点 边 的数量
func (t *avlTree) treeHeight(node *hnode) int {
	// 叶结点的高度为 0 ，空结点的高度为 -1
	if node == nil {
		return -1
	}
	return node.height
}

func (t *avlTree) updateHeight(node *hnode) {
	lh := t.treeHeight(node.left)
	rh := t.treeHeight(node.right)
	if lh > rh {
		node.height = lh + 1
	} else {
		node.height = rh + 1
	}
}

// 结点的「平衡因子 Balance Factor」是 结点的左子树高度减去右子树高度，并定义空结点的平衡因子为 0 。
// 通过平衡因子可以确定树的平衡状态
func (t *avlTree) balanceFactor(node *hnode) int {
	if node == nil {
		return 0
	}

	lh := t.treeHeight(node.left)
	rh := t.treeHeight(node.right)
	return lh - rh
}

func (t *avlTree) rightRotate(node *hnode) *hnode {
	child := node.left
	grandChild := child.right

	node.left = grandChild
	child.right = node
	// 旋转完成之后需要更新高度
	t.updateHeight(node)
	t.updateHeight(child)
	// 返回旋转后子树的根结点
	return child
}

func (t *avlTree) leftRotate(node *hnode) *hnode {
	child := node.right
	grandChild := child.left

	node.right = grandChild
	child.left = node
	// 旋转完成之后需要更新高度
	t.updateHeight(node)
	t.updateHeight(child)
	// 返回旋转后子树的根结点
	return child
}

// 对于不平衡树的四种情况进行旋转
func (t *avlTree) rotate(node *hnode) *hnode {
	balance := t.balanceFactor(node)

	if balance < -1 {
		// 右偏树
		if t.balanceFactor(node.right) > 0 {
			// 无右孙子节点
			child := node.right
			node.right = t.rightRotate(child)
		}
		// < 0, 无左孙子节点，== 0 有双孙子节点
		return t.leftRotate(node)
	}

	if balance > 1 {
		// 左偏树
		if t.balanceFactor(node.left) < 0 {
			// 无左孙子节点
			child := node.left
			node.left = t.leftRotate(child)
		}
		// > 0, 无右孙子节点，== 0 有孙子双节点
		return t.rightRotate(node)
	}

	return node
}

func (t *avlTree) Insert(val int) {
	t.root = t.insert(t.root, val)
}

func (t *avlTree) insert(node *hnode, val int) *hnode {
	if node == nil {
		return &hnode{
			height: 0,
			value:  val,
			left:   nil,
			right:  nil,
		}
	}

	if node.value < val {
		node.right = t.insert(node.right, val)
	} else if node.value > val {
		node.left = t.insert(node.left, val)
	} else {
		return node
	}

	t.updateHeight(node)
	node = t.rotate(node)
	return node
}

func (t *avlTree) Remove(val int) {
	t.root = t.remove(t.root, val)
}

func (t *avlTree) remove(root *hnode, val int) *hnode {
	if root == nil {
		return nil
	}

	if root.value < val {
		root.right = t.remove(root.right, val)
	} else if root.value > val {
		root.left = t.remove(root.left, val)
	} else {
		if root.left == nil || root.right == nil {
			child := root.left
			if child == nil {
				child = root.right
			}

			if child == nil {
				return nil
			}
			root = child
		} else {
			next := t.getInOrderNext(root.right)
			root.right = t.remove(root.right, next.value)
			root.value = next.value
		}
	}

	t.updateHeight(root)
	root = t.rotate(root)
	return root
}

// 获取中序遍历的下一个结点
func (t *avlTree) getInOrderNext(root *hnode) *hnode {
	if root == nil {
		return root
	}
	for root.left != nil {
		root = root.left
	}
	return root
}
