/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode.cn/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (44.83%)
 * Likes:    1454
 * Dislikes: 0
 * Total Accepted:    489.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
 *
 * 返回这三个数的和。
 *
 * 假定每组输入只存在恰好一个解。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,2,1,-4], target = 1
 * 输出：2
 * 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,0,0], target = 1
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= nums.length <= 1000
 * -1000 <= nums[i] <= 1000
 * -10^4 <= target <= 10^4
 *
 *
 */
package leetcode

import (
	"math"
	"sort"
	"testing"
)

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	return threeSumClosest2(nums, target)
}

func threeSumClosest2(nums []int, target int) int {
	sort.Ints(nums)

	l := len(nums)
	best := math.MaxInt32
	for i := 0; i < l; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := l - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			if abs(target-sum) < abs(target-best) {
				best = sum
			}

			if sum > target {
				k0 := k - 1
				for j < k0 && nums[k0] == nums[k] {
					k0--
				}
				k = k0
			} else {
				j0 := j + 1
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}
		}
	}
	return best
}

// func abs(a int) int {
// 	if a > 0 {
// 		return a
// 	}
// 	return -a
// }

func threeSumClosest1(nums []int, target, start int) int {
	sort.Ints(nums)

	delta := math.MaxInt
	for i := 0; i < len(nums)-2; i++ {
		sum := nums[i] + twoSumClosest(nums, target-nums[i], i+1)
		if math.Abs(float64(target-sum)) < math.Abs(float64(delta)) {
			// delta = int(math.Abs(float64(target - sum)))
			delta = target - sum
		}
	}
	return target - delta
}

func twoSumClosest(nums []int, target, start int) int {
	delta := math.MaxInt
	lo, hi := start, len(nums)-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if math.Abs(float64(target-sum)) < math.Abs(float64(delta)) {
			// delta = int(math.Abs(float64(target - sum)))
			delta = target - sum
		}
		if sum < target {
			lo++
		} else {
			hi--
		}
	}
	return target - delta
}

// @lc code=end

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2}, -2},
		{"1", args{[]int{-1, 2, 1, -4}, 1}, 2},
		{"1", args{[]int{0, 0, 0}, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
