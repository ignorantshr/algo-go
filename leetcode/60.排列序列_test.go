/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 排列序列
 *
 * https://leetcode.cn/problems/permutation-sequence/description/
 *
 * algorithms
 * Hard (53.53%)
 * Likes:    829
 * Dislikes: 0
 * Total Accepted:    136.9K
 * Total Submissions: 255.5K
 * Testcase Example:  '3\n3'
 *
 * 给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
 *
 * 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
 *
 *
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 *
 *
 * 给定 n 和 k，返回第 k 个排列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3, k = 3
 * 输出："213"
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 4, k = 9
 * 输出："2314"
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 3, k = 1
 * 输出："123"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 9
 * 1 <= k <= n!
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func getPermutation(n int, k int) string {
	res := []byte{}
	fac := make([]int, n)
	valid := make([]bool, n+1)
	fac[0] = 1
	for i := 1; i < n; i++ {
		fac[i] = i * fac[i-1]
		valid[i] = true
	}
	valid[n] = true

	// 最左侧的待定数字可分为 (n-i)! 组
	for i := 1; i <= n; i++ {
		// 找到一个有效的数字
		order := (k-1)/fac[n-i] + 1
		for j := 1; j <= n; j++ {
			if valid[j] {
				order--
				if order == 0 {
					valid[j] = false
					res = append(res, '0'+byte(j))
					break
				}
			}
		}

		k = (k-1)%fac[n-i] + 1
	}

	return string(res)
}

func getPermutationBacktrack(n int, k int) string {
	s := []byte{}
	for i := 1; i <= n; i++ {
		s = append(s, '0'+byte(i))
	}

	path := []byte{}

	var traceback2 func()
	traceback2 = func() {
		if k == 0 {
			return
		}

		for i, v := range s {
			if v == 0 {
				continue
			}
			path = append(path, s[i])
			if len(path) == n {
				k--
			}
			s[i] = 0
			traceback2()
			if k <= 0 {
				return
			}
			s[i] = v
			path = path[:len(path)-1]
		}
	}
	traceback2()

	// var traceback func(start int)
	// traceback = func(start int) {
	// 	if k == 0 {
	// 		return
	// 	}

	// 	for i := start; i < n && k > 0; i++ {
	// 		path = append(path, s[i])
	// 		for j := i; j > start; j-- { // 上浮
	// 			s[j], s[j-1] = s[j-1], s[j]
	// 		}
	// 		if len(path) == n {
	// 			k--
	// 		}
	// 		traceback(start + 1)
	// 		if k <= 0 {
	// 			return
	// 		}
	// 		for j := start; j < i; j++ { // 下沉
	// 			s[j], s[j+1] = s[j+1], s[j]
	// 		}
	// 		path = path[:len(path)-1]
	// 	}
	// }
	// traceback(0)

	return string(path)
}

// @lc code=end

func Test_getPermutation(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{1, 1}, "1"},
		{"2.1", args{2, 1}, "12"},
		{"2.2", args{2, 2}, "21"},
		{"3.1", args{3, 1}, "123"},
		{"3.2", args{3, 2}, "132"},
		{"3.3", args{3, 3}, "213"},
		{"3.4", args{3, 4}, "231"},
		{"3.5", args{3, 5}, "312"},
		{"3.6", args{3, 6}, "321"},
		{"4.9", args{4, 9}, "2314"},
		{"4.19", args{4, 19}, "4123"},
		{"4.20", args{4, 20}, "4132"},
		{"4.24", args{4, 24}, "4321"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPermutation(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
