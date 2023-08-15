package leetcode

import (
	"reflect"
	"testing"
)

/*
https%3A//leetcode-cn.com/problems/range-addition

假设你有一个长度为 n 的数组，初始情况下所有的数字均为 0，你将会被给出 k 个更新的操作。

其中，每个操作会被表示为一个三元组：[startIndex, endIndex, inc]，你需要将子数组 A[startIndex ... endIndex]（包括 startIndex 和 endIndex）增加 inc。

请你返回 k 次操作后的数组。

示例:

输入: length = 5, updates = [[1,3,2],[2,4,3],[0,2,-2]]
输出: [-2,0,3,5,3]
解释:

初始状态:
[0,0,0,0,0]

进行了操作 [1,3,2] 后的状态:
[0,2,2,2,0]

进行了操作 [2,4,3] 后的状态:
[0,2,5,5,3]

进行了操作 [0,2,-2] 后的状态:
[-2,0,3,5,3]
*/

// 差分数组
type differArray struct {
	differs []int // differs[i] = nums[i] - nums[i-1], differs[0] = nums[0]
}

func newDifferArray(nums []int) *differArray {
	size := len(nums)
	if size == 0 {
		return &differArray{differs: []int{}}
	}

	differs := make([]int, size)
	differs[0] = nums[0]
	for i := 1; i < size; i++ {
		differs[i] = nums[i] - nums[i-1]
	}
	return &differArray{differs: differs}
}

// 用 val 更新原数组的 [a,b] 区间
func (d *differArray) update(a, b, val int) {
	if a >= 0 && a < len(d.differs) {
		d.differs[a] += val
	}
	if b >= 0 && b < len(d.differs)-1 {
		d.differs[b+1] -= val
	}
}

// 恢复原数组
func (d *differArray) restore() []int {
	size := len(d.differs)
	if size == 0 {
		return []int{}
	}

	nums := make([]int, size)
	nums[0] = d.differs[0]
	for i := 1; i < size; i++ {
		nums[i] = d.differs[i] + nums[i-1]
	}
	return nums
}

func getModifiedArray(length int, updates [][]int) []int {
	differ := newDifferArray(make([]int, length))
	for _, u := range updates {
		differ.update(u[0], u[1], u[2])
	}
	return differ.restore()
}

func Test_getModifiedArray(t *testing.T) {
	type args struct {
		length  int
		updates [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0", args{
			0,
			[][]int{},
		}, []int{}},
		{"1", args{
			5,
			[][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}},
		}, []int{-2, 0, 3, 5, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getModifiedArray(tt.args.length, tt.args.updates); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getModifiedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
