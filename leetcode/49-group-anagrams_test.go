package leetcode

import (
	"strconv"
	"testing"
)

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。



示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func groupAnagrams(strs []string) [][]string {
	return groupAnagramsByIndex(strs)
}

func groupAnagramsByIndex(strs []string) [][]string {
	indexes := make(map[string][]string)
	for _, str := range strs {
		indexes[genIndex(str)] = append(indexes[genIndex(str)], str)
	}

	res := make([][]string, 0, len(indexes))
	for _, v := range indexes {
		res = append(res, v)
	}
	return res
}

func genIndex(str string) string {
	count := make([]int, 26)
	for j := range str {
		count[str[j]-'a']++
	}
	var idx string
	for _, v := range count {
		idx += strconv.Itoa(v)
	}
	return idx
}

func TestGroupAnagrams(t *testing.T) {
	table := []struct {
		name string
		strs []string
		ret  [][]string
	}{
		{
			"1",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			"2",
			[]string{""},
			[][]string{{""}},
		},
		{
			"2",
			[]string{"a"},
			[][]string{{"a"}},
		},
	}

	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			t.Logf("input:%v\noutput:%v\nexpect:%v\n", v.strs, groupAnagrams(v.strs), v.ret)
		})
	}
}
