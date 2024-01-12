/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 *
 * https://leetcode.cn/problems/insert-interval/description/
 *
 * algorithms
 * Medium (42.64%)
 * Likes:    855
 * Dislikes: 0
 * Total Accepted:    188.9K
 * Total Submissions: 442.5K
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * 给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
 *
 * 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
 * 输出：[[1,5],[6,9]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * 输出：[[1,2],[3,10],[12,16]]
 * 解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
 *
 * 示例 3：
 *
 *
 * 输入：intervals = [], newInterval = [5,7]
 * 输出：[[5,7]]
 *
 *
 * 示例 4：
 *
 *
 * 输入：intervals = [[1,5]], newInterval = [2,3]
 * 输出：[[1,5]]
 *
 *
 * 示例 5：
 *
 *
 * 输入：intervals = [[1,5]], newInterval = [2,7]
 * 输出：[[1,7]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * intervals[i].length == 2
 * 0
 * intervals 根据 intervals[i][0] 按 升序 排列
 * newInterval.length == 2
 * 0
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	insert := false

	left := newInterval[0]
	right := newInterval[1]
	for _, cur := range intervals {
		if cur[1] < left {
			res = append(res, cur)
		} else if cur[0] > right {
			if insert {
				res = append(res, cur)
			} else {
				res = append(res, []int{left, right})
				res = append(res, cur)
				insert = true
			}
		} else {
			left = min(left, cur[0])
			right = max(right, cur[1])
		}
	}

	if !insert {
		res = append(res, []int{left, right})
	}

	return res
}

func insert1(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	size := len(intervals)
	insert := false

	if size == 0 {
		return [][]int{newInterval} // 没有区间可比
	}

	for i := 0; i < size; {
		if intervals[i][1] < newInterval[0] || intervals[i][0] > newInterval[1] {
			if insert { // 插入过了
				res = append(res, intervals[i:]...)
				break
			} else {
				// 优先插入小区间
				if intervals[i][1] < newInterval[0] {
					res = append(res, intervals[i])
					if i+1 == size { // 最后一个了，没有区间可比了
						res = append(res, newInterval)
					}
				} else {
					res = append(res, newInterval)
					res = append(res, intervals[i])
					insert = true
				}
				i++
			}
			continue
		}

		// 合并连续区间
		left := min(newInterval[0], intervals[i][0])
		right := max(newInterval[1], intervals[i][1])
		j := i + 1
		for j < size && intervals[j][0] <= right {
			right = max(right, intervals[j][1])
			j++
		}

		insert = true
		i = j
		res = append(res, []int{left, right})
	}

	return res
}

// @lc code=end

func Test_insert57(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// {"", args{[][]int{}, []int{}}, [][]int{}},
		{"x", args{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{"1", args{[][]int{{1, 3}, {6, 9}}, []int{2, 5}}, [][]int{{1, 5}, {6, 9}}},
		{"2", args{[][]int{{1, 5}}, []int{6, 7}}, [][]int{{1, 5}, {6, 7}}},
		{"3", args{[][]int{{3, 5}}, []int{1, 2}}, [][]int{{1, 2}, {3, 5}}},
		{"3.1", args{[][]int{{1, 2}, {7, 8}}, []int{4, 5}}, [][]int{{1, 2}, {4, 5}, {7, 8}}},
		{"4", args{[][]int{{1, 5}}, []int{1, 2}}, [][]int{{1, 5}}},
		{"5", args{[][]int{{3, 4}, {7, 10}}, []int{1, 5}}, [][]int{{1, 5}, {7, 10}}},
		{"6", args{[][]int{{3, 4}, {7, 10}}, []int{4, 7}}, [][]int{{3, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
