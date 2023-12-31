package leetcode

import (
	"testing"
)

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func xlengthOfLongestSubstring(s string) int {
	visited := make(map[byte]int)

	length := 0
	left, right, size := 0, 0, len(s)
	for right < size {
		c := s[right]
		visited[c]++
		right++

		for visited[c] > 1 {
			if length < right-left {
				length = right - left - 1
			}
			visited[s[left]]--
			left++
		}
	}
	return max(length, right-left)
}

func xlengthOfLongestSubstring1(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	maxl := 0
	slow, fast := 0, 1
	visited := make(map[byte]int)
	visited[s[slow]] = 1

	for fast < len(s) {
		pos := visited[s[fast]]
		if pos == 0 {
			visited[s[fast]] = fast + 1
			fast++
		} else {
			if maxl < fast-slow {
				maxl = fast - slow
			}
			for i := slow; i < pos; i++ {
				delete(visited, s[slow])
				slow++
			}
		}
	}
	return max(maxl, fast-slow)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(xlengthOfLongestSubstring("abbcabcbb"))
	t.Log(xlengthOfLongestSubstring("abc"))
	t.Log(xlengthOfLongestSubstring("a"))
	t.Log(xlengthOfLongestSubstring("dvdf"))
	t.Log(xlengthOfLongestSubstring("dede"))
	t.Log(xlengthOfLongestSubstring("deed"))
}
