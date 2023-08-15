/*
 * @lc app=leetcode.cn id=1109 lang=golang
 *
 * [1109] 航班预订统计
 *
 * https://leetcode.cn/problems/corporate-flight-bookings/description/
 *
 * algorithms
 * Medium (63.62%)
 * Likes:    471
 * Dislikes: 0
 * Total Accepted:    112.5K
 * Total Submissions: 176.6K
 * Testcase Example:  '[[1,2,10],[2,3,20],[2,5,25]]\n5'
 *
 * 这里有 n 个航班，它们分别从 1 到 n 进行编号。
 *
 * 有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从
 * firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
 *
 * 请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
 * 输出：[10,55,45,25,25]
 * 解释：
 * 航班编号        1   2   3   4   5
 * 预订记录 1 ：   10  10
 * 预订记录 2 ：       20  20
 * 预订记录 3 ：       25  25  25  25
 * 总座位数：      10  55  45  25  25
 * 因此，answer = [10,55,45,25,25]
 *
 *
 * 示例 2：
 *
 *
 * 输入：bookings = [[1,2,10],[2,2,15]], n = 2
 * 输出：[10,25]
 * 解释：
 * 航班编号        1   2
 * 预订记录 1 ：   10  10
 * 预订记录 2 ：       15
 * 总座位数：      10  25
 * 因此，answer = [10,25]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 2 * 10^4
 * 1 <= bookings.length <= 2 * 10^4
 * bookings[i].length == 3
 * 1 <= firsti <= lasti <= n
 * 1 <= seatsi <= 10^4
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func corpFlightBookings(bookings [][]int, n int) []int {
	differ := newDifferArray(make([]int, n))
	for _, b := range bookings {
		differ.update(b[0]-1, b[1]-1, b[2])
	}
	return differ.restore()
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

func Test_corpFlightBookings(t *testing.T) {
	type args struct {
		bookings [][]int
		n        int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0", args{
			[][]int{},
			0,
		}, []int{}},
		{"1", args{
			[][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}},
			5,
		}, []int{10, 55, 45, 25, 25}},
		{"1", args{
			[][]int{{1, 2, 10}, {2, 2, 15}},
			2,
		}, []int{10, 25}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := corpFlightBookings(tt.args.bookings, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("corpFlightBookings() = %v, want %v", got, tt.want)
			}
		})
	}
}
