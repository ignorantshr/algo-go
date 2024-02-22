package leetcode

import (
	"strconv"
	"testing"
)

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。


示例 1：

输入：x = 121
输出：true
示例 2：

输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：

输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。


提示：

-2^31<= x <= 2^31- 1


进阶：你能不将整数转为字符串来解决这个问题吗？

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isPalindrome9(x int) bool {
	return isPalindromeReverseHalf(x)
}

func isPalindromeStr(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	str := strconv.Itoa(x)
	for i, j := 0, len(str)-1; i < j; {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindromeStack(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	stack := []int{}
	for x > 0 {
		stack = append(stack, x%10)
		x = x / 10
	}
	for i, j := 0, len(stack)-1; i < j; {
		if stack[i] != stack[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindromeReverse(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	y := 0
	o := x
	for x > 0 {
		y = y*10 + x%10
		x = x / 10
	}
	return o == y
}

func isPalindromeReverseHalf(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	if x < 10 {
		return true
	}

	y := 0
	for x > y {
		y = y*10 + x%10
		x = x / 10
	}
	return x == y || x == y/10
}

func Test_isPalindrome9(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{"1", -11, false},
		{"1", 0, true},
		{"1", 4, true},
		{"2", 10, false},
		{"2", 123, false},
		{"3", 121, true},
		{"3", 1221, true},
		{"3", 12443534421, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome9(tt.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
