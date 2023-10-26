/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径
 *
 * https://leetcode.cn/problems/all-paths-from-source-to-target/description/
 *
 * algorithms
 * Medium (78.85%)
 * Likes:    438
 * Dislikes: 0
 * Total Accepted:    115K
 * Total Submissions: 145.8K
 * Testcase Example:  '[[1,2],[3],[3],[]]'
 *
 * 给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）
 *
 * graph[i] 是一个从节点 i 可以访问的所有节点的列表（即从节点 i 到节点 graph[i][j]存在一条有向边）。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：graph = [[1,2],[3],[3],[]]
 * 输出：[[0,1,3],[0,2,3]]
 * 解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
 * 输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == graph.length
 * 2 <= n <= 15
 * 0 <= graph[i][j] < n
 * graph[i][j] != i（即不存在自环）
 * graph[i] 中的所有元素 互不相同
 * 保证输入为 有向无环图（DAG）
 *
 *
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	size := len(graph)
	path := make([]int, 0)

	var dfs func(idx int)
	dfs = func(idx int) {
		path = append(path, idx)
		if idx == size-1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			path = path[:len(path)-1] // 这里 return 要正确维护 path
			return
		}

		for _, i := range graph[idx] {
			dfs(i)
		}
		path = path[:len(path)-1]
	}
	dfs(0)

	return res
}

// @lc code=end

func Test_allPathsSourceTarget(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  [][]int
	}{
		{"", [][]int{{}}, [][]int{{0}}},
		{"", [][]int{{1, 2}, {3}, {3}, {}}, [][]int{{0, 1, 3}, {0, 2, 3}}},
		{"", [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}, [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPathsSourceTarget(tt.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
