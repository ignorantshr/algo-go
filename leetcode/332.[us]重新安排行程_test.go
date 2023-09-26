/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 *
 * https://leetcode.cn/problems/reconstruct-itinerary/description/
 *
 * algorithms
 * Hard (47.81%)
 * Likes:    842
 * Dislikes: 0
 * Total Accepted:    94.5K
 * Total Submissions: 197.3K
 * Testcase Example:  '[["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]'
 *
 * 给你一份航线列表 tickets ，其中 tickets[i] = [fromi, toi]
 * 表示飞机出发和降落的机场地点。请你对该行程进行重新规划排序。
 *
 * 所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK
 * 开始。如果存在多种有效的行程，请你按字典排序返回最小的行程组合。
 *
 *
 * 例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前。
 *
 *
 * 假定所有机票至少存在一种合理的行程。且所有的机票 必须都用一次 且 只能用一次。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
 * 输出：["JFK","MUC","LHR","SFO","SJC"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：tickets =
 * [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
 * 输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
 * 解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"] ，但是它字典排序更大更靠后。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * tickets[i].length == 2
 * fromi.length == 3
 * toi.length == 3
 * fromi 和 toi 由大写英文字母组成
 * fromi != toi
 *
 *
 */
package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// @lc code=start
func findItinerary(tickets [][]string) []string {
	voyages := make(map[string]*destination332) // 起飞机场，[]{目的机场}
	for _, tickticket := range tickets {
		if _, has := voyages[tickticket[0]]; !has {
			voyages[tickticket[0]] = &destination332{&pair332{tickticket[1], false}}
		} else {
			*voyages[tickticket[0]] = append(*voyages[tickticket[0]], &pair332{tickticket[1], false})
		}
	}
	path := make([]string, 0)

	for _, destination := range voyages {
		sort.Sort(destination)
	}

	path = append(path, "JFK")
	var backtrace func(src string) bool
	backtrace = func(src string) bool {
		if len(path) == len(tickets)+1 {
			return true
		}

		for _, pair := range *voyages[src] {
			if pair.visited {
				continue
			}

			path = append(path, pair.dest)
			pair.visited = true
			if backtrace(pair.dest) {
				return true
			}
			path = path[:len(path)-1]
			pair.visited = true
		}
		return false
	}
	backtrace("JFK")
	return path
}

type pair332 struct {
	dest    string
	visited bool
}

type destination332 []*pair332

func (a destination332) Len() int           { return len(a) }
func (a destination332) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a destination332) Less(i, j int) bool { return a[i].dest < a[j].dest }

// @lc code=end

func Test_findItinerary(t *testing.T) {
	tests := []struct {
		name    string
		tickets [][]string
		want    []string
	}{
		{"1,", [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}, []string{"JFK", "MUC", "LHR", "SFO", "SJC"}},
		{"1,", [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}, []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findItinerary(tt.tickets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findItinerary() = %v, want %v", got, tt.want)
			}
		})
	}
}
