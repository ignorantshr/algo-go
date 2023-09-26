package leetcode

import (
	"reflect"
	"testing"
)

var phoneByteMapRV = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations_RV(digits string) []string {
	res := make([]string, 0)
	path := make([]byte, 0)
	var backtrack func(count int)
	backtrack = func(count int) {
		if count == len(digits) {
			if len(path) == 0 {
				return
			}
			res = append(res, string(path))
			return
		}

		for i := 0; i < len(phoneByteMapRV[digits[count]]); i++ {
			path = append(path, phoneByteMapRV[digits[count]][i])
			backtrack(count + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}

func Test_letterCombinations_RV(t *testing.T) {
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
			if got := letterCombinations_RV(tt.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations_RV() = %v, want %v", got, tt.want)
			}
		})
	}
}
