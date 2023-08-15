package leetcode

import (
	"math"
	"testing"
)

/*
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0

提示：

-2^31 <= x <= 2^31 - 1

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reverse(x int) int {
	sum := 0
	for x != 0 {
		if sum > math.MaxInt32/10 || sum < math.MinInt32/10 {
			return 0
		}
		remain := x % 10
		sum = sum*10 + remain
		x /= 10
	}

	return int(sum)
}

func TestReverseInteger(t *testing.T) {
	testCases := []struct {
		desc   string
		intput int
		expect int
	}{
		{
			"1",
			1563847412,
			0,
		},
		{
			"1",
			1534236469,
			0,
		},
		{
			"1",
			1230,
			321,
		},
		{
			"1",
			-1230,
			-321,
		},
		{
			"1",
			0,
			0,
		},
		{
			"1",
			2 << 31,
			0,
		},
		{
			"1",
			-(2 << 31) - 1,
			0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if output := reverse(tC.intput); output != tC.expect {
				t.Fail()
			}
		})
	}
}
