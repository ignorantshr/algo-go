/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode.cn/problems/permutations/description/
 *
 * algorithms
 * Medium (78.90%)
 * Likes:    2689
 * Dislikes: 0
 * Total Accepted:    927.1K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1]
 * 输出：[[0,1],[1,0]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 6
 * -10 <= nums[i] <= 10
 * nums 中的所有整数 互不相同
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var choose func(path []int, options []int, l int)
	choose = func(path, options []int, l int) {
		if len(options) == l {
			// fmt.Println("result", path, options[l:])
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		// fmt.Println("choose a path", path, options[l:])
		for i := l; i < len(options); i++ {
			path = append(path, options[i])
			options[i], options[l] = options[l], options[i]
			choose(path, options, l+1)
			options[l], options[i] = options[i], options[l]
			path = path[:len(path)-1]
		}
	}

	choose([]int{}, nums, 0)
	return res
}

// @lc code=end

func Test_permute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"1", []int{1}, [][]int{{1}}},
		{"1", []int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{"1", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.nums); !equalSliceMatrix(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
