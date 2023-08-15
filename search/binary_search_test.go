package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 4, 56, 233, 234, 345, 1000}

	t.Run("binarySearch1", func(t *testing.T) {
		fmt.Println(binarySearch1(arr, 1))
		fmt.Println(binarySearch1(arr, 3))
		fmt.Println(binarySearch1(arr, 1000))
		fmt.Println(binarySearch1(arr, 233))
		fmt.Println(binarySearch1(arr, 2))
	})

	t.Run("binarySearch2", func(t *testing.T) {
		fmt.Println(binarySearch2(arr, 1))
		fmt.Println(binarySearch2(arr, 3))
		fmt.Println(binarySearch2(arr, 1000))
		fmt.Println(binarySearch2(arr, 233))
		fmt.Println(binarySearch2(arr, 2))
	})
}
