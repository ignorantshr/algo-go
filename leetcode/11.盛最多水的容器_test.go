package leetcode

import "testing"

/*
11. 盛最多水的容器
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

示例 1：

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例 2：

输入：height = [1,1]
输出：1

提示：

n == height.length
2 <= n <= 105
0 <= height[i] <= 104

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
假设状态 S(i,j) 下 h[i]<h[j] ，在向内移动短板至 S(i+1,j) ，
则相当于消去了 S(i,j−1),S(i,j−2),...,S(i,i+1) 状态集合。而所有消去状态的面积一定都小于当前面积（即 <S(i,j)），因为这些状态：

短板高度：相比 S(i,j) 相同或更短（即 ≤h[i] ）；
底边宽度：相比 S(i,j) 更短；

作者：jyd
链接：https://leetcode.cn/problems/container-with-most-water/solution/container-with-most-water-shuang-zhi-zhen-fa-yi-do/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func maxArea(height []int) int {
	// 双指针，排除小值
	m := 0
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			m = max(m, height[i]*(j-i))
			i++
		} else {
			m = max(m, height[j]*(j-i))
			j--
		}
	}
	return m
}

// 超时
func maxAreaDp(height []int) int {
	// dp[i] = max(area(height[i-1]), dp[i-1])
	dp := make([]int, len(height)+1)
	dp[1] = 0
	for i := 2; i <= len(height); i++ {
		m := dp[i-1]
		for j := 1; j < i; j++ {
			a := min(height[i-1], height[j-1]) * (i - j)
			m = max(m, a)
		}
		dp[i] = m
	}
	return dp[len(height)]
}

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"2", []int{1, 1}, 1},
		{"3", []int{1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
