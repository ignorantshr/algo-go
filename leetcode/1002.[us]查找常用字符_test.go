/*
- @lc app=leetcode.cn id=1002 lang=golang

给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（ 包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答案。

示例 1：

输入：words = ["bella","label","roller"] 输出：["e","l","l"]

示例 2：

输入：words = ["cool","lock","cook"] 输出：["c","o"]

提示：

1 <= words.length <= 100 1 <= words[i].length <= 100 words[i] 由小写英文字母组成
*/
package leetcode

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// @lc code=start
func commonChars(words []string) []string {
	counter := [26]int{}
	for i := 0; i < 26; i++ {
		counter[i] = math.MaxInt
	}
	for _, str := range words {
		singleCounter := [26]int{}
		for _, r := range str {
			singleCounter[r-'a']++
		}
		for j, v := range singleCounter {
			if counter[j] > v {
				counter[j] = v
			}
		}
	}

	if len(words) == 0 {
		return []string{}
	}
	var res []string
	for i, v := range counter {
		for j := 0; j < v; j++ {
			res = append(res, fmt.Sprintf("%c", i+'a'))
		}
	}
	return res
}

// @lc code=end

func Test_commonChars(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{"0", []string{}, []string{}},
		{"1", []string{"bella", "label", "roller"}, []string{"e", "l", "l"}},
		{"1", []string{"cool", "lock", "cook"}, []string{"c", "o"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonChars(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
