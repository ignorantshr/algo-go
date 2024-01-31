/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 *
 * https://leetcode.cn/problems/interleaving-string/description/
 *
 * algorithms
 * Medium (44.64%)
 * Likes:    965
 * Dislikes: 0
 * Total Accepted:    134.4K
 * Total Submissions: 300.3K
 * Testcase Example:  '"aabcc"\n"dbbca"\n"aadbbcbcac"'
 *
 * 给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。
 *
 * 两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
 *
 *
 * s = s1 + s2 + ... + sn
 * t = t1 + t2 + ... + tm
 * |n - m| <= 1
 * 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 +
 * ...
 *
 *
 * 注意：a + b 意味着字符串 a 和 b 连接。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：s1 = "", s2 = "", s3 = ""
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s1.length, s2.length <= 100
 * 0 <= s3.length <= 200
 * s1、s2、和 s3 都由小写英文字母组成
 *
 *
 *
 *
 * 进阶：您能否仅使用 O(s2.length) 额外的内存空间来解决它?
 *
 */
package leetcode

import "testing"

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	size1 := len(s1)
	size2 := len(s2)
	size3 := len(s3)
	if size1+size2 != size3 {
		return false
	}

	dp := make([][]bool, size1+1) // dp[i][j] 以 s1[:1) s2[:2) 结尾是否构成交错字符串
	for i := 0; i <= size1; i++ {
		dp[i] = make([]bool, size2+1)
	}
	dp[0][0] = true

	for i := 0; i <= size1; i++ {
		for j := 0; j <= size2; j++ {
			// 这里为什么要带 dp[i][j] || 前缀呢，因为不带会把 dp[0][0] = true 给覆盖掉
			dp[i][j] = dp[i][j] || (i > 0 && s1[i-1] == s3[i+j-1] && dp[i-1][j]) ||
				(j > 0 && s2[j-1] == s3[i+j-1] && dp[i][j-1])
		}
	}

	return dp[size1][size2]
}

// timeout
func isInterleaveBacktrack(s1 string, s2 string, s3 string) bool {
	size1 := len(s1)
	size2 := len(s2)
	size3 := len(s3)
	if size1+size2 != size3 {
		return false
	}

	var backtrack func(i1, i2, i3, who int) bool
	backtrack = func(i1, i2, i3, who int) bool {
		v1 := i1
		v2 := i2
		v3 := i3

		if v1 == size1 && v2 == size2 && v3 == size3 {
			return true
		}

		if who == 1 {
			for v1 < size1 && s1[v1] == s3[v3] {
				v1++
				v3++
			}
			for i := i1 + 1; i <= v1; i++ {
				if backtrack(i, v2, i3+i-i1, 2) {
					return true
				}
			}
			if v1 == i1 || (v3 == size3 && (v2 != size2 || v1 != size1)) {
				return false
			}
			return backtrack(v1, v2, v3, 2)
		} else {
			for v2 < size2 && s2[v2] == s3[v3] {
				v2++
				v3++
			}
			for i := i2 + 1; i <= v2; i++ {
				if backtrack(v1, i, i3+i-i2, 1) {
					return true
				}
			}
			if v2 == i2 || (v3 == size3 && (v2 != size2 || v1 != size1)) {
				return false
			}
			return backtrack(v1, v2, v3, 1)
		}
	}

	if backtrack(0, 0, 0, 1) {
		return true
	}
	return backtrack(0, 0, 0, 2)
}

// @lc code=end

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{}, true},
		{"1", args{"a", "b", "ab"}, true},
		{"1.1", args{"a", "a", "aa"}, true},
		{"1.2", args{"aa", "a", "aaa"}, true},
		{"1.3", args{"ab", "c", "abc"}, true},
		{"1.4", args{"ab", "c", "cab"}, true},
		{"1.5", args{"aabcc", "dbbca", "aadbbcbcac"}, true},
		{"1.6", args{"aabcc", "dbbca", "aadbbbaccc"}, false},
		{"1.7", args{"a", "c", "ab"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}
