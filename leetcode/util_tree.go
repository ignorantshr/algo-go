package leetcode

import "math"

type Node struct {
	Val      int
	Left     *Node
	Right    *Node
	Next     *Node
	Children []*Node
}

type PTreeNode struct {
	Val    int
	Left   *PTreeNode
	Right  *PTreeNode
	Parent *PTreeNode
}

func (t *PTreeNode) fillParent(root, parent *PTreeNode) {
	if root == nil {
		return
	}

	root.Parent = parent
	t.fillParent(root.Left, root)
	t.fillParent(root.Right, root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeByPreOrder(vals []any) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		var l, r *TreeNode
		if 2*i+1 < len(vals) {
			lv := vals[2*i+1]
			if lv != nil {
				l = &TreeNode{Val: lv.(int)}
			}
			nodes[2*i+1] = l
		}
		if 2*i+2 < len(vals) {
			rv := vals[2*i+2]
			if rv != nil {
				r = &TreeNode{Val: rv.(int)}
			}
			nodes[2*i+2] = r
		}
		if nodes[i] == nil {
			if v == nil {
				nodes[i] = nil
			} else {
				n := &TreeNode{
					Val:   v.(int),
					Left:  l,
					Right: r,
				}
				nodes[i] = n
			}
		} else {
			if 2*i+1 < len(vals) {
				nodes[i].Left = nodes[2*i+1]
			}
			if 2*i+2 < len(vals) {
				nodes[i].Right = nodes[2*i+2]
			}
		}
	}
	return nodes[0]
}

// 使用金字塔状构建二叉树
/*
		0			[0][]any{}
	1	    2		[1][]any{}
 3    4  5     6	[2][]any{}
*/
func NewTreeByPyramid(vals [][]any) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, 1<<len(vals)-1)
	for i, tier := range vals {
		for j, v := range tier {
			// 找到父节点
			loc := int(math.Pow(2, float64(i))) - 1 + j
			// 处理根节点特殊情况
			if loc == 0 {
				if v == nil {
					return nil
				}
				nodes[0] = &TreeNode{Val: v.(int)}
				continue
			}
			pnode := nodes[(loc-1)/2]
			// 没有父节点直接放弃
			if pnode == nil {
				continue
			}
			var node *TreeNode
			if v != nil {
				node = &TreeNode{Val: v.(int)}
			}
			if j&1 == 0 && v != nil {
				pnode.Left = node
			} else if j&1 == 1 && v != nil {
				pnode.Right = node
			}
			nodes[loc] = node
		}
	}
	return nodes[0]
}

func (t *TreeNode) equal(another *TreeNode) bool {
	var equals func(t1, t2 *TreeNode) bool
	equals = func(t1, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		}
		if t1 == nil || t2 == nil {
			return false
		}

		if t1.Val != t2.Val {
			return false
		}

		return equals(t1.Left, t2.Left) && equals(t1.Right, t2.Right)
	}
	return equals(t, another)
}

func (t *TreeNode) bfsPrefix() []any {
	res := make([]any, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, t)
	for len(queue) > 0 {
		size := len(queue)
		vals := make([]any, 0, size)
		for i := 0; i < size; i++ {
			ele := queue[i]
			if ele == nil {
				vals = append(vals, nil)
			} else {
				vals = append(vals, ele.Val)
				queue = append(queue, ele.Left, ele.Right)
			}
		}
		queue = queue[size:]
		allleaf := true
		for _, v := range vals {
			if v != nil {
				allleaf = false
				break
			}
		}
		if allleaf {
			break
		} else {
			res = append(res, vals...)
		}
	}
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] != nil {
			res = res[:i+1]
			break
		}
	}
	return res
}
