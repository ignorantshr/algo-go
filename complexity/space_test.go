package complexity

import "fmt"

// 空间复杂度
// https://www.hello-algo.com/chapter_computational_complexity/space_complexity/#231

func function() {
}

/* 常数阶 */
func spaceConstant(n int) {
	const a = 0
	b := 0

	// 栈帧空间」用于保存调用函数的上下文数据。系统每次调用函数都会在栈的顶部创建一个栈帧，函数返回时，栈帧空间会被释放。
	// 此种情况是 O(1)
	for i := 0; i < n; i++ {
		function()
	}

	// 循环中的变量占用 O(1) 空间
	var c int
	for i := 0; i < n; i++ {
		c = 0
	}
	fmt.Println(a, b, c)
}

/* 线性阶 */
func spaceLinear(n int) {
	_ = make([]int, n)

	var node []int
	for i := 0; i < n; i++ {
		node = append(node, n)
	}

	m := make(map[int]bool, 0)
	for i := 0; i < n; i++ {
		m[i] = true
	}
}

/* 线性阶（递归实现） */
func spaceLinearRecur(n int) {
	if n == 1 {
		return
	}
	spaceLinearRecur(n - 1)
}

/* 平方阶 */
func spaceQuadratic(n int) {
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}
}

/* 平方阶（递归实现） */
func spaceQuadraticRecur(n int) {
	if n <= 1 {
		return
	}
	_ = make([]int, 0)
	spaceLinearRecur(n - 1)
}

type treeNode struct {
	left  *treeNode
	right *treeNode
}

/* 指数阶（建立满二叉树） */
func buildTree(n int) *treeNode {
	if n <= 0 {
		return nil
	}
	root := &treeNode{}
	root.left = buildTree(n - 1)
	root.right = buildTree(n - 1)
	return root
}

/* 对数阶（循环实现）*/
func spaceLogarithmic(n int) {
	for n > 1 {
		_ = make([]int, n)
		n /= 2
	}
}

/* 对数阶（递归实现）*/
func spaceLogRecur(n int) {
	if n <= 1 {
		return
	}
	_ = make([]int, n)
	spaceLogRecur(n / 2)
}
