/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 *
 * https://leetcode.cn/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (39.36%)
 * Likes:    1051
 * Dislikes: 0
 * Total Accepted:    183.2K
 * Total Submissions: 467.2K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * 给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。
 *
 * s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。
 *
 *
 * 例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"，
 * "cdefab"，"efabcd"， 和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。
 *
 *
 * 返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "barfoothefoobarman", words = ["foo","bar"]
 * 输出：[0,9]
 * 解释：因为 words.length == 2 同时 words[i].length == 3，连接的子字符串的长度必须为 6。
 * 子串 "barfoo" 开始位置是 0。它是 words 中以 ["bar","foo"] 顺序排列的连接。
 * 子串 "foobar" 开始位置是 9。它是 words 中以 ["foo","bar"] 顺序排列的连接。
 * 输出顺序无关紧要。返回 [9,0] 也是可以的。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
 * 输出：[]
 * 解释：因为 words.length == 4 并且 words[i].length == 4，所以串联子串的长度必须为 16。
 * s 中没有子串长度为 16 并且等于 words 的任何顺序排列的连接。
 * 所以我们返回一个空数组。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
 * 输出：[6,9,12]
 * 解释：因为 words.length == 3 并且 words[i].length == 3，所以串联子串的长度必须为 9。
 * 子串 "foobarthe" 开始位置是 6。它是 words 中以 ["foo","bar","the"] 顺序排列的连接。
 * 子串 "barthefoo" 开始位置是 9。它是 words 中以 ["bar","the","foo"] 顺序排列的连接。
 * 子串 "thefoobar" 开始位置是 12。它是 words 中以 ["the","foo","bar"] 顺序排列的连接。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * 1 <= words.length <= 5000
 * 1 <= words[i].length <= 30
 * words[i] 和 s 由小写英文字母组成
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	windowSize := len(words)
	res := []int{}

	differ := make(map[string]int)
	for i := 0; i < wordLen && i <= len(s)-windowSize*wordLen; i++ {
		clear(differ)
		for j := 0; j < windowSize; j++ {
			opt := s[i+j*wordLen : i+(j+1)*wordLen]
			differ[opt]++
			differ[words[j]]--
			if differ[opt] == 0 {
				delete(differ, opt)
			}
			if differ[words[j]] == 0 {
				delete(differ, words[j])
			}
		}

		j := i
		for ; j <= len(s)-(windowSize+1)*wordLen; j += wordLen {
			if len(differ) == 0 {
				res = append(res, j)
			}
			out := s[j : j+wordLen]
			in := s[j+windowSize*wordLen : j+(windowSize+1)*wordLen]
			differ[out]--
			differ[in]++
			if differ[out] == 0 {
				delete(differ, out)
			}
			if differ[in] == 0 {
				delete(differ, in)
			}
		}
		if len(differ) == 0 {
			res = append(res, j)
		}
	}

	return res
}

// @lc code=end

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0", args{
			"a",
			[]string{"foo", "bar"},
		}, []int{}},
		{"1", args{
			"wordgoodgoodgoodbestword",
			[]string{"word", "good", "best", "good"},
		}, []int{8}},
		{"2", args{
			"barfoothefoobarman",
			[]string{"foo", "bar"},
		}, []int{0, 9}},
		{"3", args{
			"wordgoodgoodgoodbestword",
			[]string{"word", "good", "best", "word"},
		}, []int{}},
		{"4", args{
			"barfoofoobarthefoobarman",
			[]string{"bar", "foo", "the"},
		}, []int{6, 9, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
