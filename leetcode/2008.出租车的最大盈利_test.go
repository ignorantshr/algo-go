/*
 * @lc app=leetcode.cn id=2008 lang=golang
 *
 * [2008] 出租车的最大盈利
 *
 * https://leetcode.cn/problems/maximum-earnings-from-taxi/description/
 *
 * algorithms
 * Medium (47.64%)
 * Likes:    114
 * Dislikes: 0
 * Total Accepted:    11.7K
 * Total Submissions: 22K
 * Testcase Example:  '5\n[[2,5,4],[1,5,1]]'
 *
 * 你驾驶出租车行驶在一条有 n 个地点的路上。这 n 个地点从近到远编号为 1 到 n ，你想要从 1 开到 n
 * ，通过接乘客订单盈利。你只能沿着编号递增的方向前进，不能改变方向。
 *
 * 乘客信息用一个下标从 0 开始的二维数组 rides 表示，其中 rides[i] = [starti, endi, tipi] 表示第 i
 * 位乘客需要从地点 starti 前往 endi ，愿意支付 tipi 元的小费。
 *
 * 每一位 你选择接单的乘客 i ，你可以 盈利 endi - starti + tipi 元。你同时 最多 只能接一个订单。
 *
 * 给你 n 和 rides ，请你返回在最优接单方案下，你能盈利 最多 多少元。
 *
 * 注意：你可以在一个地点放下一位乘客，并在同一个地点接上另一位乘客。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 5, rides = [[2,5,4],[1,5,1]]
 * 输出：7
 * 解释：我们可以接乘客 0 的订单，获得 5 - 2 + 4 = 7 元。
 *
 *
 * 示例 2：
 *
 * 输入：n = 20, rides =
 * [[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]
 * 输出：20
 * 解释：我们可以接以下乘客的订单：
 * - 将乘客 1 从地点 3 送往地点 10 ，获得 10 - 3 + 2 = 9 元。
 * - 将乘客 2 从地点 10 送往地点 12 ，获得 12 - 10 + 3 = 5 元。
 * - 将乘客 5 从地点 13 送往地点 18 ，获得 18 - 13 + 1 = 6 元。
 * 我们总共获得 9 + 5 + 6 = 20 元。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^5
 * 1 <= rides.length <= 3 * 10^4
 * rides[i].length == 3
 * 1 <= starti < endi <= n
 * 1 <= tipi <= 10^5
 *
 *
 */
package leetcode

import (
	"sort"
	"testing"
)

// @lc code=start
func maxTaxiEarnings(n int, rides [][]int) int64 {
	sort.Slice(rides, func(i, j int) bool {
		return rides[i][1] < rides[j][1]
	})
	cn := len(rides)
	dp := make([]int64, cn) // 接单前 i 位乘客的最大盈利
	dp[0] = int64(rides[0][1] - rides[0][0] + rides[0][2])
	for i := 1; i < cn; i++ {
		pre := findEnd(rides, rides[i][0])
		chose := int64(rides[i][1] - rides[i][0] + rides[i][2])
		if pre != -1 { // 有行程不冲突的乘客
			chose += dp[pre]
		}
		// 不选乘客/选择乘客
		dp[i] = max(dp[i-1], chose)
	}

	return dp[cn-1]
}

// 找到终点 <= end 的位置
func findEnd(rides [][]int, end int) int {
	cn := len(rides)
	l := 0
	r := cn - 1

	for l <= r {
		mid := (r + l) / 2
		if rides[mid][1] < end {
			l = mid + 1
		} else if rides[mid][1] > end {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l - 1

	// for i := 0; i < len(rides); i++ {
	// 	if rides[i][1] > end {
	// 		return i - 1
	// 	}
	// }
	// return -1
}

// @lc code=end

func Test_maxTaxiEarnings(t *testing.T) {
	type args struct {
		n     int
		rides [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"1", args{
			5,
			[][]int{{2, 5, 4}, {1, 5, 1}},
		}, 7},
		{"2", args{
			20,
			[][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}},
		}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTaxiEarnings(tt.args.n, tt.args.rides); got != tt.want {
				t.Errorf("maxTaxiEarnings() = %v, want %v", got, tt.want)
			}
		})
	}
}
