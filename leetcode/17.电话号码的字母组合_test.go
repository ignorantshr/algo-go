/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (58.09%)
 * Likes:    2514
 * Dislikes: 0
 * Total Accepted:    713.5K
 * Total Submissions: 1.2M
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：digits = "23"
 * 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：digits = ""
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：digits = "2"
 * 输出：["a","b","c"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= digits.length <= 4
 * digits[i] 是范围 ['2', '9'] 的一个数字。
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start

var phoneByteMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	return letterCombinations2(digits)
}

// 回溯
func letterCombinations2(digits string) []string {
	return backtrack(digits, 0, []byte{})
}

func backtrack(digits string, start int, path []byte) []string {
	if len(path) == len(digits) {
		if len(path) != 0 {
			return []string{string(path)}
		}
		return []string{}
	}
	var res []string
	for i := start; i < len(digits); i++ {
		for _, b := range phoneByteMap[digits[i]] {
			path = append(path, b)
			res = append(res, backtrack(digits, i+1, path)...)
			path = path[:len(path)-1]
		}
	}
	return res
}

// 层级
func letterCombinations1(digits string) []string {
	n := len(digits)
	if n == 0 {
		return []string{}
	}
	elements := make([][]byte, n)
	maxN := 1
	for i := range digits {
		elements[i] = phoneByteMap[digits[i]]
		maxN *= len(elements[i])
	}
	idxed := make([]int, n) // [3,3,4], 0~(3*3*4-1)

	var incr func(level int)
	incr = func(level int) {
		if level == -1 {
			return
		}
		idx := idxed[level] + 1
		if idx == len(elements[level]) {
			idx = 0
			idxed[level] = idx
			incr(level - 1)
		}
		idxed[level] = idx
	}

	genStr := func() string {
		var b []byte
		for i := 0; i < n; i++ {
			b = append(b, elements[i][idxed[i]])
		}
		return string(b)
	}

	var res []string
	for i := 0; i < maxN; i++ {
		res = append(res, genStr())
		incr(n - 1)
	}
	return res
}

// @lc code=end

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{"1", "", []string{}},
		{"1", "2", []string{"a", "b", "c"}},
		{"1", "3", []string{"d", "e", "f"}},
		{"2", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"2", "27", []string{"ap", "aq", "ar", "as", "bp", "bq", "br", "bs", "cp", "cq", "cr", "cs"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
