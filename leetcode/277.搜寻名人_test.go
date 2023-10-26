/*
- @lc app=leetcode.cn id=277 lang=golang

给你 n 个人的社交关系（你知道任意两个人之间是否认识），然后请你找出这些人中的「名人」。

所谓「名人」有两个条件：

1、所有其他人都认识「名人」。

2、「名人」不认识任何其他人。
*/
package leetcode

// @lc code=start
// 请你实现：返回「名人」的编号
func findCelebrity(n int) int {
	return findCelebrity3(n)
	return findCelebrity2(n)
	return findCelebrity1(n)
}

func findCelebrity3(n int) int {
	// 把队列优化掉，每次只保留可能是名人的那个
	i := 0
	for other := 1; other < n; other++ {
		if knows(i, other) || !knows(other, i) {
			// pass i
			i = other
		} else {
			// pass other
		}
	}

	for other := 0; other < n; other++ {
		// 你认识其它人 或 有人不认识你
		if other != i && (knows(i, other) || !knows(other, i)) {
			return -1
		}
	}
	return i
}

func findCelebrity2(n int) int {
	candidates := make([]int, 0, n)
	for i := 0; i < n; i++ {
		candidates = append(candidates, i)
	}

	for len(candidates) >= 2 {
		p1 := candidates[len(candidates)-1]
		p2 := candidates[len(candidates)-2]
		candidates = candidates[:len(candidates)-2]

		if knows(p1, p2) || !knows(p2, p1) {
			// p1 knows p2 or p2 doesn't know p1
			// pass p1
			candidates = append(candidates, p2)
		} else {
			// p1 doesn't knows p2 and p2 know p1
			// pass p2
			candidates = append(candidates, p1)
		}
	}

	i := candidates[0]
	for other := 0; other < n; other++ {
		// 你认识其它人 或 有人不认识你
		if other != i && (knows(i, other) || !knows(other, i)) {
			return -1
		}
	}
	return i
}

// 暴力解法
func findCelebrity1(n int) int {
	for i := 0; i < n; i++ {
		other := 0
		for ; other < n; other++ {
			// 你认识其它人 或 有人不认识你
			if other != i && (knows(i, other) || !knows(other, i)) {
				break
			}
		}
		if other == n {
			return i
		}
	}

	return -1
}

// @lc code=end

// 可以直接调用，能够返回 i 是否认识 j
func knows(i int, j int) bool {
	return false
}
