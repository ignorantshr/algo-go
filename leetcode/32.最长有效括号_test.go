/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 *
 * https://leetcode.cn/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (37.43%)
 * Likes:    2411
 * Dislikes: 0
 * Total Accepted:    411.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '"(()"'
 *
 * 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "(()"
 * 输出：2
 * 解释：最长有效括号子串是 "()"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = ")()())"
 * 输出：4
 * 解释：最长有效括号子串是 "()()"
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = ""
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 3 * 104
 * s[i] 为 '(' 或 ')'
 *
 *
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func longestValidParentheses(s string) int {
	// 贪心算法
	// 遇到没有括号 和 ) 匹配的情况肯定就不匹配了，重新计数
	// 但是这样无法判断 (() 的情况，所以需要再反向遍历一次
	maxl := 0

	left, rigth := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			rigth++
		}

		if left == rigth {
			maxl = max(maxl, 2*left)
		} else if rigth > left {
			left = 0
			rigth = 0
		}
	}

	left, rigth = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			rigth++
		}

		if left == rigth {
			maxl = max(maxl, 2*left)
		} else if left > rigth {
			left = 0
			rigth = 0
		}
	}

	return maxl
}

func longestValidParenthesesDp(s string) int {
	maxl := 0
	dp := make([]int, len(s)) // dp[i] 表示以下标 i 字符结尾的最长有效括号的长度
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				// ()
				dp[i] = 2
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				// ...(...sub...))
				// 以前一个字符结尾的子串长度 加上 匹配成功时匹配的左括号之前的子串长度
				dp[i] = dp[i-1] + 2
				if i-dp[i-1] >= 2 {
					dp[i] += dp[i-dp[i-1]-2]
				}
			}

			if dp[i] > maxl {
				maxl = dp[i]
			}
		}
	}

	return maxl
}

func longestValidParenthesesStack(s string) int {
	maxl := 0
	serialStack := make([]int, 0, len(s)/2)
	serialStack = append(serialStack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			serialStack = append(serialStack, i)
			continue
		}

		serialStack = serialStack[:len(serialStack)-1]
		if len(serialStack) == 0 {
			serialStack = append(serialStack, i)
		} else {
			maxl = max(maxl, i-serialStack[len(serialStack)-1])
		}
	}

	return maxl
}

func longestValidParenthesesMe(s string) int {
	// 给每个括号编号，最后计算栈中留下的最长间隔
	charStack := make([]byte, 0, len(s)/2)
	serialStack := make([]int, 0, len(s)/2)

	for i := 0; i < len(s); i++ {
		if len(charStack) == 0 || s[i] == '(' {
			charStack = append(charStack, s[i])
			serialStack = append(serialStack, i)
			continue
		}

		if s[i] == ')' {
			if charStack[len(charStack)-1] == '(' {
				charStack = charStack[:len(charStack)-1]
				serialStack = serialStack[:len(serialStack)-1]
			} else {
				charStack = append(charStack, s[i])
				serialStack = append(serialStack, i)
			}
		}
	}

	maxl := 0
	var span int
	for i := 0; i < len(serialStack); i++ {
		if i == 0 {
			span = serialStack[i]
		} else {
			span = serialStack[i] - serialStack[i-1] - 1
		}
		if span > maxl {
			maxl = span
		}
	}
	if len(serialStack) > 0 {
		span = len(s) - 1 - serialStack[len(serialStack)-1]
		if span > maxl {
			maxl = span
		}
	} else {
		maxl = len(s)
	}

	return maxl
}

// @lc code=end

func Test_longestValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"0", "", 0},
		{"1", "(()", 2},
		{"2", ")()())", 4},
		{"3", ")(()())", 6},
		{"4", ")((()()))", 8},
		{"5", "((()()))", 8},
		{"6", "()", 2},
		{"7", "()())", 4},
		{"8", "()(()", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
			if got := longestValidParenthesesDp(tt.s); got != tt.want {
				t.Errorf("longestValidParenthesesDp() = %v, want %v", got, tt.want)
			}
			if got := longestValidParenthesesStack(tt.s); got != tt.want {
				t.Errorf("longestValidParenthesesStack() = %v, want %v", got, tt.want)
			}
			if got := longestValidParenthesesMe(tt.s); got != tt.want {
				t.Errorf("longestValidParenthesesMe() = %v, want %v", got, tt.want)
			}
		})
	}
}
