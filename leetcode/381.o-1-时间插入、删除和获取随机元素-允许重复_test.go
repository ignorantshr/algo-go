/*
* @lc app=leetcode.cn id=381 lang=golang
*
* [381] O(1) 时间插入、删除和获取随机元素 - 允许重复
*
* https://leetcode.cn/problems/insert-delete-getrandom-o1-duplicates-allowed/description/
*
  - algorithms
  - Hard (42.33%)
  - Likes:    268
  - Dislikes: 0
  - Total Accepted:    26.6K
  - Total Submissions: 62.9K
  - Testcase Example:  '["RandomizedCollection","insert","insert","insert","getRandom","remove","getRandom"]\n' +
    '[[],[1],[1],[2],[],[1],[]]'

*
* RandomizedCollection 是一种包含数字集合(可能是重复的)的数据结构。它应该支持插入和删除特定元素，以及删除随机元素。
*
* 实现 RandomizedCollection 类:
*
*
* RandomizedCollection()初始化空的 RandomizedCollection 对象。
* bool insert(int val) 将一个 val 项插入到集合中，即使该项已经存在。如果该项不存在，则返回 true ，否则返回 false
* 。
* bool remove(int val) 如果存在，从集合中移除一个 val 项。如果该项存在，则返回 true ，否则返回 false 。注意，如果
* val 在集合中出现多次，我们只删除其中一个。
* int getRandom() 从当前的多个元素集合中返回一个随机元素。每个元素被返回的概率与集合中包含的相同值的数量 线性相关 。
*
*
* 您必须实现类的函数，使每个函数的 平均 时间复杂度为 O(1) 。
*
* 注意：生成测试用例时，只有在 RandomizedCollection 中 至少有一项 时，才会调用 getRandom 。
*
*
*
* 示例 1:
*
*
* 输入
* ["RandomizedCollection", "insert", "insert", "insert", "getRandom",
* "remove", "getRandom"]
* [[], [1], [1], [2], [], [1], []]
* 输出
* [null, true, false, true, 2, true, 1]
*
* 解释
* RandomizedCollection collection = new RandomizedCollection();// 初始化一个空的集合。
* collection.insert(1);   // 返回 true，因为集合不包含 1。
* ⁠                       // 将 1 插入到集合中。
* collection.insert(1);   // 返回 false，因为集合包含 1。
* // 将另一个 1 插入到集合中。集合现在包含 [1,1]。
* collection.insert(2);   // 返回 true，因为集合不包含 2。
* // 将 2 插入到集合中。集合现在包含 [1,1,2]。
* collection.getRandom(); // getRandom 应当:
* // 有 2/3 的概率返回 1,
* // 1/3 的概率返回 2。
* collection.remove(1);   // 返回 true，因为集合包含 1。
* // 从集合中移除 1。集合现在包含 [1,2]。
* collection.getRandom(); // getRandom 应该返回 1 或 2，两者的可能性相同。
*
*
*
* 提示:
*
*
* -2^31 <= val <= 2^31 - 1
* insert, remove 和 getRandom 最多 总共 被调用 2 * 10^5 次
* 当调用 getRandom 时，数据结构中 至少有一个 元素
*
*
*/
package leetcode

import (
	"fmt"
	"math/rand"
	"testing"
)

// @lc code=start
type RandomizedCollection struct {
	nums    []int
	indices map[int]map[int]struct{} // 二级结构使用 map 而不用 list 的好处就是不需要保证二级结构最后一位是 索引 最大值，直接就能对其操作
}

func Constructor381() RandomizedCollection {
	return RandomizedCollection{[]int{}, map[int]map[int]struct{}{}}
}

func (rs *RandomizedCollection) Insert(val int) bool {
	idxes, ok := rs.indices[val]
	if !ok {
		rs.indices[val] = make(map[int]struct{})
		idxes = rs.indices[val]
	}
	idxes[len(rs.nums)] = struct{}{}
	rs.nums = append(rs.nums, val)
	return !ok
}

func (rs *RandomizedCollection) Remove(val int) bool {
	idxes, ok := rs.indices[val]
	if !ok || len(idxes) == 0 {
		return false
	}
	var idx int // 旧数据的一个索引
	for idx = range idxes {
		break
	}
	lastIdx := len(rs.nums) - 1
	lastVal := rs.nums[lastIdx]

	rs.nums[idx] = lastVal // 旧位置赋值新数据
	idxes2 := rs.indices[lastVal]
	delete(idxes, idx)      // 删掉旧索引
	delete(idxes2, lastIdx) // 删除最后一个索引
	if idx < lastIdx {
		idxes2[idx] = struct{}{} // 添加旧数据索引
	}

	if len(idxes) == 0 {
		delete(rs.indices, val)
	}
	rs.nums = rs.nums[:lastIdx] // 删掉尾部
	return true
}

func (rs *RandomizedCollection) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
func Test_RandomizedCollection(t *testing.T) {
	type args struct {
		operations []string
		vals       []int
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		// {"1", args{
		// 	[]string{"insert", "insert", "insert", "insert", "insert", "remove", "remove", "remove", "insert", "remove", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"},
		// 	[]int{1, 1, 2, 2, 2, 1, 1, 2, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// }, []any{true, false, true, false, false, true, true, true, true, true, 1, 2, 2, 2, 1, 1, 2, 1, 2, 1}},
		{"1", args{
			[]string{"insert", "insert", "insert", "remove", "insert", "remove", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"},
			[]int{2, 2, 2, 2, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}, []any{true, false, false, true, true, true, 1, 2, 2, 2, 1, 1, 2, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := help381(tt.args.operations, tt.args.vals, tt.want); !got {
				t.Errorf("RandomizedCollection %v", got)
			}
		})
	}
}

func help381(operations []string, vals []int, want []any) bool {
	res := []any{}
	c := Constructor381()
	for i, op := range operations {
		fmt.Println(c)
		fmt.Println(i, op, vals[i])
		switch op {
		case "insert":
			res = append(res, c.Insert(vals[i]))
		case "remove":
			res = append(res, c.Remove(vals[i]))
		case "getRandom":
			res = append(res, c.GetRandom())
		}
	}

	for i := range res {
		if res[i] != want[i] {
			return false
		}
	}
	return true
}
