/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 *
 * https://leetcode.cn/problems/word-break/description/
 *
 * algorithms
 * Medium (54.26%)
 * Likes:    2311
 * Dislikes: 0
 * Total Accepted:    491.5K
 * Total Submissions: 901.8K
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
 *
 * 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: s = "leetcode", wordDict = ["leet", "code"]
 * 输出: true
 * 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
 *
 *
 * 示例 2：
 *
 *
 * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
 * 输出: true
 * 解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
 * 注意，你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 *
 * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出: false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 300
 * 1 <= wordDict.length <= 1000
 * 1 <= wordDict[i].length <= 20
 * s 和 wordDict[i] 仅由小写英文字母组成
 * wordDict 中的所有字符串 互不相同
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func wordBreak139(s string, wordDict []string) bool {
	// 分解子问题
	memo := make(map[int]bool)
	words := make(map[string]bool)
	for _, v := range wordDict {
		words[v] = true
	}

	var dp func(i int) bool
	dp = func(i int) bool {
		if i == len(s) {
			return true
		}
		// v1
		// if k > len(s) {
		// 	memo[k] = false
		// 	return false
		// }
		if whether, has := memo[i]; has {
			return whether
		}

		// 找到一个单词匹配 s[i:i+len)
		// 只要 s[i+len:] 可以被拼出，s[i:] 就能被拼出
		for l := 1; i+l <= len(s); l++ {
			if words[s[i:i+l]] {
				if dp(i + l) {
					return true
				}
			}
		}
		// v1
		// for _, word := range wordDict {
		// 	if strings.HasPrefix(s[k:], word) {
		// 		if dp(k + len(word)) {
		// 			return true
		// 		}
		// 	}
		// }
		memo[i] = false
		return false
	}
	return dp(0)
}

func wordBreakBackTrace(s string, wordDict []string) bool {
	//回溯，超时
	memo := make(map[string]bool)
	var backtrace func(path string) bool
	backtrace = func(path string) bool {
		if path == s {
			return true
		}
		if len(path) >= len(s) {
			memo[path] = false
			return false
		}

		if whether, has := memo[path]; has {
			return whether
		}
		for _, v := range wordDict {
			if backtrace(path + v) {
				return true
			}
		}
		memo[path] = false
		return false
	}
	return backtrace("")
}

// @lc code=end

func Test_wordBreak139(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"a", []string{"a"}}, true},
		{"1", args{"a", []string{"b"}}, false},
		{"2", args{"leetcode", []string{"leet", "code"}}, true},
		{"2", args{"applepenapple", []string{"apple", "pen"}}, true},
		{"2", args{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak139(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
