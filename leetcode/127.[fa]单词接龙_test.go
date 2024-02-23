/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 *
 * https://leetcode.cn/problems/word-ladder/description/
 *
 * algorithms
 * Hard (48.31%)
 * Likes:    1347
 * Dislikes: 0
 * Total Accepted:    205.3K
 * Total Submissions: 423.1K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 ->
 * s2 -> ... -> sk：
 *
 *
 * 每一对相邻的单词只差一个字母。
 * 对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
 * sk == endWord
 *
 *
 * 给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列
 * 中的 单词数目 。如果不存在这样的转换序列，返回 0 。
 *
 *
 * 示例 1：
 *
 *
 * 输入：beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log","cog"]
 * 输出：5
 * 解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
 *
 *
 * 示例 2：
 *
 *
 * 输入：beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log"]
 * 输出：0
 * 解释：endWord "cog" 不在字典中，所以无法进行转换。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= beginWord.length <= 10
 * endWord.length == beginWord.length
 * 1 <= wordList.length <= 5000
 * wordList[i].length == beginWord.length
 * beginWord、endWord 和 wordList[i] 由小写英文字母组成
 * beginWord != endWord
 * wordList 中的所有字符串 互不相同
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 建图 + BFS
	startIdx, endIdx, g := buildGraph127(beginWord, endWord, wordList)
	if endIdx == -1 {
		return 0
	}

	size := len(wordList)
	queue := []int{startIdx}
	used := make([]bool, size+1)
	n := 1

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i] == endIdx {
				return n
			}
			for _, v := range g.edges[queue[i]] {
				if !used[v] {
					queue = append(queue, v)
					used[v] = true
				}
			}
		}
		queue = queue[l:]
		n++
	}

	return 0
}

func buildGraph127(beginWord string, endWord string, wordList []string) (int, int, *graph127) {
	size := len(wordList)
	startIdx := size
	endIdx := -1
	g := &graph127{vertexNum: size + 1}

	for i := 0; i < size; i++ {
		if wordList[i] == beginWord {
			startIdx = i
		}
		if wordList[i] == endWord {
			endIdx = i
		}
		g.addVertex(i)
	}
	g.addVertex(startIdx)

	for i := 0; i < size; i++ {
		if oneDiffChar(wordList[i], beginWord) {
			g.addEdge(i, startIdx)
		}
		for j := i + 1; j < size; j++ {
			if j != i {
				if oneDiffChar(wordList[i], wordList[j]) {
					g.addEdge(i, j)
				}
			}
		}
	}

	return startIdx, endIdx, g
}

type graph127 struct {
	vertexNum int
	edges     [][]int
}

func (g *graph127) addVertex(s int) {
	if g.edges == nil {
		g.edges = make([][]int, g.vertexNum)
	}
	if g.edges[s] == nil {
		g.edges[s] = make([]int, 0)
	}
}

func (g *graph127) addEdge(s, e int) {
	if g.hasEdge(s, e) {
		return
	}

	g.edges[s] = append(g.edges[s], e)
	g.edges[e] = append(g.edges[e], s)
}

func (g *graph127) hasEdge(s, e int) bool {
	for _, v := range g.edges[s] {
		if v == e {
			return true
		}
	}
	return false
}

// 超时
func ladderLengthBacktrack(beginWord string, endWord string, wordList []string) int {
	path := []string{beginWord}
	size := len(wordList)
	used := make([]bool, size)
	ans := size + 2

	var backtrack func(word string)
	backtrack = func(word string) {
		if word == endWord {
			ans = min(ans, len(path))
			return
		}

		for i := 0; i < size; i++ {
			if used[i] || !oneDiffChar(path[len(path)-1], wordList[i]) {
				continue
			}

			path = append(path, wordList[i])
			used[i] = true
			backtrack(wordList[i])
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack(beginWord)

	if ans == size+2 {
		return 0
	}
	return ans
}

func oneDiffChar(s1, s2 string) bool {
	size1 := len(s1)
	if size1 != len(s2) {
		return false
	}

	for i := 0; i < size1; i++ {
		if s1[i] != s2[i] {
			return s1[i+1:] == s2[i+1:]
		}
	}

	return false
}

// @lc code=end

func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"x.1", args{
			"hot",
			"dog",
			[]string{"hot", "dog"},
		}, 0},
		{"x.2", args{
			"hog",
			"dog",
			[]string{"hog", "dog"},
		}, 2},
		{"0", args{
			"hit",
			"cog",
			[]string{"hot"},
		}, 0},
		{"1.1", args{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
		}, 5},
		{"1.2", args{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log"},
		}, 0},
		{"1.3", args{
			"hit",
			"hot",
			[]string{"hot", "dot", "dog", "lot", "log"},
		}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
			if got := ladderLengthBacktrack(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLengthBacktrack() = %v, want %v", got, tt.want)
			}
		})
	}
}
