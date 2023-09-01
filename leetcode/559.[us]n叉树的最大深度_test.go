/*
- @lc app=leetcode.cn id=559 lang=golang

给定一个 n 叉树，找到其最大深度。

最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
*/
package leetcode

// @lc code=start
// 递归，还可以使用层序遍历
func maxDepth559(root *Node) int {
	if root == nil {
		return 0
	}

	deep := 0
	for _, n := range root.Children {
		deep = max(deep, maxDepth559(n))
	}
	return deep + 1
}

// @lc code=end
