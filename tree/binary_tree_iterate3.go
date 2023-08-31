package tree

/* 迭代遍历v3，统一风格 */

// 前序遍历
func preorderTraversal3(root *node) []int {
	if root == nil {
		return nil
	}

	var res []int
	cur := root
	stack := make([]*node, 0)
	stack = append(stack, cur)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top != nil {
			// 重新按序将节点入栈
			if top.right != nil {
				stack = append(stack, top.right) // 右节点
			}

			if top.left != nil {
				stack = append(stack, top.left) // 左节点
			}

			stack = append(stack, top) // 中节点
			stack = append(stack, nil) // 标记前一个节点，标记为待处理节点
		} else {
			top = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, top.value)
		}
	}
	return res
}

// 中序遍历
func inorderTraversal3(root *node) []int {
	if root == nil {
		return nil
	}

	var res []int
	cur := root
	stack := make([]*node, 0)
	stack = append(stack, cur)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top != nil {
			// 重新按序将节点入栈
			if top.right != nil {
				stack = append(stack, top.right) // 右节点
			}

			stack = append(stack, top) // 中节点
			stack = append(stack, nil) // 标记前一个节点，标记为待处理节点

			if top.left != nil {
				stack = append(stack, top.left) // 左节点
			}
		} else {
			top = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, top.value)
		}
	}
	return res
}

// 后序遍历
func postorderTraversal3(root *node) []int {
	if root == nil {
		return nil
	}

	var res []int
	cur := root
	stack := make([]*node, 0)
	stack = append(stack, cur)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top != nil {
			// 重新按序将节点入栈
			stack = append(stack, top) // 中节点
			stack = append(stack, nil) // 标记前一个节点，标记为待处理节点

			if top.right != nil {
				stack = append(stack, top.right) // 右节点
			}

			if top.left != nil {
				stack = append(stack, top.left) // 左节点
			}
		} else {
			top = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, top.value)
		}
	}
	return res
}
