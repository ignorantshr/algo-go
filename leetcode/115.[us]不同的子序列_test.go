/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 *
 * https://leetcode.cn/problems/distinct-subsequences/description/
 *
 * algorithms
 * Hard (52.18%)
 * Likes:    1146
 * Dislikes: 0
 * Total Accepted:    153.3K
 * Total Submissions: 298.1K
 * Testcase Example:  '"rabbbit"\n"rabbit"'
 *
 * 给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数，结果需要对 10^9 + 7 取模。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "rabbbit", t = "rabbit"
 * 输出：3
 * 解释：
 * 如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
 * rabbbit
 * rabbbit
 * rabbbit
 *
 * 示例 2：
 *
 *
 * 输入：s = "babgbag", t = "bag"
 * 输出：5
 * 解释：
 * 如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。
 * babgbag
 * babgbag
 * babgbag
 * babgbag
 * babgbag
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length, t.length <= 1000
 * s 和 t 由英文字母组成
 *
 *
 */
package leetcode

import (
	"strconv"
	"testing"
)

// @lc code=start
func numDistinct(s string, t string) int {
	return numDistinct31(s, t)
	return numDistinct3(s, t)
	return numDistinct2(s, t)
	return numDistinct1(s, t)
}

// 动态规划，空间压缩
// 每次只用到了上一行的信息，所以就不需要存储更早的行数据了
func numDistinct31(s string, t string) int {
	ssize := len(s)
	tsize := len(t)
	// dp[i][j] = s[i]==t[j] ? {dp[i+1][j](不选s[i]) + dp[i+1][j+1](选s[i])} : {dp[i+1][j](不选s[i])}
	dp := make([]int, ssize+1) // dp[i][j] 是 s[i:] 的子序列中 t[j:] 出现的个数
	for i := 0; i <= ssize; i++ {
		dp[i] = 1
	}

	for j := tsize - 1; j >= 0; j-- {
		// 下面这两行看不懂
		pre := dp[ssize]
		dp[ssize] = 0 // 覆盖
		for i := ssize - 1; i >= 0; i-- {
			tmp := dp[i] // 暂存 dp[i+1][j+1]
			if s[i] == t[j] {
				dp[i] = dp[i+1] + pre
			} else {
				dp[i] = dp[i+1]
			}
			pre = tmp
		}
	}

	return dp[0]
}

// 动态规划
func numDistinct3(s string, t string) int {
	ssize := len(s)
	tsize := len(t)
	// dp[i][j] = s[i]==t[j] ? {dp[i+1][j](不选s[i]) + dp[i+1][j+1](选s[i])} : {dp[i+1][j](不选s[i])}
	dp := make([][]int, ssize+1) // dp[i][j] 是 s[i:] 的子序列中 t[j:] 出现的个数
	for i := range dp {
		dp[i] = make([]int, tsize+1)
		dp[i][tsize] = 1
	}
	// for i := tsize-1; i >= 0; i-- { // 省略
	// 	dp[ssize][i] = 0
	// }

	for i := ssize - 1; i >= 0; i-- {
		for j := tsize - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j] + dp[i+1][j+1]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	return dp[0][0]
}

// 递归回溯
func numDistinct2(s string, t string) int {
	ssize := len(s)
	tsize := len(t)
	count := 0
	memo := make(map[string]int, 0)

	var backtrack func(is, it int)
	backtrack = func(is, it int) {
		if it == tsize {
			count++
			return
		}
		if is == ssize {
			return
		}
		key := strconv.Itoa(is) + "," + strconv.Itoa(it)
		if v, has := memo[key]; has {
			count += v
			return
		}

		pre := count
		if s[is] == t[it] {
			backtrack(is+1, it+1) // 选
		}
		backtrack(is+1, it) // 不选
		memo[key] = count - pre
	}

	backtrack(0, 0)
	return count
}

// 递归分治
func numDistinct1(s string, t string) int {
	ssize := len(s)
	tsize := len(t)
	memo := make(map[string]int, 0)

	// s[is:] t[it:] => count
	var part func(is, it int) int
	part = func(is, it int) int {
		if it == tsize {
			return 1
		}
		if is == ssize {
			return 0
		}

		key := strconv.Itoa(is) + "," + strconv.Itoa(it)
		if v, has := memo[key]; has {
			return v
		}

		count := 0
		if s[is] == t[it] {
			// 选择 s[is]，或者不选 s[is]
			count = part(is+1, it+1) + part(is+1, it)
		} else {
			// 不选 s[is]
			count = part(is+1, it)
		}

		memo[key] = count
		return count
	}
	return part(0, 0)
}

// @lc code=end

func Test_numDistinct(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"1", "1"}, 1},
		{"1", args{"rabbbit", "rabbit"}, 3},
		{"1", args{"babgbag", "bag"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("numDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
