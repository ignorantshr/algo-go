/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 *
 * https://leetcode.cn/problems/word-break-ii/description/
 *
 * algorithms
 * Hard (57.40%)
 * Likes:    726
 * Dislikes: 0
 * Total Accepted:    93.4K
 * Total Submissions: 161.6K
 * Testcase Example:  '"catsanddog"\n["cat","cats","and","sand","dog"]'
 *
 * 给定一个字符串 s 和一个字符串字典 wordDict ，在字符串 s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序
 * 返回所有这些可能的句子。
 *
 * 注意：词典中的同一个单词可能在分段中被重复使用多次。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
 * 输出:["cats and dog","cat sand dog"]
 *
 *
 * 示例 2：
 *
 *
 * 输入:s = "pineapplepenapple", wordDict =
 * ["apple","pen","applepen","pine","pineapple"]
 * 输出:["pine apple pen apple","pineapple pen apple","pine applepen apple"]
 * 解释: 注意你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 *
 * 输入:s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
 * 输出:[]
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 *
 * 1 <= s.length <= 20
 * 1 <= wordDict.length <= 1000
 * 1 <= wordDict[i].length <= 10
 * s 和 wordDict[i] 仅有小写英文字母组成
 * wordDict 中所有字符串都 不同
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	words := make(map[string]bool)
	for _, v := range wordDict {
		words[v] = true
	}

	memo := make(map[int][]string) // all sub strings int s[i:]

	var dp func(i int) []string
	dp = func(i int) []string {
		res := make([]string, 0)
		if i == len(s) {
			return res
		}

		if substrings := memo[i]; substrings != nil {
			return substrings
		}

		// s[i:]所有组合 = prefix + s[i+l:]所有组合
		for l := 1; i+l <= len(s); l++ {
			prefix := s[i : i+l]
			if words[prefix] { // 找到一种组合
				substrings := dp(i + l)
				if len(substrings) == 0 && i+l == len(s) { // 字符串末尾
					res = append(res, prefix)
				} else {
					for _, str := range substrings {
						res = append(res, prefix+" "+str)
					}
				}
			}
		}
		memo[i] = res
		return res
	}
	return dp(0)
}

// @lc code=end

func Test_wordBreak140(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{
			"catsandog",
			[]string{"cat", "cats", "and", "sand", "dog"},
		}, []string{}},
		{"", args{
			"catsanddog",
			[]string{"cat", "cats", "and", "sand", "dog"},
		}, []string{"cats and dog", "cat sand dog"}},
		{"", args{
			"pineapplepenapple",
			[]string{"apple", "pen", "applepen", "pine", "pineapple"},
		}, []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); !equalSet(got, tt.want) {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
