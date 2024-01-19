/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 *
 * https://leetcode.cn/problems/simplify-path/description/
 *
 * algorithms
 * Medium (44.42%)
 * Likes:    680
 * Dislikes: 0
 * Total Accepted:    210.6K
 * Total Submissions: 471.9K
 * Testcase Example:  '"/home/"'
 *
 * 给你一个字符串 path ，表示指向某一文件或目录的 Unix 风格 绝对路径 （以 '/' 开头），请你将其转化为更加简洁的规范路径。
 *
 * 在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..）
 * 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。任意多个连续的斜杠（即，'//'）都被视为单个斜杠 '/' 。
 * 对于此问题，任何其他格式的点（例如，'...'）均被视为文件/目录名称。
 *
 * 请注意，返回的 规范路径 必须遵循下述格式：
 *
 *
 * 始终以斜杠 '/' 开头。
 * 两个目录名之间必须只有一个斜杠 '/' 。
 * 最后一个目录名（如果存在）不能 以 '/' 结尾。
 * 此外，路径仅包含从根目录到目标文件或目录的路径上的目录（即，不含 '.' 或 '..'）。
 *
 *
 * 返回简化后得到的 规范路径 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：path = "/home/"
 * 输出："/home"
 * 解释：注意，最后一个目录名后面没有斜杠。
 *
 * 示例 2：
 *
 *
 * 输入：path = "/../"
 * 输出："/"
 * 解释：从根目录向上一级是不可行的，因为根目录是你可以到达的最高级。
 *
 *
 * 示例 3：
 *
 *
 * 输入：path = "/home//foo/"
 * 输出："/home/foo"
 * 解释：在规范路径中，多个连续斜杠需要用一个斜杠替换。
 *
 *
 * 示例 4：
 *
 *
 * 输入：path = "/a/./b/../../c/"
 * 输出："/c"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * path 由英文字母，数字，'.'，'/' 或 '_' 组成。
 * path 是一个有效的 Unix 风格绝对路径。
 *
 *
 */
package leetcode

import (
	"strings"
	"testing"
)

// @lc code=start
func simplifyPath(path string) string {
	stack := []string{}
	for _, v := range strings.Split(path, "/") {
		if v == "." || v == "" {
			continue
		}

		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, v)
	}
	return "/" + strings.Join(stack, "/")
}

func simplifyPath1(path string) string {
	path = path + "/" // 确保能解析到最后一段
	res := []string{}
	s := []rune{}
	for i, v := range path[1:] {
		if v == '/' {
			if path[i] == '/' { // 多重 /
				continue
			}

			str := string(s)
			s = s[:0]
			if str == "." {
				// skip
				continue
			}
			if str == ".." {
				if len(res) > 0 {
					res = res[:len(res)-1]
				}
				continue
			}
			res = append(res, "/"+str)
			continue
		}
		s = append(s, v)
	}

	if len(res) == 0 {
		return "/"
	}
	return strings.Join(res, "")
}

// @lc code=end

func Test_simplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"0", "/.", "/"},
		{"0.1", "/./", "/"},
		{"0.2", "/../", "/"},
		{"0.3", "/..", "/"},
		{"1", "/a", "/a"},
		{"2", "/abc/", "/abc"},
		{"3", "/aa///d//c", "/aa/d/c"},
		{"4", "/.../a", "/.../a"},
		{"5", "/./a", "/a"},
		{"6", "/a/./b/../../c/", "/c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
			if got := simplifyPath1(tt.path); got != tt.want {
				t.Errorf("simplifyPath1() = %v, want %v", got, tt.want)
			}
		})
	}
}
