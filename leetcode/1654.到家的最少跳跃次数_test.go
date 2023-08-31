/*
 * @lc app=leetcode.cn id=1654 lang=golang
 *
 * [1654] 到家的最少跳跃次数
 *
 * https://leetcode.cn/problems/minimum-jumps-to-reach-home/description/
 *
 * algorithms
 * Medium (31.15%)
 * Likes:    177
 * Dislikes: 0
 * Total Accepted:    18.3K
 * Total Submissions: 51.9K
 * Testcase Example:  '[14,4,18,1,15]\n3\n15\n9'
 *
 * 有一只跳蚤的家在数轴上的位置 x 处。请你帮助它从位置 0 出发，到达它的家。
 *
 * 跳蚤跳跃的规则如下：
 *
 *
 * 它可以 往前 跳恰好 a 个位置（即往右跳）。
 * 它可以 往后 跳恰好 b 个位置（即往左跳）。
 * 它不能 连续 往后跳 2 次。
 * 它不能跳到任何 forbidden 数组中的位置。
 *
 *
 * 跳蚤可以往前跳 超过 它的家的位置，但是它 不能跳到负整数 的位置。
 *
 * 给你一个整数数组 forbidden ，其中 forbidden[i] 是跳蚤不能跳到的位置，同时给你整数 a， b 和 x
 * ，请你返回跳蚤到家的最少跳跃次数。如果没有恰好到达 x 的可行方案，请你返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：forbidden = [14,4,18,1,15], a = 3, b = 15, x = 9
 * 输出：3
 * 解释：往前跳 3 次（0 -> 3 -> 6 -> 9），跳蚤就到家了。
 *
 *
 * 示例 2：
 *
 *
 * 输入：forbidden = [8,3,16,6,12,20], a = 15, b = 13, x = 11
 * 输出：-1
 *
 *
 * 示例 3：
 *
 *
 * 输入：forbidden = [1,6,2,14,5,17,4], a = 16, b = 9, x = 7
 * 输出：2
 * 解释：往前跳一次（0 -> 16），然后往回跳一次（16 -> 7），跳蚤就到家了。
 *
 *
 *
 *
 * 提示：
 * 1 <= forbidden.length <= 1000
 * 1 <= a, b, forbidden[i] <= 2000
 * 0 <= x <= 2000
 * forbidden 中所有位置互不相同。
 * 位置 x 不在 forbidden 中。
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func minimumJumps(forbidden []int, a int, b int, x int) int {
	return minimumJumpsBfs(forbidden, a, b, x)
	// return minimumJumpsDfs(forbidden, a, b, x)
}

func minimumJumpsBfs(forbidden []int, a int, b int, x int) int {
	bans := make(map[int]struct{}, len(forbidden))
	for _, v := range forbidden {
		bans[v] = struct{}{}
	}

	queue := make([][2]int, 0) // {path, 0/1}  0: foward; 1:backward
	count := -1

	queue = append(queue, [2]int{0, 0})
	bans[0] = struct{}{}

	for len(queue) != 0 {
		count++
		size := len(queue)
		for i := 0; i < size; i++ {
			ele := queue[i]
			if ele[0] == x {
				return count
			}

			if _, has := bans[ele[0]+a]; !has && ele[0]+a < 6000 {
				bans[ele[0]+a] = struct{}{}
				queue = append(queue, [2]int{ele[0] + a, 0})
			}
			if _, has := bans[ele[0]-b]; !has && ele[1] == 0 && ele[0]-b > 0 {
				// bans[ele[0]-b] = struct{}{} // 不可
				queue = append(queue, [2]int{ele[0] - b, 1})
			}
		}
		queue = queue[size:]
	}

	return -1
}

func minimumJumpsDfs(forbidden []int, a int, b int, x int) int {
	bans := make(map[int]struct{}, len(forbidden))
	for _, v := range forbidden {
		bans[v] = struct{}{}
	}
	minium := -1
	var skip func(p, backCount, count int) int
	skip = func(p, backCount, count int) int {
		if minium != -1 && count >= minium {
			return -1
		}
		if p == x {
			if minium == -1 || minium > count {
				minium = count
			}
			return count
		}
		if a >= b && p-(1-backCount)*b > x { // 回不来了
			return -1
		}

		c1 := -1
		if _, has := bans[p+a]; !has && p+a < 6000 {
			bans[p+a] = struct{}{}
			c1 = skip(p+a, 0, count+1)
		}

		c2 := -1
		if _, has := bans[p-b]; !has && p > b && backCount < 1 {
			// bans[p-b] = struct{}{} // 不可
			c2 = skip(p-b, 1, count+1)
		}

		if c1 == -1 {
			return c2
		}
		if c2 == -1 {
			return c1
		}
		if c1 < c2 {
			return c1
		}
		return c2
	}

	return skip(0, 0, 0)
}

// @lc code=end

func Test_minimumJumps(t *testing.T) {
	type args struct {
		a         int
		b         int
		x         int
		forbidden []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{
			14, 5, 90,
			[]int{3},
		}, 20},
		{"1", args{
			806, 1994, 326,
			[]int{1906, 1988, 1693, 483, 900, 1173, 805, 1593, 1208, 1084, 300, 614, 1325, 783, 1104, 1450, 311, 1506, 1388, 1567, 1497, 47, 102, 338, 1937, 888, 111, 195, 1041, 1570, 686, 1707, 1521, 1566, 74, 1264, 667, 1486, 960, 389, 442, 329, 1577, 1557, 1494, 1382, 1688, 779, 484, 410, 227, 1025, 1417, 1475, 1042, 1903, 1920, 1712, 870, 1813, 1137, 1732, 18, 1065, 1653, 1289, 1636, 147, 1833, 1168, 1087, 1408, 881, 1129, 71, 924, 1718, 1458, 371, 597, 1790, 889, 414, 784, 1883, 6, 1650, 1549, 552, 1233, 1467, 1514, 1568, 211, 1301, 772, 377, 1751, 1699, 1701, 1214, 1874, 324, 1991, 1006, 1413, 41, 289, 1274, 802, 1892, 1908, 1960, 1635, 69, 423, 1795, 96, 1024, 1596, 1044, 1513, 1390, 711, 1806, 1298, 968, 1160, 1232, 1315, 1646, 1178, 169, 1295, 466, 44, 10, 1250, 1283, 927, 49, 267, 1773, 342, 1828, 1949, 1291, 244, 707, 408, 798, 938, 1542, 690, 639, 1148, 1081, 431, 752, 120, 1125, 339, 480, 247, 733, 266, 596, 987, 777, 214, 1005, 1687, 160, 785, 1010, 1282, 1135, 922, 671, 1221, 250, 1982, 398, 1959, 179, 325, 1313, 577, 1053, 1436, 185, 1014, 1851, 1685, 1143, 1510, 1972, 830, 681, 390, 972, 1003, 844, 229, 1246, 1257, 668, 1765, 619, 276, 1355, 1544, 1842, 1340, 1375, 1944, 790, 606, 345, 1487, 796, 1985, 1673, 1503, 180, 1642, 498, 1805, 201, 104, 1658, 1633, 1507, 1142, 541, 865, 1193, 485, 216, 1849, 359, 1422, 391, 856, 1864, 470, 1888, 1698, 760, 1778, 572, 1057, 48, 189, 1086, 1704, 1258, 192, 825, 585, 152, 1865, 1645, 807, 225, 402, 1198, 1476, 600, 1914, 975, 1378, 1190, 24, 1550, 723, 696, 1131, 1831, 1880, 1029, 713, 486, 126, 876, 1270, 1891, 544, 61, 1356, 1676, 1239, 36, 1177, 620, 1723, 1651, 1136, 141, 1889, 1123, 624, 1519, 725, 241, 1253, 1119, 269, 763, 1120, 1620, 642, 1713, 966, 1204, 558, 1344, 550, 316, 412, 886, 1309, 1648, 599, 1893, 265, 258, 1561, 477, 1967, 66, 1296, 75, 1628, 715, 826, 1942, 1966, 1407, 159, 646, 1438, 1730, 768, 411, 287, 499, 467, 46, 302, 661, 526, 848, 1327, 1097, 166, 413, 1578, 574, 1304, 925, 504, 914, 978, 1352, 1103, 1859, 1167, 1318, 1454, 1990, 739, 1252, 132, 529, 1622, 422, 1744, 1819, 425, 945, 1767, 1791, 976, 1226, 1092, 305, 479, 174, 626, 1063, 662, 1948, 1978, 524, 512, 1255, 651, 1678, 1059},
		}, -1},
		{"1", args{
			29, 98, 80,
			[]int{162, 118, 178, 152, 167, 100, 40, 74, 199, 186, 26, 73, 200, 127, 30, 124, 193, 84, 184, 36, 103, 149, 153, 9, 54, 154, 133, 95, 45, 198, 79, 157, 64, 122, 59, 71, 48, 177, 82, 35, 14, 176, 16, 108, 111, 6, 168, 31, 134, 164, 136, 72, 98},
		}, 121},
		{"1", args{
			3, 15, 9,
			[]int{14, 4, 18, 1, 15},
		}, 3},
		{"1", args{
			15, 13, 11,
			[]int{8, 3, 16, 6, 12, 20},
		}, -1},
		{"1", args{
			16, 9, 7,
			[]int{1, 6, 2, 14, 5, 17, 4},
		}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumJumps(tt.args.forbidden, tt.args.a, tt.args.b, tt.args.x); got != tt.want {
				t.Errorf("minimumJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}
