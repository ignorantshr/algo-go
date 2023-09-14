package leetcode

func insertIntoBSTRV2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val < val {
		root.Right = insertIntoBSTRV2(root.Right, val)
	} else {
		root.Left = insertIntoBSTRV2(root.Left, val)
	}
	return root
}

func insertIntoBSTRV1(root *TreeNode, val int) *TreeNode {
	var pre *TreeNode
	cur := root
	for cur != nil {
		pre = cur
		if cur.Val < val {
			cur = cur.Right
		} else if cur.Val > val {
			cur = cur.Left
		}
	}
	if pre == nil {
		return &TreeNode{Val: val}
	}
	if pre.Val < val {
		pre.Right = &TreeNode{Val: val}
	} else {
		pre.Left = &TreeNode{Val: val}
	}
	return root
}
