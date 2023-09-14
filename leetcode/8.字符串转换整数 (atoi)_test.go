package leetcode

import (
	"math"
	"testing"
)

/*
请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。

函数 myAtoi(string s) 的算法如下：

读入字符串并丢弃无用的前导空格
检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
如果整数数超过 32 位有符号整数范围 [−2^31,  2^31 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −2^31 的整数应该被固定为 −2^31 ，大于 2^31 − 1 的整数应该被固定为 2^31 − 1 。
返回整数作为最终结果。
注意：

本题中的空白字符只包括空格字符 ' ' 。
除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。


示例 1：

输入：s = "42"
输出：42
解释：加粗的字符串为已经读入的字符，插入符号是当前读取的字符。
第 1 步："42"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："42"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
         ^
第 3 步："42"（读入 "42"）
           ^
解析得到整数 42 。
由于 "42" 在范围 [-2^31, 2^31 - 1] 内，最终结果为 42 。
示例 2：

输入：s = "   -42"
输出：-42
解释：
第 1 步："   -42"（读入前导空格，但忽视掉）
            ^
第 2 步："   -42"（读入 '-' 字符，所以结果应该是负数）
             ^
第 3 步："   -42"（读入 "42"）
               ^
解析得到整数 -42 。
由于 "-42" 在范围 [-2^31, 2^31 - 1] 内，最终结果为 -42 。
示例 3：

输入：s = "4193 with words"
输出：4193
解释：
第 1 步："4193 with words"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："4193 with words"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
         ^
第 3 步："4193 with words"（读入 "4193"；由于下一个字符不是一个数字，所以读入停止）
             ^
解析得到整数 4193 。
由于 "4193" 在范围 [-2^31, 2^31 - 1] 内，最终结果为 4193 。


提示：

0 <= s.length <= 200
s 由英文字母（大写和小写）、数字（0-9）、' '、'+'、'-' 和 '.' 组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/string-to-integer-atoi
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func myAtoi(s string) int {
	return myAtoiDFA(s)
}

func myAtoiOder(s string) int {
	sum := 0
	s = trimPrefix(s)
	sign, s := trimSign(s)
	for i := range s {
		c := s[i]
		if c >= '0' && c <= '9' {
			// if nagtive && ((-sum < math.MinInt32/10) || (-sum == math.MinInt32/10 && c > '8')) {
			// 	return math.MinInt32
			// }
			// if !nagtive && (sum > math.MaxInt32/10 || sum == math.MaxInt32/10 && c > '7') {
			// 	return math.MaxInt32
			// }

			if sum > (math.MaxInt32-int(c-'0'))/10 { // 太妙啦
				if sign == -1 {
					return math.MinInt32
				}
				return math.MaxInt32
			}
			sum = sum*10 + int(c-'0')
		} else {
			break
		}
	}

	if sign == -1 {
		return -sum
	}

	return sum
}

func trimPrefix(s string) string {
	for i := range s {
		if s[i] != ' ' {
			return s[i:]
		}
	}
	return s
}

func trimSign(s string) (int, string) {
	if len(s) == 0 {
		return 1, s
	}

	if s[0] == '-' {
		return -1, s[1:]
	}
	if s[0] == '+' {
		return 1, s[1:]
	}
	return 1, s
}

const (
	end int = iota
	start
	sign
	number
)

const (
	cspace int = iota
	csign
	cnumber
	cother
)

var dfa = [4][4]int{
	start: {
		cspace:  start,
		csign:   sign,
		cnumber: number,
		cother:  end,
	},
	sign: {
		cspace:  end,
		csign:   end,
		cnumber: number,
		cother:  end,
	},
	number: {
		cspace:  end,
		csign:   end,
		cnumber: number,
		cother:  end,
	},
	end: {
		cspace:  end,
		csign:   end,
		cnumber: end,
		cother:  end,
	},
}

// 有限状态机
func myAtoiDFA(s string) int {
	state := start
	signed := 1
	sum := 0
	for i := range s {
		c := s[i]
		switch {
		case c == ' ':
			state = dfa[state][cspace]
		case c == '+':
			state = dfa[state][csign]
		case c == '-':
			state = dfa[state][csign]
			if state == sign {
				signed = -1
			}
		case c >= '0' && c <= '9':
			state = dfa[state][cnumber]
			if sum > (math.MaxInt32-int(c-'0'))/10 {
				if signed == -1 {
					return math.MinInt32
				}
				return math.MaxInt32
			}
			sum = sum*10 + int(c-'0')
		default:
			state = dfa[state][cother]
		}
		if state == end {
			break
		}
	}

	return sum * signed
}

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"1", "123-", 123},
		{"1", "words and 987", 0},
		{"1", "123", 123},
		{"2", "  123", 123},
		{"3", "  +123", 123},
		{"4", "  -123", -123},
		{"5", "  -123  ", -123},
		{"6", "", 0},
		{"6", "-", 0},
		{"6", "+", 0},
		{"7", "4193 with words", 4193},
		{"8", "-2147483648", math.MinInt32},
		{"8", "-2147483649", math.MinInt32},
		{"8", "-2147483647", -2147483647},
		{"9", "2147483647", math.MaxInt32},
		{"8", "2147483648", math.MaxInt32},
		{"8", "2147483646", 2147483646},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
