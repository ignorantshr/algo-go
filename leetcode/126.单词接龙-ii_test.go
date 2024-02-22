/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 *
 * https://leetcode.cn/problems/word-ladder-ii/description/
 *
 * algorithms
 * Hard (37.41%)
 * Likes:    703
 * Dislikes: 0
 * Total Accepted:    58K
 * Total Submissions: 156.2K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 按字典 wordList 完成从单词 beginWord 到单词 endWord 转化，一个表示此过程的 转换序列 是形式上像 beginWord ->
 * s1 -> s2 -> ... -> sk 这样的单词序列，并满足：
 *
 *
 *
 *
 * 每对相邻的单词之间仅有单个字母不同。
 * 转换过程中的每个单词 si（1 <= i <= k）必须是字典 wordList 中的单词。注意，beginWord 不必是字典 wordList
 * 中的单词。
 * sk == endWord
 *
 *
 * 给你两个单词 beginWord 和 endWord ，以及一个字典 wordList 。请你找出并返回所有从 beginWord 到 endWord
 * 的 最短转换序列 ，如果不存在这样的转换序列，返回一个空列表。每个序列都应该以单词列表 [beginWord, s1, s2, ..., sk]
 * 的形式返回。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log","cog"]
 * 输出：[["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
 * 解释：存在 2 种最短的转换序列：
 * "hit" -> "hot" -> "dot" -> "dog" -> "cog"
 * "hit" -> "hot" -> "lot" -> "log" -> "cog"
 *
 *
 * 示例 2：
 *
 *
 * 输入：beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log"]
 * 输出：[]
 * 解释：endWord "cog" 不在字典 wordList 中，所以不存在符合要求的转换序列。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= beginWord.length <= 5
 * endWord.length == beginWord.length
 * 1 <= wordList.length <= 500
 * wordList[i].length == beginWord.length
 * beginWord、endWord 和 wordList[i] 由小写英文字母组成
 * beginWord != endWord
 * wordList 中的所有单词 互不相同
 *
 * 题解：https://leetcode.cn/problems/word-ladder-ii/solutions/277612/yan-du-you-xian-bian-li-shuang-xiang-yan-du-you--2/
 */
package leetcode

import (
	"os"
	"runtime/pprof"
	"testing"
)

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// BFS + DFS
	hasend, g := buildGraph126(beginWord, endWord, wordList)
	if !hasend {
		return nil
	}

	ans := make([][]string, 0)

	// 如果从 beginWord 开始寻找的话测试用例会超时
	var backtrack func(idx string, path []string)
	backtrack = func(idx string, path []string) {
		if idx == beginWord {
			tmp := []string{}
			for i := len(path) - 1; i >= 0; i-- {
				tmp = append(tmp, path[i])
			}
			ans = append(ans, tmp)
			return
		}

		for _, v := range g.edges[idx] {
			path = append(path, v)
			backtrack(v, path)
			path = path[:len(path)-1]
		}
	}
	backtrack(endWord, []string{endWord})

	return ans
}

// bfs 构建有向无权图
func buildGraph126(beginWord string, endWord string, wordList []string) (bool, *graph126) {
	g := &graph126{make(map[string][]string)}
	dict := make(map[string]bool)
	found := false

	for _, v := range wordList {
		dict[v] = true
	}
	if !dict[endWord] {
		return false, nil
	}
	// 删掉，防止回绕
	delete(dict, beginWord)

	steps := map[string]int{}
	step := 1
	steps[beginWord] = step
	queue := []string{beginWord}
	for len(queue) > 0 {
		step++
		l := len(queue)
		for i := 0; i < l; i++ {
			chars := []byte(queue[i])
			for j := 0; j < len(chars); j++ {
				old := chars[j]
				for c := byte('a'); c <= 'z'; c++ {
					if c == queue[i][j] {
						continue
					}
					chars[j] = c
					nextWord := string(chars)

					/* 下一个单词已被使用时的添加新边逻辑 */

					// 可能会有多个单词走到同一目的地，比如 log -> lot, hot -> lot
					// 但是 只能往高层走，同层或回头会绕远
					if steps[nextWord] != 0 && steps[nextWord] == step {
						g.addEdge(nextWord, queue[i])
					}

					/* 下一个单词未被使用时的添加新边逻辑 */

					// 拦截 已访问的单词 和 列表中不存在 的单词
					if !dict[nextWord] {
						continue
					}

					delete(dict, nextWord) //可直接删除，因为 steps 中已经有记录了

					// dict 和 steps 承担了判断是否已经访问的功能
					queue = append(queue, nextWord)

					g.addEdge(nextWord, queue[i])

					steps[nextWord] = step

					if nextWord == endWord {
						found = true
					}
				}
				chars[j] = old
			}
		}
		if found {
			break
		}
		queue = queue[l:]
	}

	return found, g
}

type graph126 struct {
	edges map[string][]string
}

func (g *graph126) addEdge(s, e string) {
	g.edges[s] = append(g.edges[s], e)
	// g.edges[e] = append(g.edges[e], s)
}

// @lc code=end

func Test_findLadders(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"x.1", args{
			"hot",
			"dog",
			[]string{"hot", "dog", "dot"},
		}, [][]string{{"hot", "dot", "dog"}}},
		{"1.1", args{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
		}, [][]string{{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"}}},
		{"1.3", args{
			"hit",
			"hot",
			[]string{"hot", "dot", "dog", "lot", "log"},
		}, [][]string{{"hit", "hot"}}},
		{"1.2", args{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log"},
		}, [][]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLadders(tt.args.beginWord, tt.args.endWord, tt.args.wordList); !equalSetMatrix(got, tt.want) {
				t.Errorf("findLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	// 创建 CPU pprof 文件
	cpuFile, _ := os.Create("cpu.pprof")
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	// 执行 benchmark
	for i := 0; i < b.N; i++ {
		findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	}

	// 创建内存 pprof 文件
	memFile, _ := os.Create("mem.pprof")
	pprof.WriteHeapProfile(memFile)
	defer memFile.Close()
}
