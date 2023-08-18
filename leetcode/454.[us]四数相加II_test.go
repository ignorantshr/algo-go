/*
- @lc app=leetcode.cn id=454 lang=golang

给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。

为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -2^28 到 2^28 - 1 之间，最终结果不会超过 2^31 - 1 。

例如:

输入:

A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]
输出:

2

解释:

两个元组如下:

(0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
(1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
*/
package leetcode

import "testing"

// @lc code=start

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	res := 0
	statistics := make(map[int]int)
	for _, va := range nums1 {
		for _, vb := range nums2 {
			statistics[va+vb]++
		}
	}

	for _, vc := range nums3 {
		for _, vd := range nums4 {
			res += statistics[0-vc-vd]
		}
	}

	return res
}

// @lc code=end

func Test_fourSumCount(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{
			[]int{},
			[]int{},
			[]int{},
			[]int{},
		}, 0},
		{"1", args{
			[]int{1, 2},
			[]int{-2, -1},
			[]int{-1, 2},
			[]int{0, 2},
		}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSumCount(tt.args.nums1, tt.args.nums2, tt.args.nums3, tt.args.nums4); got != tt.want {
				t.Errorf("fourSumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
