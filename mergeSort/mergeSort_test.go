package mergeSort

import (
	"reflect"
	"testing"
)

// TestMergeSort 测试 mergeSort 函数
func TestMergeSort(t *testing.T) {
	//输入预期
	input := []int{4, 2, 3, 1}
	//预期结果
	want := []int{1, 2, 3, 4}

	MergeSort(input, len(input))

	if !reflect.DeepEqual(input, want) {
		t.Errorf("mergeSort(%v) = %v, want %v", input, input, want)
	}
}
