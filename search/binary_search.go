package search

// https://www.hello-algo.com/chapter_searching/binary_search/

// 二分查找, [0,n] 写法
func binarySearch1(arr []int, val int) int {
	i := 0
	j := len(arr) - 1

	for i <= j { // 搜索区域为空时跳出
		m := i + (j-i)/2 // 预防 (i+j)/2 写法在部分情况下因为数值过大超过 int 范围
		if arr[m] < val {
			i = m + 1
			continue
		}

		if val < arr[m] {
			j = m - 1
			continue
		}

		return m
	}

	return -1
}

// 二分查找, [0,n) 写法
func binarySearch2(arr []int, val int) int {
	i := 0
	j := len(arr)

	for i < j { // 搜索区域为空时跳出
		m := i + (j-i)/2 // 预防 (i+j)/2 写法在部分情况下因为数值过大超过 int 范围
		if arr[m] < val {
			i = m + 1
			continue
		}

		if val < arr[m] {
			j = m
			continue
		}

		return m
	}

	return -1
}
