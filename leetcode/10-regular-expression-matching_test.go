package leetcode

import (
	"testing"
)

/*
10. 正则表达式匹配
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

示例 1：

输入：s = "aa", p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。
示例 2:

输入：s = "aa", p = "a*"
输出：true
解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3：

输入：s = "ab", p = ".*"
输出：true
解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

提示：

1 <= s.length <= 20
1 <= p.length <= 20
s 只包含从 a-z 的小写字母。
p 只包含从 a-z 的小写字母，以及字符 . 和 *。
保证每次出现字符 * 时，前面都匹配到有效的字符

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/regular-expression-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isMatch(s string, p string) bool {
	// f[i][j]: s[:i] 和 p[:j] 的匹配情况
	f := make([][]bool, len(s)+1)
	for i := range f {
		f[i] = make([]bool, len(p)+1)
	}
	f[0][0] = true

	match := func(a, b int) bool {
		if a == 0 {
			return false
		}
		return s[a-1] == p[b-1] || p[b-1] == '.'
	}

	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] != '*' {
				f[i][j] = match(i, j) && f[i-1][j-1]
			} else if match(i, j-1) {
				f[i][j] = f[i][j-2] || f[i-1][j]
			} else {
				f[i][j] = f[i][j-2]
			}
		}
	}

	return f[len(s)][len(p)]
}

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{"aa", "aa*."}, true},
		{"0", args{"aa", "aa*a"}, true},
		{"0", args{"aa", "aaa"}, false},
		{"1", args{"abc", "a*bc"}, true},
		{"1", args{"bc", "a*bc"}, true},
		{"1", args{"aaabc", "a*bc"}, true},
		{"2", args{"abc", "ab*c"}, true},
		{"2", args{"ac", "ab*c"}, true},
		{"2", args{"abbbc", "ab*c"}, true},
		{"3", args{"abc", "abc*"}, true},
		{"3", args{"ab", "abc*"}, true},
		{"3", args{"abcccc", "abc*"}, true},
		{"4", args{"abc", ".bc"}, true},
		{"4", args{"abc", "a.c"}, true},
		{"4", args{"abc", "ab."}, true},
		{"5", args{"abc", "ab.*"}, true},
		{"5", args{"bc", "ab.*"}, false},
		{"6", args{"helloworld", "h.*o.*w.*d"}, true},
		{"6", args{"hello", "world"}, false},
		{"6", args{"hello", "h*llo"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
