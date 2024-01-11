/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode.cn/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (49.64%)
 * Likes:    2214
 * Dislikes: 0
 * Total Accepted:    766.7K
 * Total Submissions: 1.5M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi]
 * 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
 * 输出：[[1,6],[8,10],[15,18]]
 * 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[1,4],[4,5]]
 * 输出：[[1,5]]
 * 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= intervals.length <= 10^4
 * intervals[i].length == 2
 * 0 <= starti <= endi <= 10^4
 *
 *
 */
package leetcode

import (
	"sort"
	"testing"
)

// @lc code=start
func merge(intervals [][]int) [][]int {
	// 双指针
	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	for i := 0; i < len(intervals); {
		j := i + 1
		most := intervals[i][1]
		for ; j < len(intervals) && intervals[j][0] <= most; j++ {
			most = max(most, intervals[j][1])
		}
		res = append(res, []int{intervals[i][0], most})
		i = j
	}
	return res
}

func mergeIterator(intervals [][]int) [][]int {
	// 按照起点排序
	// 与上一个区间对比，合并

	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		last := res[len(res)-1]
		if cur[0] <= last[1] {
			last[1] = max(cur[1], last[1])
			res[len(res)-1] = last
		} else {
			res = append(res, cur)
		}
	}

	return res
}

// @lc code=end

func Test_merge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{"1", [][]int{{1, 3}}, [][]int{{1, 3}}},
		{"2", [][]int{{1, 3}, {2, 6}}, [][]int{{1, 6}}},
		{"3", [][]int{{2, 6}, {8, 10}}, [][]int{{2, 6}, {8, 10}}},
		{"4", [][]int{{2, 11}, {8, 10}}, [][]int{{2, 11}}},
		{"5", [][]int{{2, 11}, {1, 10}}, [][]int{{1, 11}}},
		{"6", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := make([][]int, len(tt.intervals))
			copy(dst, tt.intervals)
			if got := merge(tt.intervals); !equalSliceMatrix(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
			if got := mergeIterator(dst); !equalSliceMatrix(got, tt.want) {
				t.Errorf("mergeIterator() = %v, want %v", got, tt.want)
			}
		})
	}
}
