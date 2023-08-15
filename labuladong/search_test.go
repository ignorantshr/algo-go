package labuladong

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfsPrefix1(r *TreeNode, res *[]int) {
	if r == nil {
		return
	}
	*res = append(*res, r.Val)
	dfsPrefix1(r.Left, res)
	dfsPrefix1(r.Right, res)
}

func dfsPrefix2(r *TreeNode) []int {
	if r == nil {
		return []int{}
	}
	var res []int
	res = append(res, r.Val)
	res = append(res, dfsPrefix2(r.Left)...)
	res = append(res, dfsPrefix2(r.Right)...)
	return res
}
