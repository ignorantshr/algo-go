package leetcode

import (
	"testing"
)

/*
给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。



示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestPalindrome(s string) string {
	return longestPalindromeCenterDiffuse(s)
}

func longestPalindromeCenterDiffuse(s string) string {
	idxl, idxr := 0, 0
	for i := 0; i < len(s)-1; i++ {
		l1, r1 := diameter(s, i, i)
		l2, r2 := diameter(s, i, i+1)
		if r1-l1 > idxr-idxl {
			idxl, idxr = l1, r1
		}
		if r2-l2 > idxr-idxl {
			idxl, idxr = l2, r2
		}
	}
	return s[idxl : idxr+1]
}

func diameter(s string, idxl, idxr int) (int, int) {
	for idxl >= 0 && idxr < len(s) && s[idxl] == s[idxr] {
		idxl--
		idxr++
	}
	return idxl + 1, idxr - 1
}

func longestPalindromeDP(s string) string {
	matrix := make([][]bool, len(s))
	for i := range matrix {
		matrix[i] = make([]bool, len(s))
		matrix[i][i] = true
	}

	left := 0
	maxl := 0
	if len(matrix) > 0 {
		maxl = 1
	}
	// 没必要从 1 开始，因为初始化时已经考虑了 1 的情况
	for l := 2; l <= len(s); l++ {
		for i := 0; i <= len(s)-l; i++ {
			j := i + l - 1
			if s[j] == s[i] && (l == 2 || matrix[i+1][j-1]) {
				matrix[i][j] = true
				if l > maxl {
					maxl = l
					left = i
				}
			}
		}
	}
	return s[left : left+maxl]
}

func TestLongestPalindrome(t *testing.T) {
	table := []struct {
		name string
		in   string
		out  string
	}{
		{
			"1",
			"babad",
			"bab",
		},
		{
			"2",
			"bab",
			"bab",
		},
		{
			"3",
			"babada",
			"bab",
		},
		{
			"4",
			"a",
			"a",
		},
		{
			"5",
			"cbbd",
			"bb",
		},
		{
			"6",
			"cs",
			"c",
		},
	}

	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			if re := longestPalindrome(v.in); v.out != re {
				t.Fatalf("in:%v\nout:%v\nexpect:%v\n", v.in, re, v.out)
			}
		})
	}
}
