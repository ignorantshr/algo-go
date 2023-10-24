/*
 * @lc app=leetcode.cn id=514 lang=golang
 *
 * [514] 自由之路
 *
 * https://leetcode.cn/problems/freedom-trail/description/
 *
 * algorithms
 * Hard (51.53%)
 * Likes:    275
 * Dislikes: 0
 * Total Accepted:    28.9K
 * Total Submissions: 56.2K
 * Testcase Example:  '"godding"\n"gd"'
 *
 * 电子游戏“辐射4”中，任务 “通向自由” 要求玩家到达名为 “Freedom Trail Ring” 的金属表盘，并使用表盘拼写特定关键词才能开门。
 *
 * 给定一个字符串 ring ，表示刻在外环上的编码；给定另一个字符串 key ，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。
 *
 * 最初，ring 的第一个字符与 12:00 方向对齐。您需要顺时针或逆时针旋转 ring 以使 key 的一个字符在 12:00
 * 方向对齐，然后按下中心按钮，以此逐个拼写完 key 中的所有字符。
 *
 * 旋转 ring 拼出 key 字符 key[i] 的阶段中：
 *
 *
 * 您可以将 ring 顺时针或逆时针旋转 一个位置 ，计为1步。旋转的最终目的是将字符串 ring 的一个字符与 12:00
 * 方向对齐，并且这个字符必须等于字符 key[i] 。
 * 如果字符 key[i] 已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作 1 步。按完之后，您可以开始拼写 key
 * 的下一个字符（下一阶段）, 直至完成所有拼写。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 *
 *
 * 输入: ring = "godding", key = "gd"
 * 输出: 4
 * 解释:
 * ⁠对于 key 的第一个字符 'g'，已经在正确的位置, 我们只需要1步来拼写这个字符。
 * ⁠对于 key 的第二个字符 'd'，我们需要逆时针旋转 ring "godding" 2步使它变成 "ddinggo"。
 * ⁠当然, 我们还需要1步进行拼写。
 * ⁠因此最终的输出是 4。
 *
 *
 * 示例 2:
 *
 *
 * 输入: ring = "godding", key = "godding"
 * 输出: 13
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= ring.length, key.length <= 100
 * ring 和 key 只包含小写英文字母
 * 保证 字符串 key 一定可以由字符串  ring 旋转拼出
 *
 *
 */
package leetcode

import (
	"math"
	"testing"
)

// @lc code=start
func findRotateSteps(ring string, key string) int {
	size := len(ring)
	ksize := len(key)
	idxes := make(map[byte][]int)

	// 转盘指针位于 ring[i] 时转出 key[:j] 的最小步数
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, ksize)
		for j := 0; j < ksize; j++ {
			dp[i][j] = math.MaxInt
		}
		idxes[ring[i]] = append(idxes[ring[i]], i)
	}

	steps := func(i, j int) int {
		return min(abs(i-j), size-abs(i-j)) + 1
	}

	for j := 0; j < ksize; j++ {
		for _, i := range idxes[key[j]] { // 对每个ring中以key[i]结尾的字符进行遍历取最值
			if j == 0 {
				dp[i][0] = steps(i, 0) // base case
				continue
			}
			for _, k := range idxes[key[j-1]] { // key[i]结尾取决于每种以 key[i-1]结尾的结果加上到达key[i]的步数
				dp[i][j] = min(dp[i][j], dp[k][j-1]+steps(k, i))
			}
		}
	}

	minv := dp[0][ksize-1]
	for i := 0; i < size; i++ {
		minv = min(minv, dp[i][ksize-1])
	}
	return minv
}

// @lc code=end

func Test_findRotateSteps(t *testing.T) {
	type args struct {
		ring string
		key  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"a", "a"}, 1},
		{"1", args{"ab", "a"}, 1},
		{"2", args{"ab", "b"}, 2},
		{"2", args{"ab", "ab"}, 3},
		{"2", args{"abc", "abc"}, 5},
		{"3", args{"godding", "gd"}, 4},
		{"4", args{"godding", "godding"}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRotateSteps(tt.args.ring, tt.args.key); got != tt.want {
				t.Errorf("findRotateSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
