package leetcode

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	slice := make([]map[string]string, 0)
	slice = []map[string]string{{"0": "0"}, {"1": "0"}, {"2": "0"}, {"3": "0"}, {"4": "0"}, {"5": "0"}, {"6": "0"}}
	t.Log(slice)
	slice = updateSlice(slice, 0)
	t.Log(slice)
	slice = updateSlice(slice, 1)
	t.Log(slice)
	slice = updateSlice(slice, 2)
	t.Log(slice)
}

func updateSlice(slice []map[string]string, n int) []map[string]string {
	if n < len(slice) {
		fmt.Println((slice)[:n], (slice)[n+1:])
		// tmp := append([]map[string]string{}, (slice)[:n]...)
		// tmp = append(tmp, (slice)[n+1:]...)
		// return tmp
		slice = append(slice[:n], slice[n+1:]...)
		return slice
	}
	return slice
}

func updateSlicePoint(slice *[]map[string]string, n int) {
	if n < len(*slice) {
		fmt.Println((*slice)[:n], (*slice)[n+1:])
		(*slice) = append((*slice)[:n], (*slice)[n+1:]...)
	}
}
