package leetcode

import (
	"strings"
	"testing"
)

/*
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给你一个整数，将其转为罗马数字。



示例 1:

输入: num = 3
输出: "III"
示例 2:

输入: num = 4
输出: "IV"
示例 3:

输入: num = 9
输出: "IX"
示例 4:

输入: num = 58
输出: "LVIII"
解释: L = 50, V = 5, III = 3.
示例 5:

输入: num = 1994
输出: "MCMXCIV"
解释: M = 1000, CM = 900, XC = 90, IV = 4.


提示：

1 <= num <= 3999

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/integer-to-roman
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type RS int

func (r RS) String() string {
	switch r {
	case I:
		return "I"
	case V:
		return "V"
	case X:
		return "X"
	case L:
		return "L"
	case C:
		return "C"
	case D:
		return "D"
	case M:
		return "M"
	}
	return ""
}

const (
	I RS = 1
	V RS = 5
	X RS = 10
	L RS = 50
	C RS = 100
	D RS = 500
	M RS = 1000
)

// [9|4][9|4][9|4]
var rss = [7]RS{M, D, C, L, X, V, I}

var romas = map[int]map[int]RS{
	1000: {
		1000: M,
		4:    M,
		9:    M,
	},
	100: {
		100: C,
		4:   D,
		9:   M,
	},
	10: {
		10: X,
		4:  L,
		9:  C,
	},
	1: {
		1: I,
		4: V,
		9: X,
	},
}

func intToRoman(num int) string {
	return intToRoman4(num)
}

func intToRoman4(num int) string {
	rs := []struct {
		value int
		str   string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	res := make([]byte, 0)
	for _, v := range rs {
		for num >= v.value {
			num -= v.value
			res = append(res, v.str...)
		}
		if num == 0 {
			break
		}
	}
	return string(res)
}

func intToRoman3(num int) string {
	res := ""
	divisor := 1000

	for divisor > 0 && num > 0 {
		q, r := num/divisor, num%divisor
		if q == 9 || q == 4 {
			res += romas[divisor][divisor].String() + romas[divisor][q].String()
			num = r
			divisor /= 10
			continue
		}
		if q >= 5 {
			res += romas[divisor][4].String()
			q -= 5
		}
		for j := 0; j < q; j++ {
			res += romas[divisor][divisor].String()
		}
		divisor /= 10
		num = r
	}
	return res
}

func intToRoman2(num int) string {
	res := ""
	for i := 0; i < 7; i++ {
		rs := rss[i]
		q, r := remain(num, rs)
		if rs == L || rs == D || rs == V {
			if r/int(rss[i+1]) == 4 && q <= 1 {
				if q == 0 {
					res += (rss[i+1]).String() + rs.String()
				} else {
					res += (rss[i+1]).String() + (rss[i-1]).String()
				}
				num = r % int(rss[i+1])
				continue
			}
		}
		for j := 0; j < q; j++ {
			res += rs.String()
		}
		if r == 0 {
			break
		}
		num = r
	}
	return res
}

func intToRomanStr(num int) string {
	res := ""
	for i := 0; i < 7; i++ {
		q, r := remain(num, rss[i])
		for j := 0; j < q; j++ {
			res += rss[i].String()
		}
		if r == 0 {
			break
		}
		num = r
	}
	return format(res)
}

func remain(num int, r RS) (int, int) {
	return num / int(r), num % int(r)
}

func format(s string) string {
	s = strings.ReplaceAll(s, "DCCCC", "CM")
	s = strings.ReplaceAll(s, "CCCC", "CD")
	s = strings.ReplaceAll(s, "LXXXX", "XC")
	s = strings.ReplaceAll(s, "XXXX", "XL")
	s = strings.ReplaceAll(s, "VIIII", "IX")
	s = strings.ReplaceAll(s, "IIII", "IV")
	return s
}

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{"1", 1000, "M"},
		{"1", 2000, "MM"},
		{"1", 3999, "MMMCMXCIX"},
		{"2", 4, "IV"},
		{"2", 9, "IX"},
		{"2", 40, "XL"},
		{"2", 90, "XC"},
		{"2", 400, "CD"},
		{"2", 900, "CM"},
		{"2", 1994, "MCMXCIV"},
		{"3", 58, "LVIII"},
		{"3", 68, "LXVIII"},
		{"3", 37, "XXXVII"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
