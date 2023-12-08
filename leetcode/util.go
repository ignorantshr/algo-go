package leetcode

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// go1.21 内置了 max, min 函数，不再需要了
// type compare interface {
// 	~int | ~int64
// }

// func max[T compare](a, b T) T {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }
