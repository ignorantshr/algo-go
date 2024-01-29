/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] 扰乱字符串
 *
 * https://leetcode.cn/problems/scramble-string/description/
 *
 * algorithms
 * Hard (47.20%)
 * Likes:    557
 * Dislikes: 0
 * Total Accepted:    61.6K
 * Total Submissions: 130.6K
 * Testcase Example:  '"great"\n"rgeat"'
 *
 * 使用下面描述的算法可以扰乱字符串 s 得到字符串 t ：
 *
 * 如果字符串的长度为 1 ，算法停止
 * 如果字符串的长度 > 1 ，执行下述步骤：
 *
 * 在一个随机下标处将字符串分割成两个非空的子字符串。即，如果已知字符串 s ，则可以将其分成两个子字符串 x 和 y ，且满足 s = x + y
 * 。
 * 随机 决定是要「交换两个子字符串」还是要「保持这两个子字符串的顺序不变」。即，在执行这一步骤之后，s 可能是 s = x + y 或者 s = y +
 * x 。
 * 在 x 和 y 这两个子字符串上继续从步骤 1 开始递归执行此算法。
 *
 *
 *
 *
 * 给你两个 长度相等 的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。如果是，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s1 = "great", s2 = "rgeat"
 * 输出：true
 * 解释：s1 上可能发生的一种情形是：
 * "great" --> "gr/eat" // 在一个随机下标处分割得到两个子字符串
 * "gr/eat" --> "gr/eat" // 随机决定：「保持这两个子字符串的顺序不变」
 * "gr/eat" --> "g/r / e/at" // 在子字符串上递归执行此算法。两个子字符串分别在随机下标处进行一轮分割
 * "g/r / e/at" --> "r/g / e/at" // 随机决定：第一组「交换两个子字符串」，第二组「保持这两个子字符串的顺序不变」
 * "r/g / e/at" --> "r/g / e/ a/t" // 继续递归执行此算法，将 "at" 分割得到 "a/t"
 * "r/g / e/ a/t" --> "r/g / e/ a/t" // 随机决定：「保持这两个子字符串的顺序不变」
 * 算法终止，结果字符串和 s2 相同，都是 "rgeat"
 * 这是一种能够扰乱 s1 得到 s2 的情形，可以认为 s2 是 s1 的扰乱字符串，返回 true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s1 = "abcde", s2 = "caebd"
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：s1 = "a", s2 = "a"
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * s1.length == s2.length
 * 1
 * s1 和 s2 由小写英文字母组成
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func isScramble(s1 string, s2 string) bool {
	// 链接：https://leetcode.cn/problems/scramble-string/solutions/51990/miao-dong-de-qu-jian-xing-dpsi-lu-by-sha-yu-la-jia/
	if len(s1) != len(s2) {
		return false
	}

	size := len(s1)
	dp := make([][][]bool, size) // dp[i][j][k] s1[i:i+k] == s2[j:j+k]
	for i := 0; i < size; i++ {
		dp[i] = make([][]bool, size)
		for j := 0; j < size; j++ {
			dp[i][j] = make([]bool, size+1)
		}
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			dp[i][j][1] = s1[i] == s2[j]
		}
	}

	// 枚举区间长度
	for l := 2; l <= size; l++ {
		// 枚举 s1 的起点位置
		for i := 0; i <= size-l; i++ {
			// 枚举 s2 的起点位置
			for j := 0; j <= size-l; j++ {
				// 枚举划分位置
				for k := 1; k <= l-1; k++ {
					// 第一种情况：S1 -> T1, S2 -> T2
					if dp[i][j][k] && dp[i+k][j+k][l-k] {
						dp[i][j][l] = true
						break
					}

					// 第二种情况：S1 -> T2, S2 -> T1
					// S1 起点 i，T2 起点 j + 前面那段长度 len-k ，S2 起点 i + 前面长度k
					if dp[i][j+l-k][k] && dp[i+k][j][l-k] {
						dp[i][j][l] = true
						break
					}
				}
			}
		}
	}

	return dp[0][0][size]
}

// @lc code=end

func Test_isScramble(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1.1", args{"g", "a"}, false},
		{"1.2", args{"g", "g"}, true},

		{"2.1", args{"ga", "ga"}, true},
		{"2.2", args{"ga", "ag"}, true},
		{"2.3", args{"ga", "ab"}, false},

		{"3.1", args{"gva", "gav"}, true},

		{"true", args{"great", "great"}, true},
		{"true.1", args{"great", "rgeat"}, true},
		{"false", args{"abcde", "caebd"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isScramble(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isScramble() = %v, want %v", got, tt.want)
			}
		})
	}
}
