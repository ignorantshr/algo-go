package sort

// 时间复杂度 O(nlogn)：划分形成高度为 logn 的递归树，每层合并的总操作数量为 n，总体使用 O(nlogn) 时间。
// 空间复杂度 O(n): 额外辅助空间 O(n),递归深度为 logn，使用 O(logn) 大小的栈帧空间
// 非原地排序：辅助数组需要使用 O(n) 额外空间。
// 稳定排序：在合并时可保证相等元素的相对位置不变。
// 非自适应排序：对于任意输入数据，归并排序的时间复杂度皆相同。
func mergeSortFromTop(arr []int) {
	_mergeFromTop(arr, 0, len(arr))
}

func _mergeFromTop(arr []int, left, right int) {
	if right-left <= 1 {
		return
	}
	mid := left + (right-left)/2
	_mergeFromTop(arr, left, mid)
	_mergeFromTop(arr, mid, right)
	merge(arr, left, mid, right)
}

func mergeSortFromBottom(arr []int) {
	_mergeFromBottom(arr, 0, len(arr))
}

func _mergeFromBottom(arr []int, left, right int) {
	for step := 1; step < (right - left); step *= 2 {
		for i := left; i+step < right && i < right; i += 2 * step {
			end := min(right, i+step*2)
			merge(arr, i, i+step, end)
		}
	}
}

func merge(arr []int, left, mid, right int) {
	if mid-1 >= left && arr[mid-1] < arr[mid] {
		return
	}
	a := make([]int, mid-left)
	b := make([]int, right-mid)
	copy(a, arr[left:mid])
	copy(b, arr[mid:right])
	for i, j := 0, 0; i < len(a) || j < len(b); {
		idx := left + i + j
		if i >= len(a) {
			arr[idx] = b[j]
			j++
			continue
		}

		if j >= len(b) {
			arr[idx] = a[i]
			i++
			continue
		}

		if a[i] < b[j] {
			arr[idx] = a[i]
			i++
		} else {
			arr[idx] = b[j]
			j++
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
