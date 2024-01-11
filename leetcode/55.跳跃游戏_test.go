/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 *
 * https://leetcode.cn/problems/jump-game/description/
 *
 * algorithms
 * Medium (43.40%)
 * Likes:    2636
 * Dislikes: 0
 * Total Accepted:    846.7K
 * Total Submissions: 2M
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
 *
 * 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,3,1,1,4]
 * 输出：true
 * 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,2,1,0,4]
 * 输出：false
 * 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^4
 * 0 <= nums[i] <= 10^5
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func canJump(nums []int) bool {
	farest := 0
	end := len(nums) - 1
	for i := 0; i <= end && i <= farest; i++ {
		if i+nums[i] > farest {
			farest = i + nums[i]
		}
		if farest >= end {
			return true
		}
	}
	return false
}

// @lc code=end

func Test_canJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"1", []int{2, 3, 1, 1, 4}, true},
		{"2", []int{3, 2, 1, 0, 4}, false},
		{"3", []int{0, 2, 1, 0, 4}, false},
		{"4", []int{1, 1, 1, 1, 4}, true},
		{"5", []int{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
