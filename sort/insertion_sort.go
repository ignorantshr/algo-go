package sort

// 最差时间复杂度：n-1,...,1 = n(n-1+1)/2
//原地排序，稳定排序，自适应排序
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			base := arr[i]
			j := i - 1
			for ; j >= 0 && base < arr[j]; j-- {
				arr[j+1] = arr[j]
			}
			arr[j+1] = base
		}
	}
}
