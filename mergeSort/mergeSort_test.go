package mergeSort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "排序测试",
			arr:  []int{4, 3, 2, 1},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "空数组测试",
			arr:  []int{},
			want: []int{},
		},
		{
			name: "单个数字测试",
			arr:  []int{1},
			want: []int{1},
		},
		{
			name: "顺序测试",
			arr:  []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, v := range tests {
		MergeSort(v.arr, len(v.arr))
		if !reflect.DeepEqual(v.arr, v.want) {
			t.Errorf("MergeSort() = %v, want %v", v.arr, v.want)
		}
	}

}
