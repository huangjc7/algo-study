package LRUCacheOut

import (
	"testing"
)

// 测试 add 函数能否正确地向链表中添加节点
func TestAdd(t *testing.T) {
	head := &Node{} // 创建一个空的头节点作为dummy head

	// 向链表中添加一个节点
	add(head, &Node{key: 1, value: 100})

	// 验证是否正确添加了节点
	if head.next == nil {
		t.Errorf("No node was added.")
	} else if head.next.value != 100 {
		t.Errorf("Node with value 100 was not added correctly. Got value %d", head.next.value)
	}
}

// 测试 queryHashTable 函数能否正确地查询并返回预期的结果
func TestQueryHashTable(t *testing.T) {
	head := &Node{key: 0} // 初始化头节点

	// 构建链表
	add(head, &Node{key: 1, value: 175})
	add(head, &Node{key: 2, value: 180})
	add(head, &Node{key: 3, value: 170})

	qType := []rune("friend")
	// 假设queryHashTable现在返回一个值和一个布尔标志
	value, found := queryHashTable(qType, 2, head)

	if !found {
		t.Errorf("Expected to find the key %d, but it was not found.", 1)
	}

	if value != 180 {
		t.Errorf("Expected value %d for key %d, got %d", 180, 2, value)
	}
}
