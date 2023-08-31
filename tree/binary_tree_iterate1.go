package tree

/* 迭代遍历v1 */

// 前序遍历
func preorderTraversal1(root *node) []int {
	var res []int
	stack := make([]*node, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			// 前序遍历位置
			// 遍历和处理位置一致
			res = append(res, root.value)
			stack = append(stack, root)
			root = root.left
		}
		peek := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = peek.right
	}
	return res
}

// 中序遍历
func inorderTraversal1(root *node) []int {
	res := make([]int, 0)
	stack := []*node{} // 栈处理节点
	cur := root        // 指针遍历节点

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.value)
		cur = top.right // 父节点遍历完成后马上就遍历右节点
	}
	return res
}

// 后序遍历
func postorderTraversal1(root *node) []int {
	result := []int{}
	stack := []*node{}
	var lastVisited *node

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}

		peek := stack[len(stack)-1]

		// 如果:
		// 		当前节点没有右子节点（不需要处理右节点）
		// 或者
		// 		上一次访问的节点是当前节点的右子节点（右节点已经处理过了）
		// 那么我们就可以访问当前节点（中节点）。
		// 否则，说明当前节点的右子树还未被遍历，需要先处理右子树。
		if peek.right == nil || lastVisited == peek.right {
			result = append(result, peek.value)
			stack = stack[:len(stack)-1]
			lastVisited = peek
		} else {
			root = peek.right
		}
	}

	return result
}
