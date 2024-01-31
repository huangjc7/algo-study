package binarySearch

import (
	"testing"
)

func TestBsearch(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		value  int
		result int
	}{
		{
			name:   "重复值查找第一个给定值",
			arr:    []int{1, 2, 3, 3, 3, 3, 3, 4, 5, 5, 5, 5, 5, 6, 7, 8, 9, 9, 9, 9, 9, 9, 9},
			value:  5,
			result: 8,
		},
		{
			name:   "给定一个不存在的值 是否会数组越界",
			arr:    []int{1, 2, 3, 3, 3, 3, 3, 4, 5, 5, 5, 5, 5, 6, 7, 8, 9, 9, 9, 9, 9, 9, 9},
			value:  20,
			result: -1,
		},
	}

	for _, v := range tests {
		got := Bsearch(v.arr, v.value)
		if got != v.result {
			t.Errorf("Test '%s' failed: binarySearch() = %d, want %d", v.name, got, v.result)

		}

	}
}
