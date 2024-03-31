package tree

import (
	"testing"
)

// 测试 BinarySearchTree 的 insert 和 find 方法
func TestBinarySearchTree(t *testing.T) {
	// 创建一个新的 BinarySearchTree
	bst := BinarySearchTree{}

	// 插入元素
	bst.insert(4)
	bst.insert(2)
	bst.insert(6)
	bst.insert(1)
	bst.insert(3)
	bst.insert(5)
	bst.insert(7)

	// 测试 find 方法
	testCases := []struct {
		value    int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{6, true},
		{7, true},
		{8, false},  // 不存在于树中
		{11, false}, // 不存在于树中
		{0, false},  // 不存在于树中
	}

	for _, testCase := range testCases {
		found := bst.find(testCase.value)
		if found != testCase.expected {
			t.Errorf("find(%d) = %v; want %v", testCase.value, found, testCase.expected)
		}
	}
}
