/*
- @lc app=leetcode.cn id=202 lang=golang

编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。

如果 n 是快乐数就返回 True ；不是，则返回 False 。

示例：

输入：19
输出：true
解释：
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
*/
package leetcode

import "testing"

// @lc code=start
func isHappy(n int) bool {
	return isHappySet(n)
}

func isHappySet(n int) bool {
	computed := make(map[int]bool)
	sum := n
	for sum != 1 {
		sum = sum202(sum)
		if computed[sum] {
			return false
		}
		computed[sum] = true
	}
	return true
}

func isHappyPointer(n int) bool {
	slow := n
	fast := n
	// 快慢指针
	for slow != 1 && fast != 1 {
		slow, fast = sum202(slow), sum202(sum202(fast))
		if slow == fast {
			return false
		}
	}
	return true
}

func sum202(n int) int {
	sum := 0
	for n > 0 {
		remains := n % 10
		sum += remains * remains
		n = n / 10
	}
	return sum
}

// @lc code=end

func Test_happyNum(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"1", 19, true},
		{"1", 1, true},
		{"1", 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.n); got != tt.want {
				t.Errorf("happyNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
