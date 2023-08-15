package sort

/*
time complexity: n-1,...,1 = n(n-1+1)/2
原地排序，稳定排序，自适应排序
*/
func bubbleSort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 优化，最佳时间复杂度为 O(n)
func bubbleSortImprove(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		flag := true
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}
