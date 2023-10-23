/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] 地下城游戏
 *
 * https://leetcode.cn/problems/dungeon-game/description/
 *
 * algorithms
 * Hard (48.71%)
 * Likes:    788
 * Dislikes: 0
 * Total Accepted:    67.8K
 * Total Submissions: 139.3K
 * Testcase Example:  '[[-2,-3,3],[-5,-10,1],[10,30,-5]]'
 *
 * table.dungeon, .dungeon th, .dungeon td {
 * ⁠ border:3px solid black;
 * }
 *
 * ⁠.dungeon th, .dungeon td {
 * ⁠   text-align: center;
 * ⁠   height: 70px;
 * ⁠   width: 70px;
 * }
 *
 * 恶魔们抓住了公主并将她关在了地下城 dungeon 的 右下角 。地下城是由 m x n 个房间组成的二维网格。我们英勇的骑士最初被安置在 左上角
 * 的房间里，他必须穿过地下城并通过对抗恶魔来拯救公主。
 *
 * 骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。
 *
 * 有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；其他房间要么是空的（房间里的值为
 * 0），要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。
 *
 * 为了尽快解救公主，骑士决定每次只 向右 或 向下 移动一步。
 *
 * 返回确保骑士能够拯救到公主所需的最低初始健康点数。
 *
 * 注意：任何房间都可能对骑士的健康点数造成威胁，也可能增加骑士的健康点数，包括骑士进入的左上角房间以及公主被监禁的右下角房间。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：dungeon = [[-2,-3,3],[-5,-10,1],[10,30,-5]]
 * 输出：7
 * 解释：如果骑士遵循最佳路径：右 -> 右 -> 下 -> 下 ，则骑士的初始健康点数至少为 7 。
 *
 * 示例 2：
 *
 *
 * 输入：dungeon = [[0]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == dungeon.length
 * n == dungeon[i].length
 * 1 <= m, n <= 200
 * -1000 <= dungeon[i][j] <= 1000
 *
 *
 */
package leetcode

import (
	"math"
	"testing"
)

// @lc code=start
func calculateMinimumHP(dungeon [][]int) int {
	rowSize := len(dungeon)
	colSize := len(dungeon[0])
	// dp[i][j]:
	// ❌ 到达 grid[i][j] 时的最大生命值
	// ❌ 从 grid[0][0] 到达右下角所需的最小生命值
	// ✅ 从 grid[i][j] 到达右下角所需的最小生命值
	dp := make([][]int, rowSize+1)
	for i := 0; i <= rowSize; i++ {
		dp[i] = make([]int, colSize+1)
		dp[i][colSize] = math.MaxInt
	}

	for j := 0; j <= colSize; j++ {
		dp[rowSize][j] = math.MaxInt
	}

	dp[rowSize-1][colSize] = 1 // 在边界处预设一个值，就不用在循环里单独判断 dp[rowSize-1][colSize-1] 的情况了

	for i := rowSize - 1; i >= 0; i-- {
		for j := colSize - 1; j >= 0; j-- {
			res := min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
			if res > 0 {
				dp[i][j] = res
			} else {
				dp[i][j] = 1
			}
		}
	}

	return dp[0][0]
}

// @lc code=end

func Test_calculateMinimumHP(t *testing.T) {
	tests := []struct {
		name    string
		dungeon [][]int
		want    int
	}{
		{"1", [][]int{{0}}, 1},
		{"1", [][]int{{1}}, 1},
		{"2", [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}, 7},
		{"3", [][]int{{0, -10, 20}, {0, 0, 100}, {-1, -1, 0}}, 1},
		{"4", [][]int{{0, 0, 0}, {0, -100, 10}, {0, 0, -5}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMinimumHP(tt.dungeon); got != tt.want {
				t.Errorf("calculateMinimumHP() = %v, want %v", got, tt.want)
			}
		})
	}
}
