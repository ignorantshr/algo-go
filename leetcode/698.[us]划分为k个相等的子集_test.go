/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 *
 * https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/description/
 *
 * algorithms
 * Medium (42.07%)
 * Likes:    979
 * Dislikes: 0
 * Total Accepted:    107.1K
 * Total Submissions: 255.2K
 * Testcase Example:  '[4,3,2,3,5,2,1]\n4'
 *
 * 给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
 * 输出： True
 * 说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1,2,3,4], k = 3
 * 输出: false
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= len(nums) <= 16
 * 0 < nums[i] < 10000
 * 每个元素的频率在 [1,4] 范围内
 *
 *
 */
package leetcode

import (
	"sort"
	"testing"
)

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) < k {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum/k == 0 {
		return false
	}

	// return canPartitionKSubsetsByNum(nums, k, sum/k)
	return canPartitionKSubsetsByBucket(nums, k, sum/k)
}

// 从数字的视角来划分
func canPartitionKSubsetsByNum(nums []int, k, target int) bool {
	bucket := make([]int, k)
	// 降序排列，更容易触发剪枝操作，大的数字会先被分配到 bucket 中，对于之后的数字，bucket[i] + nums[idx] 会更大
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	var backtrack func(idx int) bool
	backtrack = func(idx int) bool {
		if idx == len(nums) {
			for j := 0; j < k; j++ {
				if bucket[j] != target {
					return false
				}
			}
			return true
		}

		// 选择放进 k 个桶中的一个
		for i := 0; i < k; i++ {
			// 剪枝
			if bucket[i]+nums[idx] > target {
				continue
			}

			bucket[i] += nums[idx]
			if backtrack(idx + 1) {
				return true
			}
			bucket[i] -= nums[idx]

			// nums[idx] 作为第一个加入桶里面的元素就失败了，那么加到其
			if bucket[i] == 0 {
				break
			}
		}
		return false
	}
	return backtrack(0)
}

// 从桶的视角来划分
func canPartitionKSubsetsByBucket(nums []int, k, target int) bool {
	bucket := make([]int, k)
	// choosed := make([]bool, len(nums)) // 每次转换会很耗时
	choosed := 0               // 使用位图来优化
	memo := make(map[int]bool) // 记录桶装满的情况，下次再遇到其他桶也是这种的话就可以直接返回
	var backtrack func(bucketIdx, numsIdx int) bool
	backtrack = func(bucketIdx, numsIdx int) bool {
		// 所有的桶都满了
		if bucketIdx == k {
			return true
		}
		// 此桶装满，装下一个
		if bucket[bucketIdx] == target {
			// 缓存结果
			if res, has := memo[choosed]; !has {
				return res
			}
			res := backtrack(bucketIdx+1, 0)
			memo[choosed] = res
			return res
		}

		// 对每个数做出选择
		for i := numsIdx; i < len(nums); i++ {
			if (choosed>>i)&1 == 1 {
				continue
			}
			if bucket[bucketIdx]+nums[i] > target {
				continue
			}

			choosed |= 1 << i
			bucket[bucketIdx] += nums[i]
			if backtrack(bucketIdx, i+1) {
				return true
			}
			choosed ^= 1 << i
			bucket[bucketIdx] -= nums[i]
		}
		return false
	}
	return backtrack(0, 0)
}

// @lc code=end

func Test_canPartitionKSubsets(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1}, 1}, true},
		{"1", args{[]int{2}, 1}, true},
		{"1", args{[]int{2}, 3}, false},
		{"1", args{[]int{4, 3, 2, 3, 5, 2, 1}, 4}, true},
		{"1", args{[]int{1, 2, 3, 4}, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartitionKSubsets(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("canPartitionKSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
