/*
 * @lc app=leetcode.cn id=68 lang=golang
 *
 * [68] 文本左右对齐
 *
 * https://leetcode.cn/problems/text-justification/description/
 *
 * algorithms
 * Hard (52.85%)
 * Likes:    401
 * Dislikes: 0
 * Total Accepted:    67.4K
 * Total Submissions: 126.4K
 * Testcase Example:  '["This", "is", "an", "example", "of", "text", "justification."]\n16'
 *
 * 给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
 *
 * 你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth
 * 个字符。
 *
 * 要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
 *
 * 文本的最后一行应为左对齐，且单词之间不插入额外的空格。
 *
 * 注意:
 *
 *
 * 单词是指由非空格字符组成的字符序列。
 * 每个单词的长度大于 0，小于等于 maxWidth。
 * 输入单词数组 words 至少包含一个单词。
 *
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: words = ["This", "is", "an", "example", "of", "text", "justification."],
 * maxWidth = 16
 * 输出:
 * [
 * "This    is    an",
 * "example  of text",
 * "justification.  "
 * ]
 *
 *
 * 示例 2:
 *
 *
 * 输入:words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
 * 输出:
 * [
 * "What   must   be",
 * "acknowledgment  ",
 * "shall be        "
 * ]
 * 解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
 * 因为最后一行应为左对齐，而不是左右两端对齐。
 * ⁠    第二行同样为左对齐，这是因为这行只包含一个单词。
 *
 *
 * 示例 3:
 *
 *
 * 输入:words =
 * ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]，maxWidth
 * = 20
 * 输出:
 * [
 * "Science  is  what we",
 * ⁠"understand      well",
 * "enough to explain to",
 * "a  computer.  Art is",
 * "everything  else  we",
 * "do                  "
 * ]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= words.length <= 300
 * 1 <= words[i].length <= 20
 * words[i] 由小写英文字母和符号组成
 * 1 <= maxWidth <= 100
 * words[i].length <= maxWidth
 *
 *
 */
package leetcode

import (
	"slices"
	"strings"
	"testing"
)

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	sizes := make([]int, len(words))
	for i, v := range words {
		sizes[i] = len(v)
	}
	allocation := make([][]string, 0, 1)
	linenum := 0
	idx := 0
	tmpl := 0
	lineSizes := make(map[int]int)
	for idx <= len(words) {
		if len(allocation) == linenum {
			allocation = append(allocation, make([]string, 0, 1))
		}
		if idx != len(words) && sizes[idx]+tmpl <= maxWidth {
			allocation[linenum] = append(allocation[linenum], words[idx])
			tmpl += sizes[idx] + 1 // space
			idx++
		} else {
			lineSizes[linenum] = tmpl - 1 // delete a space
			if idx == len(words) {
				break
			}
			linenum++
			tmpl = 0
		}
	}

	res := []string{}
	for i, line := range allocation {
		ls := lineSizes[i]
		spaceNum := maxWidth - ls
		if len(line) == 1 || i == linenum {
			res = append(res, strings.Join(line, " ")+strings.Repeat(" ", spaceNum))
			continue
		}
		if spaceNum > 0 {
			perNum := spaceNum / (len(line) - 1) // 最后一个不分配空格
			reNum := spaceNum % (len(line) - 1)
			s1 := strings.Join(line[:reNum+1], strings.Repeat(" ", perNum+2))
			s2 := strings.Join(line[reNum+1:], strings.Repeat(" ", perNum+1))
			res = append(res, s1+strings.Repeat(" ", perNum+1)+s2)
			continue

			// // 先均分
			// if perNum > 0 {
			// 	for i := 0; i < len(line)-1; i++ {
			// 		line[i] += strings.Repeat(" ", perNum)
			// 	}
			// }
			// // 分配多余的空格
			// for i := 0; i < reNum; i++ {
			// 	line[i] += " "
			// }
		}
		res = append(res, strings.Join(line, " "))
	}

	return res
}

// @lc code=end

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// {"x", args{[]string{},}, []string{}},
		{"1", args{[]string{"a", "b", "c"}, 5}, []string{"a b c"}},
		{"1.1", args{[]string{"a", "b"}, 1}, []string{"a", "b"}},
		{"1.2", args{[]string{"a", "b"}, 2}, []string{"a ", "b "}},
		{"2", args{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16}, []string{
			"This    is    an",
			"example  of text",
			"justification.  "}},
		{"2.1", args{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16}, []string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		}},
		{"2.2", args{[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20}, []string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  ",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !slices.Equal[[]string](got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
