/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 *
 * https://leetcode.cn/problems/decode-ways/description/
 *
 * algorithms
 * Medium (33.33%)
 * Likes:    1489
 * Dislikes: 0
 * Total Accepted:    292.3K
 * Total Submissions: 873K
 * Testcase Example:  '"12"'
 *
 * 一条包含字母 A-Z 的消息通过以下映射进行了 编码 ：
 *
 *
 * 'A' -> "1"
 * 'B' -> "2"
 * ...
 * 'Z' -> "26"
 *
 * 要 解码 已编码的消息，所有数字必须基于上述映射的方法，反向映射回字母（可能有多种方法）。例如，"11106" 可以映射为：
 *
 *
 * "AAJF" ，将消息分组为 (1 1 10 6)
 * "KJF" ，将消息分组为 (11 10 6)
 *
 *
 * 注意，消息不能分组为  (1 11 06) ，因为 "06" 不能映射为 "F" ，这是由于 "6" 和 "06" 在映射中并不等价。
 *
 * 给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数 。
 *
 * 题目数据保证答案肯定是一个 32 位 的整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "12"
 * 输出：2
 * 解释：它可以解码为 "AB"（1 2）或者 "L"（12）。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "226"
 * 输出：3
 * 解释：它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "06"
 * 输出：0
 * 解释："06" 无法映射到 "F" ，因为存在前导零（"6" 和 "06" 并不等价）。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 100
 * s 只包含数字，并且可能包含前导零。
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1) // s[,i) 前 i 个字符的组合情况
	dp[0] = 1
	size := len(s)

	for i := 1; i <= size; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1] // ..1
		}
		if i > 1 && s[i-2] != '0' && s[i-2:i] <= "26" { // ..10
			dp[i] += dp[i-2]
		}
	}

	return dp[size]
}

// 超时
func numDecodingsDFS(s string) int {
	if s[0] == '0' {
		return 0
	}

	size := len(s)
	ans := 0
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == size {
			ans++
			return
		}

		if s[idx] == '0' {
			return
		}

		dfs(idx + 1)
		if idx+1 < size && s[idx:idx+2] <= "26" {
			dfs(idx + 2)
		}
	}
	dfs(0)

	return ans
}

// @lc code=end

func Test_numDecodings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"x.4", "30", 0},
		{"x.3", "301", 0},
		{"x.2", "230", 0},
		{"x.1", "111111111111111111111111111111111111111111111", 1836311903},
		{"0", "06", 0},
		{"0.1", "600", 0},
		{"0.2", "6002", 0},
		{"1", "11106", 2},
		{"1.1", "12", 2},
		{"1.2", "226", 3},
		{"1.3", "32", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
			// if got := numDecodingsDFS(tt.s); got != tt.want {
			// 	t.Errorf("numDecodingsDFS() = %v, want %v", got, tt.want)
			// }
		})
	}
}
