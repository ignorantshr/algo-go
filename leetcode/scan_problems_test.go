package leetcode

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// 比较字符串数字，数字必须是正数或 0
//
// n1 > n2: return 1; n1 == n2: return 0; n1 < n2: return -1;
func compareNumStr(s1, s2 string) int {
	s1 = strings.TrimLeft(s1, "0")
	s2 = strings.TrimLeft(s2, "0")
	pointLoc1 := strings.LastIndexByte(s1, '.')
	pointLoc2 := strings.LastIndexByte(s2, '.')

	compareInteger := func(a, b string) int {
		if len(a) < len(b) {
			return -1
		}
		if len(a) > len(b) {
			return 1
		}
		for i := 0; i < len(a); i++ {
			if a[i] > b[i] {
				return 1
			}
			if a[i] < b[i] {
				return -1
			}
		}
		return 0
	}

	compareDecimal := func(a, b string) int {
		a = strings.TrimRight(a, "0")
		b = strings.TrimRight(b, "0")

		i, j := 0, 0
		for ; i < len(a) && j < len(b); i, j = i+1, j+1 {
			if a[i] > b[j] {
				return 1
			}
			if a[i] < b[j] {
				return -1
			}
		}

		if j < len(b) {
			return 1
		}
		if i < len(a) {
			return -1
		}
		return 0
	}

	// 不是小数
	if pointLoc1 == -1 && pointLoc2 == -1 {
		return compareInteger(s1, s2)
	} else {
		if pointLoc1 == -1 {
			pointLoc1 = len(s1)
			s1 += "."
		}
		if pointLoc2 == -1 {
			pointLoc2 = len(s2)
			s2 += "."
		}
		res := compareInteger(s1[:pointLoc1], s2[:pointLoc2])
		if res != 0 {
			return res
		}
		return compareDecimal(s1[pointLoc1:], s2[pointLoc2:])
	}

}

// 给定一个范围，查找遗漏的题目
func TestScanLeakProblems(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}

	sort.Slice(entries, func(i, j int) bool {
		num1, err := strconv.Atoi(strings.SplitN(entries[i].Name(), ".", 2)[0])
		if err != nil {
			return false
		}

		num2, err := strconv.Atoi(strings.SplitN(entries[j].Name(), ".", 2)[0])
		if err != nil {
			return true
		}

		s1, s2 := strings.SplitN(entries[i].Name(), ".", 2)[0], strings.SplitN(entries[j].Name(), ".", 2)[0]
		res := compareNumStr(s1, s2)
		if num1 < num2 && res != -1 {
			fmt.Println(s1, s2, num1, num2, res)
		}
		if num1 == num2 && res != 0 {
			fmt.Println(s1, s2, num1, num2, res)
		}
		if num1 > num2 && res != 1 {
			fmt.Println(s1, s2, num1, num2, res)
		}
		return num1 < num2
	})

	// (l,e]
	pre := 0
	end := 300
	n := 0
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		raw := strings.SplitN(entry.Name(), ".", 2)[0]
		if num, err := strconv.Atoi(raw); err == nil {
			if num == pre || num == pre+1 {
				pre = num
				continue
			}
			pre++
			for pre < num && pre <= end {
				n++
				fmt.Printf("%d, ", pre)
				pre++
			}
			if pre <= end+1 {
				fmt.Println("[", n, "]")
			}
		}
	}

}

func Test_compareNumStr(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"小于", args{"1", "3"}, -1},
		{"小于", args{"1", "2"}, -1},
		{"大于", args{"2", "1"}, 1},
		{"大于", args{"1650", "1644"}, 1},
		{"等于", args{"1", "1"}, 0},
		{"小于 前缀", args{"01", "03"}, -1},
		{"大于 前缀", args{"003", "01"}, 1},
		{"大于 后缀", args{"00100", "010"}, 1},
		{"等于 前缀", args{"003", "03"}, 0},
		{"小于 小数 整数部分", args{"1.1", "2.1"}, -1},
		{"小于 小数 小数部分", args{"1.1", "1.20"}, -1},
		{"大于 小数 整数部分", args{"3.02", "2.12"}, 1},
		{"大于 小数 小数部分", args{"2.22", "2.12"}, 1},
		{"等于 小数 小数为空", args{"55.", "55."}, 0},
		{"等于 小数 小数为0", args{"55.0", "55.00"}, 0},
		{"等于 小数 整数为空", args{".23", ".23"}, 0},
		{"等于 小数 整数为0", args{"0.34", "0.34"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareNumStr(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("compareNumStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
