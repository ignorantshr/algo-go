/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 *
 * https://leetcode.cn/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Medium (53.46%)
 * Likes:    1979
 * Dislikes: 0
 * Total Accepted:    544.4K
 * Total Submissions: 1M
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
 *
 * 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
 *
 * 示例 1：
 *
 * 输入：nums = [100,4,200,1,3,2]
 * 输出：4
 * 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
 *
 * 示例 2：
 *
 * 输入：nums = [0,3,7,2,5,8,4,6,0,1]
 * 输出：9
 *
 * 提示：
 *
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 *
 */
package leetcode

import "testing"

// @lc code=start
func longestConsecutive(nums []int) int {
	return longestConsecutiveSet(nums)
	return longestConsecutiveUF(nums)
}

func longestConsecutiveSet(nums []int) int {
	set := make(map[int]struct{})
	for _, v := range nums {
		set[v] = struct{}{}
	}

	ans := 0
	for k := range set {
		if _, ok := set[k-1]; ok {
			continue
		}
		cur := k
		for {
			if _, ok := set[cur]; ok {
				cur++
			} else {
				break
			}
		}
		ans = max(ans, cur-k)
	}

	return ans
}

// 并查集，时间复杂度不是 O(n)
func longestConsecutiveUF(nums []int) int {
	set := make(map[int]int)
	idx := 1
	for _, v := range nums {
		if set[v] == 0 {
			set[v] = idx
			idx++
		}
	}

	uf := newuf(len(set))
	for k := range set {
		if v := set[k-1]; v != 0 {
			uf.union(v-1, set[k]-1)
		}
	}

	ans := 0
	for _, v := range uf.size {
		ans = max(ans, v)
	}

	return ans
}

type uf struct {
	parent []int
	size   []int // 每个联通区间的节点数量
}

func newuf(size int) *uf {
	u := &uf{make([]int, size), make([]int, size)}
	for i := 0; i < size; i++ {
		u.parent[i] = i
		u.size[i] = 1
	}
	return u
}

func (u *uf) findRoot(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.findRoot(u.parent[x])
	}
	return u.parent[x]
}

func (u *uf) union(x, y int) {
	xr := u.findRoot(x)
	yr := u.findRoot(y)

	if xr == yr {
		return
	}

	u.parent[yr] = xr
	u.size[xr] += u.size[yr]
}

// @lc code=end

func Test_longestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"0", []int{}, 0},
		{"1", []int{1}, 1},
		{"1.1", []int{1, 4}, 1},
		{"2", []int{3, 4}, 2},
		{"2.1", []int{100, 4, 200, 1, 3, 2}, 4},
		{"2.2", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
			if got := longestConsecutiveUF(tt.nums); got != tt.want {
				t.Errorf("longestConsecutiveUF() = %v, want %v", got, tt.want)
			}
		})
	}
}
