package leetcode

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
