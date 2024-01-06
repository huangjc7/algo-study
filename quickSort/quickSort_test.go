package quickSort

import (
	"reflect"
	"testing"
)

func TestquickSort(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		want []int
	}{
		{
			name: "正常排序测试",
			A:    []int{4, 2, 1, 3, 2, 10},
			want: []int{1, 2, 2, 3, 4, 10},
		},
		{
			name: "单个字符测试",
			A:    []int{1},
			want: []int{1},
		},
		{
			name: "空字符测试",
			A:    []int{},
			want: []int{},
		},
		{
			name: "顺序测试",
			A:    []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, v := range tests {
		quickSort(v.A, 0, len(v.A)-1)
		if reflect.DeepEqual(v.A, v.want) {
			t.Errorf("MergeSort() = %v, want %v", v.A, v.want)
		}
	}
}
