/*
 * @lc app=leetcode.cn id=1094 lang=golang
 *
 * [1094] 拼车
 *
 * https://leetcode.cn/problems/car-pooling/description/
 *
 * algorithms
 * Medium (51.73%)
 * Likes:    265
 * Dislikes: 0
 * Total Accepted:    75.8K
 * Total Submissions: 146.8K
 * Testcase Example:  '[[2,1,5],[3,3,7]]\n4'
 *
 * 车上最初有 capacity 个空座位。车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向）
 *
 * 给定整数 capacity 和一个数组 trips ,  trip[i] = [numPassengersi, fromi, toi] 表示第 i
 * 次旅行有 numPassengersi 乘客，接他们和放他们的位置分别是 fromi 和 toi 。这些位置是从汽车的初始位置向东的公里数。
 *
 * 当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：trips = [[2,1,5],[3,3,7]], capacity = 4
 * 输出：false
 *
 *
 * 示例 2：
 *
 *
 * 输入：trips = [[2,1,5],[3,3,7]], capacity = 5
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= trips.length <= 1000
 * trips[i].length == 3
 * 1 <= numPassengersi <= 100
 * 0 <= fromi < toi <= 1000
 * 1 <= capacity <= 10^5
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func carPooling(trips [][]int, capacity int) bool {
	differ := newDifferArray(make([]int, 1001))
	for _, t := range trips {
		// t[2] 时已下车
		differ.update(t[1], t[2]-1, t[0])
	}
	for _, c := range differ.restore() {
		if c > capacity {
			return false
		}
	}
	return true
}

// // 差分数组
// type differArray struct {
// 	differs []int // differs[i] = nums[i] - nums[i-1], differs[0] = nums[0]
// }

// func newDifferArray(nums []int) *differArray {
// 	size := len(nums)
// 	if size == 0 {
// 		return &differArray{differs: []int{}}
// 	}

// 	differs := make([]int, size)
// 	differs[0] = nums[0]
// 	for i := 1; i < size; i++ {
// 		differs[i] = nums[i] - nums[i-1]
// 	}
// 	return &differArray{differs: differs}
// }

// // 用 val 更新原数组的 [a,b] 区间
// func (d *differArray) update(a, b, val int) {
// 	if a >= 0 && a < len(d.differs) {
// 		d.differs[a] += val
// 	}
// 	if b >= 0 && b < len(d.differs)-1 {
// 		d.differs[b+1] -= val
// 	}
// }

// // 恢复原数组
// func (d *differArray) restore() []int {
// 	size := len(d.differs)
// 	if size == 0 {
// 		return []int{}
// 	}

// 	nums := make([]int, size)
// 	nums[0] = d.differs[0]
// 	for i := 1; i < size; i++ {
// 		nums[i] = d.differs[i] + nums[i-1]
// 	}
// 	return nums
// }

// @lc code=end

func Test_carPooling(t *testing.T) {
	type args struct {
		trips    [][]int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{
			[][]int{},
			0,
		}, true},
		{"1", args{
			[][]int{{2, 1, 5}, {3, 3, 7}},
			5,
		}, true},
		{"1", args{
			[][]int{{2, 1, 5}, {3, 3, 7}},
			4,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carPooling(tt.args.trips, tt.args.capacity); got != tt.want {
				t.Errorf("carPooling() = %v, want %v", got, tt.want)
			}
		})
	}
}
