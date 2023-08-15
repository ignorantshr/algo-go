package sort

// https://www.hello-algo.com/chapter_sorting/quick_sort/

func quickSort(arr []int) {
	_quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, i, j int) {
	if i >= j {
		return
	}
	m := partition(arr, i, j)
	_quickSort(arr, i, m-1)
	_quickSort(arr, m+1, j)
}

func partition(arr []int, left, right int) int {
	i := left
	j := right
	for {
		//1、2 顺序不可颠倒，要保证 arr[j] <= arr[left]，
		//顺序颠倒时的异常用例：arr = []int{6, 23, 9, 54, 2, 3, 1, 1, 5}
		for i < j && arr[j] >= arr[left] { // 1。 先从右向左找
			j--
		}
		for i < j && arr[i] <= arr[left] { // 2。再从左向右找
			i++
		}

		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	arr[j], arr[left] = arr[left], arr[j]
	return j
}

func quickSortBaseNumImprove(arr []int) {
	_quickSort(arr, 0, len(arr)-1)
}

func _quickSortBaseNumImprove(arr []int, i, j int) {
	if i >= j {
		return
	}
	m := partition(arr, i, j)
	_quickSort(arr, i, m-1)
	_quickSort(arr, m+1, j)
}

func partitionBaseNumImprove(arr []int, left, right int) int {
	m := getMiddle(arr, left, right)
	arr[m], arr[left] = arr[left], arr[m]

	i := left
	j := right
	for {
		for i < j && arr[j] >= arr[left] { // 1。 先从右向左找
			j--
		}
		for i < j && arr[i] <= arr[left] { // 2。再从左向右找
			i++
		}

		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	arr[j], arr[left] = arr[left], arr[j]
	return j
}

func getMiddle(arr []int, left, right int) int {
	mid := left + (right-left)/2
	// 此处使用异或运算来简化代码（!= 在这里起到异或的作用）
	// 异或规则为 0 ^ 0 = 1 ^ 1 = 0, 0 ^ 1 = 1 ^ 0 = 1
	if (arr[left] < arr[mid]) != (arr[left] < arr[right]) {
		return left
	} else if (arr[mid] < arr[left]) != (arr[mid] < arr[right]) {
		return mid
	}
	return right
}

func quickSortTailCallImprove(arr []int) {
	_quickSortTailCallImprove(arr, 0, len(arr)-1, 0)
}

func _quickSortTailCallImprove(arr []int, left, right, deep int) {
	//name := RandStringBytes(5)
	//pre := ""
	//for i := 0; i < deep; i++ {
	//	pre += "\t"
	//}
	//fmt.Printf("%v%v enter sort [%v, %v]\n", pre, name, left, right)
	//defer func() {
	//	fmt.Printf("%v%v exist sort [%v, %v]\n", pre, name, left, right)
	//}()
	// 子数组长度为 1 时终止
	for left < right {
		//fmt.Printf("%vloop, [%v, %v]\n", pre, left, right)
		// 哨兵划分操作
		pivot := partition(arr, left, right)
		// 对两个子数组中较短的那个执行快排
		if pivot-left < right-pivot {
			//fmt.Printf("%vsort left, [%v, %v, %v]\n", pre, left, pivot, right)
			_quickSortTailCallImprove(arr, left, pivot-1, deep+1) // 递归排序左子数组
			left = pivot + 1                                      // 剩余待排序区间为 [pivot + 1, right]
		} else {
			//fmt.Printf("%vsort right, [%v, %v, %v]\n", pre, left, pivot, right)
			_quickSortTailCallImprove(arr, pivot+1, right, deep+1) // 递归排序右子数组
			right = pivot - 1                                      // 剩余待排序区间为 [left, pivot - 1]
		}
	}
}
