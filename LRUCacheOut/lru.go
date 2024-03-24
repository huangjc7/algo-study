package LRUCacheOut

import (
	"fmt"
)

// Hashable 是一个接口，要求实现它的类型有一个 HashCode 方法。
type Hashable interface {
	HashCode() int
}

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

// String 是一个示例类型，用来表示包含字符串值和缓存的哈希码的结构体。
type String struct {
	value []rune // 字符数组
	hash  int    // 缓存的哈希码，初始值为0，表示还没有计算
}

// HashCode 实现了 Hashable 接口，计算并返回字符串的哈希码。
// 如果哈希码已经计算过（即不为0），直接返回缓存的值。
func (s *String) HashCode() int {
	if s.hash == 0 && len(s.value) > 0 {
		for _, c := range s.value {
			s.hash = 31*s.hash + int(c)
		}
	}
	return s.hash
}

// Hash 函数根据提供的键（必须实现了 Hashable 接口）和散列表的容量，
// 计算并返回散列表的索引。
func Hash(key Hashable, capacity int) int {
	h := key.HashCode()
	if capacity <= 0 {
		panic("capacity must be greater than 0")
	}
	// (h ^ (h >>> 16)) & (capacity -1) 的Go语言实现
	//99302346 & 15
	//101111010110011101111001010 & 000000000000000000000001111
	//10111101011 ^ 10111101011
	return (h ^ (h >> 16)) & (capacity - 1)
}

func add(head, newNode *Node) {
	current := head

	for {
		if current.next == nil {
			// 记录当前指针
			break
		} else {
			current = current.next
		}
	}
	current.next = newNode
	current.next.prev = current

}

func query(head **Node, key int) int {
	if *head == nil {
		fmt.Println("List is empty")
		return -1
	}

	var tmp *Node
	current := *head

	for current != nil {

		if current.prev == nil {
			//暂存当前节点
			tmp = current.next

			if tmp.key == key {
				return tmp.value
			}
		}

		if current.key == key {
			// if the current node is not head node ,move to head
			// 4 -> 5
			current.prev.next = current.next
			// 5 <- 4
			if current.next != nil {
				current.next.prev = current.prev
			}

			// current node move to the head
			current.next = tmp
			current.prev = *head

			tmp.prev = current

			(*head).next = current
			(*head).prev = nil
			return current.value
		}
		current = current.next
	}
	return current.value
}

func delete(head *Node, key int) {
	for current := head; current != nil; current = current.next {
		if current.key == key {
			current.prev.next = current.next
			current.next.prev = current.prev
		}
	}

}

func queryHashTable(types []rune, key int, head *Node) (int, bool) {
	str := &String{
		value: types,
	}
	capacity := 16 // 假设散列表的大小为16
	index := Hash(str, capacity)
	hashTable := make([]*Node, 16)
	hashTable[index] = head
	result := query(&head, key)
	fmt.Printf("查找类型: %c\n对应散列槽位: %d\n散列表存放的头指针地址: %p\n原始的头指针地址 %p\n对应值: %d\n",
		types, index, hashTable[index], head, result)
	return result, true
}

func Traverse(head *Node) {
	current := head
	for current != nil {
		fmt.Println(current.key, current.value, "上一个内存地址", current.prev, "下一个内存地址", current.next)
		current = current.next
	}
}

//func main() {
//	head := &Node{}
//	xh := &Node{
//		key:   1,
//		value: 175,
//	}
//	hjc := &Node{
//		key:   2,
//		value: 180,
//	}
//	lz := &Node{
//		key:   3,
//		value: 170,
//	}
//	dl := &Node{
//		key:   4,
//		value: 155,
//	}
//	xx := &Node{
//		key:   50,
//		value: 1000,
//	}
//
//	add(head, xh)
//	add(head, hjc)
//	add(head, lz)
//	add(head, dl)
//	add(head, xx)
//
//	//query(&head, 2)
//	//delete(head, 2)
//
//	qType := []rune("friend")
//	queryHashTable(qType, 3, head) // 这里传入头节点 单纯是因为懒不想改造使用成员函数来来数据传递
//	//Traverse(head)
//
//}
