/*
- @lc app=leetcode.cn id=501 lang=golang

给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：

给定 BST [1,null,2,2],
返回[2].

提示：如果众数超过1个，不需考虑输出顺序

进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
*/
package leetcode

import (
	"testing"
)

// @lc code=start
func findMode(root *TreeNode) []int {
	// return findModeDfs(root)
	return findModeIterate(root)
}

func findModeIterate(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode
	mostFreq := 1
	count := 1

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil {
			if top.Val == pre.Val {
				count++
			} else {
				if count >= mostFreq {
					if count > mostFreq {
						res = []int{}
						mostFreq = count
					}
					res = append(res, pre.Val)
				}
				count = 1 // 初始化计数器
			}
		}
		pre = top
		cur = top.Right
	}
	// 统计最后一个节点
	if count >= mostFreq {
		if count > mostFreq {
			res = []int{}
		}
		res = append(res, pre.Val)
	}
	return res
}

func findModeDfs(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	mostv := 1
	count := 1
	var pre *TreeNode
	var travseral func(r *TreeNode)
	travseral = func(r *TreeNode) {
		if r == nil {
			return
		}

		travseral(r.Left)
		if pre != nil {
			if r.Val == pre.Val {
				count++
			} else {
				if mostv <= count {
					if mostv < count {
						res = []int{}
						mostv = count
					}
					res = append(res, pre.Val)
				}
				count = 1
			}
		}
		pre = r
		travseral(r.Right)
	}
	travseral(root)
	// 统计最后一个节点
	if mostv <= count {
		if mostv < count {
			res = []int{}
		}
		res = append(res, pre.Val)
	}
	return res
}

func findModeExtraSpace(root *TreeNode) []int {
	// 没有利用二叉搜索树的有序性
	nums := make(map[int]int, 0)

	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			nums[queue[i].Val]++
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
	}

	maxv := 0
	res := make([]int, 0)
	for k, v := range nums {
		if v >= maxv {
			if v > maxv {
				res = []int{}
			}
			res = append(res, k)
		}
	}
	return res
}

// @lc code=end

func Test_findMode(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{"0", NewTreeByPreOrder([]any{}), []int{}},
		{"0", NewTreeByPreOrder([]any{1}), []int{1}},
		{"1", NewTreeByPreOrder([]any{1, 2, 3}), []int{1, 2, 3}},
		{"1", NewTreeByPreOrder([]any{2, 2, 3}), []int{2}},
		{"1", NewTreeByPreOrder([]any{2, 2, 2}), []int{2}},
		{"1", NewTreeByPreOrder([]any{10, 5, 15, 5, 6, 10, 15}), []int{5, 10, 15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMode(tt.root); !equalSet(got, tt.want) {
				t.Errorf("findMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
