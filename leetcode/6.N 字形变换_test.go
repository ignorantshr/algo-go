package leetcode

import (
	"testing"
)

/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);

示例 1：

输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I
示例 3：

rn=5
1    9
2   8 10
3  7   11
4 6     12
5        13

m=5,n=5,4,...,1
2n-3, 2n-3, ...				7
2n-3, 2(m)-3 - (2n-3)-1 	5, 7-5-1=1
2n-3, 2(m)-3 - (2n-3)-1		3, 7-3-1=3
2n-3=-1时，取 2(m)-3

0，0
1，1
2，3
3，5
4，7

输入：s = "A", numRows = 1
输出："A"

提示：

1 <= s.length <= 1000
s 由英文字母（小写和大写）、',' 和 '.' 组成
1 <= numRows <= 1000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/zigzag-conversion
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func convert(s string, numRows int) string {
	return convertOddEvenSpan(s, numRows)
}

func convertOddEvenSpan(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var res []byte
	tspan := 2 * (numRows - 1)
	for offset := 0; offset < numRows; offset++ {
		tmp := 2 * offset
		for i := offset; i < len(s); i += tspan {
			res = append(res, s[i]) // 周期内第一个字符
			nspan := tspan - tmp
			if nspan == 0 {
				nspan = tspan
			}
			if i+nspan < len(s) && nspan < tspan {
				res = append(res, s[i+nspan]) // 周期内第二个字符
			}
		}
	}
	return string(res)
}

func convertRemain(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make(map[int][]byte, numRows)
	n := 0
	incr := true
	for i := range s {
		rows[n] = append(rows[n], s[i])
		if incr {
			n++
		} else {
			n--
		}

		if n == 0 || n >= numRows-1 {
			incr = !incr
		}
	}
	var res string
	for i := 0; i < numRows; i++ {
		res += string(rows[i])
	}
	return res
}

func TestZigZagConvertion(t *testing.T) {
	table := []struct {
		name   string
		input  string
		rows   int
		expect string
	}{
		{
			"1",
			"PAYPALISHIRING",
			3,
			"PAHNAPLSIIGYIR",
		},
		{
			"2",
			"PAYPALISHIRING",
			4,
			"PINALSIGYAHRPI",
		},
		{
			"3",
			"a",
			1,
			"a",
		},
		{
			"4",
			"abc",
			1,
			"abc",
		},
	}

	for _, item := range table {
		t.Run(item.name, func(t *testing.T) {
			output := convert(item.input, item.rows)
			if output != item.expect {
				t.Logf("\noutput: %s, expect: %s", output, item.expect)
				t.Fail()
			}
		})
	}
}
