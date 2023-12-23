package mergeSort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{4, 3, 2, 1}
	want := []int{1, 2, 3, 4}

	MergeSort(arr, len(arr))
	if !reflect.DeepEqual(arr, want) {
		t.Errorf("MergeSort() = %v, want %v", arr, want)
	}
}
