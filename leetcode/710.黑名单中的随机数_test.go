/*
* @lc app=leetcode.cn id=710 lang=golang
*
* [710] 黑名单中的随机数
*
* https://leetcode.cn/problems/random-pick-with-blacklist/description/
*
  - algorithms
  - Hard (43.67%)
  - Likes:    230
  - Dislikes: 0
  - Total Accepted:    30.3K
  - Total Submissions: 69.3K
  - Testcase Example:  '["Solution","pick","pick","pick","pick","pick","pick","pick"]\n' +
    '[[7,[2,3,5]],[],[],[],[],[],[],[]]'

*
* 给定一个整数 n 和一个 无重复 黑名单整数数组 blacklist 。设计一种算法，从 [0, n - 1] 范围内的任意整数中选取一个 未加入
* 黑名单 blacklist 的整数。任何在上述范围内且不在黑名单 blacklist 中的整数都应该有 同等的可能性 被返回。
*
* 优化你的算法，使它最小化调用语言 内置 随机函数的次数。
*
* 实现 Solution 类:
*
*
* Solution(int n, int[] blacklist) 初始化整数 n 和被加入黑名单 blacklist 的整数
* int pick() 返回一个范围为 [0, n - 1] 且不在黑名单 blacklist 中的随机整数
*
*
*
*
* 示例 1：
*
*
* 输入
* ["Solution", "pick", "pick", "pick", "pick", "pick", "pick", "pick"]
* [[7, [2, 3, 5]], [], [], [], [], [], [], []]
* 输出
* [null, 0, 4, 1, 6, 1, 0, 4]
*
* 解释
* Solution solution = new Solution(7, [2, 3, 5]);
* solution.pick(); // 返回0，任何[0,1,4,6]的整数都可以。注意，对于每一个pick的调用，
* ⁠                // 0、1、4和6的返回概率必须相等(即概率为1/4)。
* solution.pick(); // 返回 4
* solution.pick(); // 返回 1
* solution.pick(); // 返回 6
* solution.pick(); // 返回 1
* solution.pick(); // 返回 0
* solution.pick(); // 返回 4
*
*
*
*
* 提示:
*
*
* 1 <= n <= 10^9
* 0 <= blacklist.length <= min(10^5, n - 1)
* 0 <= blacklist[i] < n
* blacklist 中所有值都 不同
* pick 最多被调用 2 * 10^4 次
*
*
*/
package leetcode

import (
	"math/rand"
	"testing"
)

// @lc code=start
type Solution struct {
	size        int
	replacelist map[int]int
}

func Constructor710(n int, blacklist []int) Solution {
	tail := make(map[int]int, len(blacklist))
	s := Solution{n - len(blacklist), make(map[int]int, len(blacklist))}
	for _, v := range blacklist {
		tail[v] = -1
	}

	last := n - 1
	for i := 0; i < len(blacklist) && last >= 0; {
		if blacklist[i] >= s.size { // 如果已经在 [size, n) 中也不需要映射了
			i++
			continue
		}
		if tail[last] != -1 { // 找一个不在黑名单中的位置
			s.replacelist[blacklist[i]] = last // 用 last 替换黑名单中的值
			i++
		}
		last--
	}

	return s
}

func (this *Solution) Pick() int {
	res := rand.Intn(this.size)
	if v, ok := this.replacelist[res]; ok {
		return v
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
// @lc code=end

func TestConstructor710(t *testing.T) {
	n, list := 4, []int{2, 1}
	// n, list := 4, []int{1, 3}
	blackmap := make(map[int]struct{})
	for _, v := range list {
		blackmap[v] = struct{}{}
	}

	s := Constructor710(n, list)
	for i := 0; i < 10; i++ {
		v := s.Pick()
		if _, ok := blackmap[v]; ok {
			t.Errorf("Pick() = %v", v)
		}
	}
}
