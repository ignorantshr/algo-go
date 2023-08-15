package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var table = []struct {
	name string
	sort func([]int)
}{
	{"bubbleSort", bubbleSort},
	{"bubbleSortImprove", bubbleSortImprove},
	{"insertionSort", insertionSort},
	{"quickSort", quickSort},
	{"quickSortBaseNumImprove", quickSortBaseNumImprove},
	{"quickSortTailCallImprove", quickSortTailCallImprove},
	{"mergeSortFromTop", mergeSortFromTop},
	{"mergeSortFromBottom", mergeSortFromBottom},
}

func TestSort(t *testing.T) {
	//ori := []int{1, 54, 1, 5, 23, 2, 9, 3, 6}
	ori := genArr()

	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			var arr []int
			defer func() {
				if e := recover(); e != nil {
					fmt.Println(arr)
					fmt.Println(e)
				}
			}()
			t.Parallel()
			arr = make([]int, len(ori))
			copy(arr, ori)
			for i := 0; i < 20; i++ {
				shuffle(arr)
				v.sort(arr)
				if !checkSort(arr) {
					t.Fatalf("sort failed\n method:%v\n before:%v\n after:\t%v\n", v.name, ori, ori)
				}
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	arr := genArr()

	for _, v := range table {
		b.Run(v.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				shuffle(arr)
				v.sort(arr)
			}
		})
	}
}

func TestQuickSortTailCallImprove(t *testing.T) {
	//arr := []int{0, 6, 9, 5, 5, 3, 0, 0, 5, 0, 9, 7, 3, 4, 2, 2, 1, 3, 4, 7}
	//arr := genCustomArr(10, 10)
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	t.Log(arr)
	quickSortTailCallImprove(arr)
	t.Log(arr)
}

func TestMergeSort(t *testing.T) {
	arr := []int{62, 71, 6, 86, 7}
	t.Log(arr)
	mergeSortFromBottom(arr)
	t.Log(arr)
}
