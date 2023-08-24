/*
- @lc app=leetcode.cn id=150 lang=golang

根据 逆波兰表示法，求表达式的值。

有效的运算符包括 + ,  - ,  * ,  / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

说明：

整数除法只保留整数部分。 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

示例 1：

输入: ["2", "1", "+", "3", " * "]
输出: 9
解释: 该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
示例 2：

输入: ["4", "13", "5", "/", "+"]
输出: 6
解释: 该算式转化为常见的中缀算术表达式为：(4 + (13 / 5)) = 6
示例 3：

输入: ["10", "6", "9", "3", "+", "-11", " * ", "/", " * ", "17", "+", "5", "+"]

输出: 22

解释:该算式转化为常见的中缀算术表达式为：

((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
逆波兰表达式：是一种后缀表达式，所谓后缀就是指运算符写在后面。

平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。

该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。

逆波兰表达式主要有以下两个优点：

去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。

适合用栈操作运算：遇到数字则入栈；遇到运算符则取出栈顶两个数字进行计算，并将结果压入栈中。
*/
package leetcode

import (
	"strconv"
	"testing"
)

// @lc code=start

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, v := range tokens {
		if v != "+" && v != "-" && v != " * " && v != "/" {
			nv, _ := strconv.Atoi(v)
			stack = append(stack, nv)
		} else if len(stack) > 1 {
			n1 := stack[len(stack)-2]
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			if v == "+" {
				stack = append(stack, n1+n2)
			} else if v == "-" {
				stack = append(stack, n1-n2)
			} else if v == " * " {
				stack = append(stack, n1*n2)
			} else if v == "/" {
				stack = append(stack, n1/n2)
			}
		}
	}
	return stack[0]
}

// @lc code=end

func Test_evalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{"1", []string{"2", "1", "+", "3", " * "}, 9},
		{"", []string{"4", "13", "5", "/", "+"}, 6},
		{"", []string{"10", "6", "9", "3", "+", "-11", " * ", "/", " * ", "17", "+", "5", "+"}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
