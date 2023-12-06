/*
 * @lc app=leetcode.cn id=2646 lang=golang
 *
 * [2646] 最小化旅行的价格总和
 *
 * https://leetcode.cn/problems/minimize-the-total-price-of-the-trips/description/
 *
 * algorithms
 * Hard (51.59%)
 * Likes:    80
 * Dislikes: 0
 * Total Accepted:    9.4K
 * Total Submissions: 14.1K
 * Testcase Example:  '4\n[[0,1],[1,2],[1,3]]\n[2,2,10,6]\n[[0,3],[2,1],[2,3]]'
 *
 * 现有一棵无向、无根的树，树中有 n 个节点，按从 0 到 n - 1 编号。给你一个整数 n 和一个长度为 n - 1 的二维整数数组 edges
 * ，其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条边。
 *
 * 每个节点都关联一个价格。给你一个整数数组 price ，其中 price[i] 是第 i 个节点的价格。
 *
 * 给定路径的 价格总和 是该路径上所有节点的价格之和。
 *
 * 另给你一个二维整数数组 trips ，其中 trips[i] = [starti, endi] 表示您从节点 starti 开始第 i
 * 次旅行，并通过任何你喜欢的路径前往节点 endi 。
 *
 * 在执行第一次旅行之前，你可以选择一些 非相邻节点 并将价格减半。
 *
 * 返回执行所有旅行的最小价格总和。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 4, edges = [[0,1],[1,2],[1,3]], price = [2,2,10,6], trips =
 * [[0,3],[2,1],[2,3]]
 * 输出：23
 * 解释：
 * 上图表示将节点 2 视为根之后的树结构。第一个图表示初始树，第二个图表示选择节点 0 、2 和 3 并使其价格减半后的树。
 * 第 1 次旅行，选择路径 [0,1,3] 。路径的价格总和为 1 + 2 + 3 = 6 。
 * 第 2 次旅行，选择路径 [2,1] 。路径的价格总和为 2 + 5 = 7 。
 * 第 3 次旅行，选择路径 [2,1,3] 。路径的价格总和为 5 + 2 + 3 = 10 。
 * 所有旅行的价格总和为 6 + 7 + 10 = 23 。可以证明，23 是可以实现的最小答案。
 *
 * 示例 2：
 *
 * 输入：n = 2, edges = [[0,1]], price = [2,2], trips = [[0,0]]
 * 输出：1
 * 解释：
 * 上图表示将节点 0 视为根之后的树结构。第一个图表示初始树，第二个图表示选择节点 0 并使其价格减半后的树。
 * 第 1 次旅行，选择路径 [0] 。路径的价格总和为 1 。
 * 所有旅行的价格总和为 1 。可以证明，1 是可以实现的最小答案。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 50
 * edges.length == n - 1
 * 0 <= ai, bi <= n - 1
 * edges 表示一棵有效的树
 * price.length == n
 * price[i] 是一个偶数
 * 1 <= price[i] <= 1000
 * 1 <= trips.length <= 100
 * 0 <= starti, endi <= n - 1
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	table := make([][]int, n)
	for _, row := range edges {
		table[row[0]] = append(table[row[0]], row[1])
		table[row[1]] = append(table[row[1]], row[0])
	}

	// 统计途径 node 的数量
	count := make([]int, n)
	var dfs func(int, int, int) bool
	dfs = func(node, parent, end int) bool {
		if node == end {
			count[node]++
			return true
		}

		for _, child := range table[node] {
			if child == parent {
				continue
			}
			if dfs(child, node, end) {
				count[node]++
				return true
			}
		}
		return false
	}

	for _, trip := range trips {
		dfs(trip[0], -1, trip[1])
	}

	var dp func(node, parent int) []int
	dp = func(node, parent int) []int {
		res := []int{price[node] * count[node], price[node] * count[node] / 2}
		for _, child := range table[node] {
			if child == parent {
				continue
			}
			v := dp(child, node)
			// node 没有减半，因此可以取子树减半或没减半的最小值
			// node 减半了，只能取子树没有减半的情况
			res[0], res[1] = res[0]+min(v[0], v[1]), res[1]+v[0]
		}
		return res
	}

	res := dp(0, -1)
	return min(res[0], res[1])
}

// @lc code=end

func Test_minimumTotalPrice(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
		price []int
		trips [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{
			4,
			[][]int{{0, 1}, {1, 2}, {1, 3}},
			[]int{2, 2, 10, 6},
			[][]int{{0, 3}, {2, 1}, {2, 3}},
		}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotalPrice(tt.args.n, tt.args.edges, tt.args.price, tt.args.trips); got != tt.want {
				t.Errorf("minimumTotalPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
