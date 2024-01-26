/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode.cn/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.45%)
 * Likes:    3383
 * Dislikes: 0
 * Total Accepted:    743.2K
 * Total Submissions: 959.3K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：["()"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 8
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	path := make([]byte, 0)

	var backtrack func(l, r int)
	backtrack = func(l, r int) {
		if l == n && r == n {
			res = append(res, string(path))
			return
		}

		if l < n {
			path = append(path, '(')
			backtrack(l+1, r)
			path = path[:len(path)-1]
		}
		if r < l && r < n {
			path = append(path, ')')
			backtrack(l, r+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)
	return res
}

// @lc code=end

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{"1", 1, []string{"()"}},
		{"2", 2, []string{"()()", "(())"}},
		{"3", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.n); !equalSet(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
