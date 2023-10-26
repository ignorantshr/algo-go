/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 *
 * https://leetcode.cn/problems/course-schedule-ii/description/
 *
 * algorithms
 * Medium (57.69%)
 * Likes:    893
 * Dislikes: 0
 * Total Accepted:    208.8K
 * Total Submissions: 362K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中
 * prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。
 *
 *
 * 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
 *
 *
 * 返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：[0,1]
 * 解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
 * 输出：[0,2,1,3]
 * 解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
 * 因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
 *
 * 示例 3：
 *
 *
 * 输入：numCourses = 1, prerequisites = []
 * 输出：[0]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= numCourses <= 2000
 * 0 <= prerequisites.length <= numCourses * (numCourses - 1)
 * prerequisites[i].length == 2
 * 0 <= ai, bi < numCourses
 * ai != bi
 * 所有[ai, bi] 互不相同
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	return findOrderBfs(numCourses, prerequisites)
	return findOrderDfs(numCourses, prerequisites)
}

func findOrderBfs(numCourses int, prerequisites [][]int) []int {
	table, indegree := build210Graph2(numCourses, prerequisites)
	queue := make([]int, 0)
	order := make([]int, 0)
	for i, v := range indegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	for {
		size := len(queue)
		if size == 0 {
			break
		}

		count += size
		for i := 0; i < size; i++ {
			order = append(order, queue[i])
			for _, v := range table[queue[i]] {
				indegree[v]--
				if indegree[v] == 0 {
					queue = append(queue, v)
				}
			}
		}
		queue = queue[size:]
	}

	if count != numCourses {
		return []int{}
	}

	return order
}

func findOrderDfs(numCourses int, prerequisites [][]int) []int {
	table := build210Graph(numCourses, prerequisites)
	visted := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	hasCycle := false
	postorder := make([]int, 0, numCourses)

	var cyclic func(idx int)
	cyclic = func(idx int) {
		if onPath[idx] {
			hasCycle = true
		}
		if visted[idx] || hasCycle {
			return
		}

		visted[idx] = true
		onPath[idx] = true
		for _, v := range table[idx] {
			cyclic(v)
		}
		onPath[idx] = false
		postorder = append(postorder, idx)
	}

	for i := 0; i < numCourses; i++ { // ⚠️这里需要遍历所有节点
		cyclic(i)
	}
	if hasCycle {
		return []int{}
	}

	// reverse postorder
	// 因为边的定义是 被依赖 关系，所以需要反转；如果边是按照 依赖 关系定义的，则不需要反转
	for i := 0; i < numCourses/2; i++ {
		postorder[i], postorder[numCourses-1-i] = postorder[numCourses-1-i], postorder[i]
	}
	return postorder
}

// 构建邻接表和每个节点的入度
func build210Graph2(numCourses int, prerequisites [][]int) ([][]int, []int) {
	table := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, cond := range prerequisites {
		if table[cond[1]] == nil {
			table[cond[1]] = make([]int, 0)
		}
		table[cond[1]] = append(table[cond[1]], cond[0])
		indegree[cond[0]]++
	}
	return table, indegree
}

// 构建邻接表
func build210Graph(numCourses int, prerequisites [][]int) [][]int {
	table := make([][]int, numCourses)
	for _, cond := range prerequisites {
		if table[cond[1]] == nil {
			table[cond[1]] = make([]int, 0)
		}
		table[cond[1]] = append(table[cond[1]], cond[0])
	}
	return table
}

// @lc code=end

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{1, [][]int{}}, []int{0}},
		{"", args{2, [][]int{{0, 1}}}, []int{1, 0}},
		{"", args{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}}, []int{0, 2, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
