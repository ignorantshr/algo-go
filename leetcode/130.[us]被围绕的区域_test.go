/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 *
 * https://leetcode.cn/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (46.28%)
 * Likes:    1054
 * Dislikes: 0
 * Total Accepted:    250.3K
 * Total Submissions: 540.9K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X'
 * 填充。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：board =
 * [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
 * 输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
 * 解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O'
 * 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
 *
 *
 * 示例 2：
 *
 *
 * 输入：board = [["X"]]
 * 输出：[["X"]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == board.length
 * n == board[i].length
 * 1
 * board[i][j] 为 'X' 或 'O'
 *
 */
package leetcode

import "testing"

// @lc code=start
func solve(board [][]byte) {
	// todo dfs
	solveUF(board)
}

func solveUF(board [][]byte) {
	m := len(board)
	n := len(board[0])

	// 坐标转下标：下标=i*n+j
	uf := InitUF130(m*n + 1)
	dummy := m * n // 所有边缘 0 及与其相连 0 的根结点

	// 首行和末行的 0
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			uf.Union(dummy, j)
		}
		if board[m-1][j] == 'O' {
			uf.Union(dummy, (m-1)*n+j)
		}
	}

	// 首列和末列的 0
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			uf.Union(dummy, i*n)
		}
		if board[i][n-1] == 'O' {
			uf.Union(dummy, i*n+n-1)
		}
	}

	orientation := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' {
				// ⚠️将此 O 与上下左右的 O 连通
				for k := 0; k < 4; k++ {
					x := i + orientation[k][0]
					y := j + orientation[k][1]
					if board[x][y] == 'O' {
						uf.Union(x*n+y, i*n+j)
					}
				}
			}
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' && !uf.Connected(dummy, i*n+j) {
				board[i][j] = 'X'
			}
		}
	}
}

// 并查集
type UF130 struct {
	parent []int // 父节点索引
}

func InitUF130(n int) *UF130 {
	u := &UF130{
		parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.parent[i] = -1
	}
	return u
}

func (u *UF130) Find(x int) int {
	r := x
	for u.parent[r] > 0 {
		r = u.parent[r]
	}

	for u.parent[x] > 0 {
		// tmp := u.parent[x]
		x, u.parent[x] = u.parent[x], r
		// x = tmp
	}

	return x
}

func (u *UF130) Union(x1, x2 int) {
	r1 := u.Find(x1)
	r2 := u.Find(x2)

	if r1 == r2 {
		return
	}

	u.parent[r1] = r2
}

func (u UF130) Connected(x1, x2 int) bool {
	return u.Find(x1) == u.Find(x2)
}

// @lc code=end

func Test_solve(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		want  [][]byte
	}{
		{"0", [][]byte{{'X'}}, [][]byte{{'X'}}},
		{"0", [][]byte{{'O'}}, [][]byte{{'O'}}},
		{"1", [][]byte{
			{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'},
		}, [][]byte{
			{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'},
		}},
		{"2", [][]byte{
			{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'O', 'X', 'X'},
		}, [][]byte{
			{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'O', 'X', 'X'},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if solve(tt.board); !equalSliceMatrix[byte](tt.board, tt.want) {
				t.Errorf("solve() = %v, want %v", tt.board, tt.want)
			}
		})
	}
}
