/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 *
 * https://leetcode.cn/problems/course-schedule/description/
 *
 * algorithms
 * Medium (53.87%)
 * Likes:    1813
 * Dislikes: 0
 * Total Accepted:    349.8K
 * Total Submissions: 649.3K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
 *
 * 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi]
 * ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
 *
 *
 * 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
 *
 *
 * 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：true
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
 *
 * 示例 2：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
 * 输出：false
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= numCourses <= 2000
 * 0 <= prerequisites.length <= 5000
 * prerequisites[i].length == 2
 * 0 <= ai, bi < numCourses
 * prerequisites[i] 中的所有课程对 互不相同
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 检测是不是无环图 DAG

	return canFinishBfs(numCourses, prerequisites)
	return canFinishDfs(numCourses, prerequisites)
}

func canFinishDfs(numCourses int, prerequisites [][]int) bool {
	table := build207Graph(numCourses, prerequisites)
	visted := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	hasCycle := false

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
	}

	for v := range table { // 这里遍历存在的根节点即可
		cyclic(v)
	}
	return !hasCycle
}

func canFinishBfs(numCourses int, prerequisites [][]int) bool {
	table, indegree := build207Graph2(numCourses, prerequisites)
	queue := make([]int, 0)
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
			for _, v := range table[queue[i]] { // 遍历相邻节点
				indegree[v]--
				if indegree[v] == 0 {
					queue = append(queue, v)
				}
			}
		}
		queue = queue[size:]
	}

	return count == numCourses
}

// 构建邻接表和每个节点的入度
func build207Graph2(numCourses int, prerequisites [][]int) ([][]int, []int) {
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
func build207Graph(numCourses int, prerequisites [][]int) [][]int {
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

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{2, [][]int{{1, 0}, {0, 1}}}, false},
		{"", args{2, [][]int{{1, 0}}}, true},
		{"", args{6, [][]int{{2, 1}, {3, 1}, {4, 3}, {5, 3}, {4, 2}}}, true},
		{"", args{6, [][]int{{2, 1}, {3, 1}, {4, 3}, {5, 3}, {5, 2}, {3, 2}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
