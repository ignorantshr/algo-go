package tree

/* 另一种思路的二叉树的迭代遍历方法，比较容易理解 */

// 前序遍历
func preorderTraversal2(root *node) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := []*node{}

	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.value)

		if top.right != nil { // 右节点线入栈，因为出栈时要保持 中左右的顺序
			stack = append(stack, top.right)
		}
		if top.left != nil {
			stack = append(stack, top.left)
		}
	}

	return res
}

// 后序遍历
func postorderTraversal2(root *node) []int {
	res := make([]int, 0)
	stack := []*node{} // 栈处理节点
	cur := root        // 指针遍历节点

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		var parent *node
		if len(stack) > 1 {
			parent = stack[len(stack)-2]
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.value)

		// 判断父节点的右节点是不是就是自己：
		// 		是的话就不要被入栈两次重复处理
		// 		不是的话就入栈等待处理
		if parent != nil && parent.right != top {
			// stack = append(stack, parent.right) 右节点入栈，移到了签名的for循环中
			cur = parent.right
		}
	}
	return res
}
