/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 *
 * https://leetcode.cn/problems/jump-game-ii/description/
 *
 * algorithms
 * Medium (44.82%)
 * Likes:    2387
 * Dislikes: 0
 * Total Accepted:    595.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
 *
 * 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j]
 * 处:
 *
 *
 * 0 <= j <= nums[i]
 * i + j < n
 *
 *
 * 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [2,3,1,1,4]
 * 输出: 2
 * 解释: 跳到最后一个位置的最小跳跃数是 2。
 * 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [2,3,0,1,4]
 * 输出: 2
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 10^4
 * 0 <= nums[i] <= 1000
 * 题目保证可以到达 nums[n-1]
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func jump(nums []int) int {
	// 贪心算法
	// 每次根据自己的可能情况找经过下一个距离的中介后跳的最远的
	size := len(nums)
	step := 0
	far := 0 // 遍历过的路径中能跳的最远距离
	end := 0 // 边界

	for i := 0; i < size-1; i++ {
		far = max(far, i+nums[i])
		if i == end {
			// 每次到达边界的时候就需要在遍历过的路径中选择一个最远的来突破边界，否则就只能在边界里面待着
			// s[...从 s 可以跳到的位置...] 从 s 起跳的话跳一步最远只能在这里了
			// s[.................(++++]+++++) 想突破边界限制就必须再跳一次
			end = far
			step++
		}
	}

	/* 通俗易懂写法
	for i := 0; i < size-1; {
		most := 0
		nidx := i
		for j := 1; j <= nums[i]; j++ { // 遍历下一跳的情况，选择跳的最远的
			if i+j >= size-1 {
				return step + 1
			}

			if most < nums[i+j]+j {
				nidx = i + j
				most = nums[i+j] + j
			}
		}
		step++
		i = nidx
	} */

	return step
}

func jumpGreed1(nums []int) int {
	// 倒序贪心算法
	pos := len(nums) - 1
	step := 0
	for pos > 0 {
		for i := 0; i < pos; i++ {
			if i+nums[i] >= pos { // 找到距离终点最远的那个位置，也就是索引最小的位置
				pos = i
				step++
				break
			}
		}
	}
	return step
}

// 很慢
func jumpDp(nums []int) int {
	size := len(nums)
	// dp = min(1+dp[j]) j+nums[j]>=i
	dp := make([]int, size) // dp[i] 以 i 索引结尾的最小步数
	// dp[0] = 0
	for i := 1; i < size; i++ {
		dp[i] = 10000
	}

	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[size-1]
}

// @lc code=end

func Test_jump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"0", []int{2}, 0},
		{"1", []int{2, 3, 1, 1, 4}, 2},
		{"2", []int{2, 3, 0, 1, 4}, 2},
		{"3", []int{1, 1, 1, 1}, 3},
		{"4", []int{2, 2, 0, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
			if got := jumpDp(tt.nums); got != tt.want {
				t.Errorf("jumpDp() = %v, want %v", got, tt.want)
			}
		})
	}
}
