/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 *
 * https://leetcode.cn/problems/reverse-string/description/
 *
 * algorithms
 * Easy (79.69%)
 * Likes:    786
 * Dislikes: 0
 * Total Accepted:    769K
 * Total Submissions: 965K
 * Testcase Example:  '["h","e","l","l","o"]'
 *
 * 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
 *
 * 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = ["h","e","l","l","o"]
 * 输出：["o","l","l","e","h"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = ["H","a","n","n","a","h"]
 * 输出：["h","a","n","n","a","H"]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^5
 * s[i] 都是 ASCII 码表中的可打印字符
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// @lc code=end

func Test_reverseString(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{"1", []byte{}, []byte{}},
		{"1", []byte{'a'}, []byte{'a'}},
		{"1", []byte("abc"), []byte("cba")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.input)
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Fatalf("reverseString = %v, want = %v\n", tt.input, tt.want)
			}
		})
	}
}
