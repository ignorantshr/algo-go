/*
- @lc app=leetcode.cn id=93 lang=golang

给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。

示例 1：

输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
示例 2：

输入：s = "0000"
输出：["0.0.0.0"]
示例 3：

输入：s = "1111"
输出：["1.1.1.1"]
示例 4：

输入：s = "010010"
输出：["0.10.0.10","0.100.1.0"]
示例 5：

输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
提示：

0 <= s.length <= 3000
s 仅由数字组成
*/
package leetcode

import (
	"strconv"
	"strings"
	"testing"
)

// @lc code=start
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	path := make([]string, 0)

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == len(s) {
			if len(path) == 4 {
				res = append(res, strings.Join(path, "."))
			}
			return
		}

		isSegment := func(start, end int) bool {
			tmp := s[start : end+1]
			num, _ := strconv.Atoi(tmp)
			if end != start && tmp[0] == '0' {
				return false
			}
			return num >= 0 && num <= 255
		}

		for i := idx; i < len(s); i++ {
			if !isSegment(idx, i) || len(path) == 4 {
				break
			}

			path = append(path, s[idx:i+1])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}

// @lc code=end

func Test_restoreIpAddresses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"0", "", []string{}},
		{"0", "1", []string{}},
		{"0", "1234567891", []string{}},
		{"1", "0000", []string{"0.0.0.0"}},
		{"1", "1111", []string{"1.1.1.1"}},
		{"2", "010010", []string{"0.10.0.10", "0.100.1.0"}},
		{"2", "25525511135", []string{"255.255.11.135", "255.255.111.35"}},
		{"2", "101023", []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.s); !equalSet(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
