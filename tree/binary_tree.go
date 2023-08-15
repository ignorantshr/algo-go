package tree

import "container/list"

// 广度优先遍历
func bfs(root TreeNode) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		n := queue.Remove(queue.Front()).(TreeNode)
		print(n, " ")
		if !n.Left().isEmpty() {
			queue.PushBack(n.Left())
		}
		if !n.Right().isEmpty() {
			queue.PushBack(n.Right())
		}
	}
}

// 深度优先遍历，0: 前序, 1: 中序，2：后序
func dfs(root TreeNode, order int) {
	if root.isEmpty() {
		return
	}

	if order == 0 {
		print(root, " ")
	}

	if !root.Left().isEmpty() {
		dfs(root.Left(), order)
	}
	if order == 1 {
		print(root, " ")
	}
	if !root.Right().isEmpty() {
		dfs(root.Right(), order)
	}
	if order == 2 {
		print(root, " ")
	}
}

func dfsNodes(root *node, order int) []*node {
	if root.isEmpty() {
		return nil
	}

	res := make([]*node, 0)
	if order == 0 {
		res = append(res, root)
	}

	if !root.left.isEmpty() {
		res = append(res, dfsNodes(root.left, order)...)
	}
	if order == 1 {
		res = append(res, root)
	}
	if !root.right.isEmpty() {
		res = append(res, dfsNodes(root.right, order)...)
	}
	if order == 2 {
		res = append(res, root)
	}
	return res
}

// 迭代版本
func dfsIterate(root *node) []int {
	var res []int
	stack := make([]*node, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			// 前序遍历位置
			stack = append(stack, root)
			root = root.left
		}
		peek := stack[len(stack)-1]
		// 中序遍历位置
		stack = stack[:len(stack)-1]
		root = peek.right
	}
	return res
}

// 后序遍历 迭代版本
func dfsPostorderTraversal(root *node) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	stack := []*node{}
	var lastVisited *node

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}

		peek := stack[len(stack)-1]

		// 如果当前节点没有右子节点或者上一次访问的节点是当前节点的右子节点，那么我们就可以访问当前节点。
		// 否则，说明当前节点的右子树还未被完全遍历，需要先处理右子树。
		if peek.right != nil && lastVisited != peek.right {
			root = peek.right
		} else {
			result = append(result, peek.value)
			lastVisited = peek
			stack = stack[:len(stack)-1]
		}
	}

	return result
}

// mirrorsDfsInOrder 是一个 mirrors 遍历算法，用于进行中序遍历
func mirrorsDfsInOrder(root *node) []int {
	var res []int
	cur := root

	for cur != nil {
		if cur.left == nil { // 如果当前节点的左子树为空
			res = append(res, cur.value) // 将当前节点的值添加到结果集中
			cur = cur.right              // 移动到右子节点
		} else {
			pre := cur.left
			if pre.right != nil && pre.right != cur { // 找到前驱节点
				pre = pre.right
			}

			if pre.right == nil { // 还没有访问 pre 的右子树
				pre.right = cur // 将 pre 的右子节点指向 cur，构成线索，在遍历完左子树后将回到 cur
				cur = cur.left  // 移动到左子节点
			} else { // 已经访问过 pre 的右子树
				pre.right = nil              // 恢复原始树的结构
				res = append(res, cur.value) // 将当前节点的值添加到结果集中
				cur = cur.right              // 移动到右子节点
			}
		}
	}

	return res // 返回结果集
}
