package binarySearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name  string
		arr   []int
		value int
		want  int
	}{
		{
			name:  "左半区查找",
			arr:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			value: 3,
			want:  2,
		},
		{
			name:  "右半区查找",
			arr:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			value: 8,
			want:  7,
		},
		{
			name:  "溢出测试",
			arr:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			value: 10,
			want:  -1,
		},
	}

	for _, v := range tests {
		got := BinarySearch(v.arr, v.value)
		if got != v.want {
			t.Errorf("Test '%s' failed: binarySearch() = %v, want %v", v.name, got, v.want)
		}
	}
}
