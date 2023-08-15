/*
* @lc app=leetcode.cn id=303 lang=golang
*
* [303] 区域和检索 - 数组不可变
*
* https://leetcode.cn/problems/range-sum-query-immutable/description/
*
  - algorithms
  - Easy (76.85%)
  - Likes:    565
  - Dislikes: 0
  - Total Accepted:    223.1K
  - Total Submissions: 290.1K
  - Testcase Example:  '["NumArray","sumRange","sumRange","sumRange"]\n' +
    '[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]'

*
* 给定一个整数数组  nums，处理以下类型的多个查询:
*
*
* 计算索引 left 和 right （包含 left 和 right）之间的 nums 元素的 和 ，其中 left <= right
*
*
* 实现 NumArray 类：
*
*
* NumArray(int[] nums) 使用数组 nums 初始化对象
* int sumRange(int i, int j) 返回数组 nums 中索引 left 和 right 之间的元素的 总和 ，包含 left 和
* right 两点（也就是 nums[left] + nums[left + 1] + ... + nums[right] )
*
*
*
*
* 示例 1：
*
*
* 输入：
* ["NumArray", "sumRange", "sumRange", "sumRange"]
* [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
* 输出：
* [null, 1, -1, -3]
*
* 解释：
* NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
* numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
* numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))
* numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))
*
*
*
*
* 提示：
*
*
* 1 <= nums.length <= 10^4
* -10^5 <= nums[i] <= 10^5
* 0 <= i <= j < nums.length
* 最多调用 10^4 次 sumRange 方法
*
*
*/
package leetcode

import "testing"

// @lc code=start
type NumArray struct {
	nums    []int
	preNums []int // 前缀和
}

func Constructor304(nums []int) NumArray {
	preNums := make([]int, len(nums)+1)
	for i, v := range nums {
		preNums[i+1] = preNums[i] + v
	}
	return NumArray{preNums: preNums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preNums[right+1] - this.preNums[left]
}

func Constructor1(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) SumRange1(left int, right int) int {
	sum := 0
	for left <= right {
		sum += this.nums[left]
		left++
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

func TestNumArray_SumRange(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		this NumArray
		args args
		want int
	}{
		{"1", Constructor304([]int{1}), args{0, 0}, 1},
		{"1", Constructor304([]int{1, 2, 3, 4, 5}), args{0, 0}, 1},
		{"1", Constructor304([]int{1, 2, 3, 4, 5}), args{0, 1}, 3},
		{"1", Constructor304([]int{-1, -1, 0, 3, 4, 5}), args{1, 4}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.SumRange(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("NumArray.SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
