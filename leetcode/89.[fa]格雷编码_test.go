/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 *
 * https://leetcode.cn/problems/gray-code/description/
 *
 * algorithms
 * Medium (75.51%)
 * Likes:    661
 * Dislikes: 0
 * Total Accepted:    125.4K
 * Total Submissions: 166.3K
 * Testcase Example:  '2'
 *
 * n 位格雷码序列 是一个由 2^n 个整数组成的序列，其中：
 *
 * 每个整数都在范围 [0, 2^n - 1] 内（含 0 和 2^n - 1）
 * 第一个整数是 0
 * 一个整数在序列中出现 不超过一次
 * 每对 相邻 整数的二进制表示 恰好一位不同 ，且
 * 第一个 和 最后一个 整数的二进制表示 恰好一位不同
 *
 *
 * 给你一个整数 n ，返回任一有效的 n 位格雷码序列 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 2
 * 输出：[0,1,3,2]
 * 解释：
 * [0,1,3,2] 的二进制表示是 [00,01,11,10] 。
 * - 00 和 01 有一位不同
 * - 01 和 11 有一位不同
 * - 11 和 10 有一位不同
 * - 10 和 00 有一位不同
 * [0,2,3,1] 也是一个有效的格雷码序列，其二进制表示是 [00,10,11,01] 。
 * - 00 和 10 有一位不同
 * - 10 和 11 有一位不同
 * - 11 和 01 有一位不同
 * - 01 和 00 有一位不同
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[0,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 16
 *
 *
 */
package leetcode

import (
	"bytes"
	"math"
	"reflect"
	"strconv"
	"testing"
)

// @lc code=start
// 公式法
func grayCode(n int) []int {
	ans := make([]int, 1<<n)
	for i := range ans {
		ans[i] = i>>1 ^ i
	}
	return ans
}

// 回溯
func grayCodeBacktrack(n int) []int {
	ans := make([]int, 0)
	count := int(math.Pow(2, float64(n)))

	numstr := bytes.Repeat([]byte{'0'}, n)
	visited := make(map[string]bool)
	visited[string(numstr)] = true

	var backtrack func(k, idx int, nums []string) bool
	backtrack = func(k, idx int, nums []string) bool {
		if k == 0 {
			// check
			diff := 0
			for i := 0; i < n && diff <= 1; i++ {
				if nums[0][i] != nums[count-1][i] {
					diff++
				}
			}
			if diff == 1 {
				for _, v := range nums {
					num, _ := strconv.ParseInt(v, 2, 64)
					ans = append(ans, int(num))
				}
			}
			return diff == 1
		}

		for i := 0; i < n; i++ {
			numstr[i] = '1' - numstr[i] + '0'
			if !visited[string(numstr)] {
				visited[string(numstr)] = true
				nums = append(nums, string(numstr))
				if backtrack(k-1, i, nums) {
					return true
				}
				nums = nums[:len(nums)-1]
				visited[string(numstr)] = false
			}
			numstr[i] = '1' - numstr[i] + '0'
		}
		return false
	}

	backtrack(count-1, 0, []string{string(numstr)})

	return ans
}

// @lc code=end

func Test_grayCode(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"1", 1, []int{0, 1}},
		{"2", 2, []int{0, 1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grayCode(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grayCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
