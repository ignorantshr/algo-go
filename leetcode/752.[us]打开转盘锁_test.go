/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 *
 * https://leetcode.cn/problems/open-the-lock/description/
 *
 * algorithms
 * Medium (52.71%)
 * Likes:    633
 * Dislikes: 0
 * Total Accepted:    123.4K
 * Total Submissions: 234.3K
 * Testcase Example:  '["0201","0101","0102","1212","2002"]\n"0202"'
 *
 * 你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8',
 * '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
 *
 * 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
 *
 * 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
 *
 * 字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
 * 输出：6
 * 解释：
 * 可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
 * 注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
 * 因为当拨动到 "0102" 时这个锁就会被锁定。
 *
 *
 * 示例 2:
 *
 *
 * 输入: deadends = ["8888"], target = "0009"
 * 输出：1
 * 解释：把最后一位反向旋转一次即可 "0000" -> "0009"。
 *
 *
 * 示例 3:
 *
 *
 * 输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"],
 * target = "8888"
 * 输出：-1
 * 解释：无法旋转到目标数字且不被锁定。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= deadends.length <= 500
 * deadends[i].length == 4
 * target.length == 4
 * target 不在 deadends 之中
 * target 和 deadends[i] 仅由若干位数字组成
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func openLock(deadends []string, target string) int {
	return openLockTwoWayBFS(deadends, target)
	// return openLockBfs(deadends, target)
}

// 双向bfs
// 局限性：必须同时知道终点和起点
func openLockTwoWayBFS(deadends []string, target string) int {
	var step int
	used := make(map[string]struct{})
	for _, v := range deadends {
		used[v] = struct{}{}
	}
	q1 := make(map[string]bool)
	q2 := make(map[string]bool)
	q1["0000"] = true
	q2[target] = true

	for len(q1) != 0 && len(q2) != 0 {
		tmp := make(map[string]bool)

		for cur := range q1 {
			if q2[cur] {
				return step
			}
			if _, has := used[cur]; has {
				continue
			}
			used[cur] = struct{}{}

			for i := 0; i < 4; i++ {
				up := plusOnce(i, cur)
				if _, has := used[up]; !has {
					tmp[up] = true // 存储本轮扩散结果
				}
				down := minusOnce(i, cur)
				if _, has := used[down]; !has {
					tmp[down] = true
				}
			}
		}
		step++
		// 交换扩散顺序，存储扩散结果
		q1 = q2
		q2 = tmp
	}
	return -1
}

func openLockBfs(deadends []string, target string) int {
	var step int
	visited := make(map[string]struct{})
	for _, v := range deadends {
		visited[v] = struct{}{}
	}
	queue := make([]string, 0)
	queue = append(queue, "0000")

	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			if cur == target {
				return step
			}
			if _, has := visited[cur]; has {
				continue
			}
			visited[cur] = struct{}{}

			for j := 0; j < 4; j++ {
				lk := plusOnce(j, cur)
				if _, has := visited[lk]; !has {
					queue = append(queue, lk)
				}

				lk = minusOnce(j, cur)
				if _, has := visited[lk]; !has {
					queue = append(queue, lk)
				}
			}
		}
		step++
		queue = queue[size:]
	}
	return -1
}
func plusOnce(i int, lock string) string {
	ch := []byte(lock)
	if ch[i] == '9' {
		ch[i] = '0'
	} else {
		ch[i] += 1
	}
	return string(ch)
}

func minusOnce(i int, lock string) string {
	ch := []byte(lock)
	if ch[i] == '0' {
		ch[i] = '9'
	} else {
		ch[i] -= 1
	}
	return string(ch)
}

// @lc code=end

func Test_openLock(t *testing.T) {
	type args struct {
		deadends []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"8888"}, "0000"}, 0},
		{"1", args{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202"}, 6},
		{"1", args{[]string{"8888"}, "0009"}, 1},
		{"1", args{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openLock(tt.args.deadends, tt.args.target); got != tt.want {
				t.Errorf("openLock() = %v, want %v", got, tt.want)
			}
		})
	}
}
