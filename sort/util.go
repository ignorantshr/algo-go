package sort

import "math/rand"

func genArr() []int {
	return genCustomArr(2000, 100)
}

func genCustomArr(n, ran int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(ran)
	}
	return arr
}

func shuffle(arr []int) {
	for i := len(arr) - 1; i >= 1; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
	//fmt.Println(arr)
}

func checkSort(arr []int) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
