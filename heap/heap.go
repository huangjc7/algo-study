package heap

import "fmt"

type Heap struct {
	arr   []int
	count int
}

func newHeap() *Heap {

	return &Heap{
		arr:   make([]int, 10),
		count: 0,
	}
}

func (h *Heap) insert(data int) {
	h.count++
	h.arr[h.count] = data
	i := h.count

	if h.count < 2 {
		return
	}

	// 从下往上堆化
	//i/2 > 0 必须要大于0，如果没有大于0那么就会给index[0]赋值 不满足heap特性 是边界条件之一
	for i/2 > 0 && h.arr[i] > h.arr[i/2] {
		//swap
		h.arr[i], h.arr[i/2] = h.arr[i/2], h.arr[i]
		//持续往上比较
		i = i / 2
	}
}

func (h *Heap) removeMax() {
	h.arr[1] = h.arr[h.count] //最后元素移至堆顶
	h.count--

	i := 1

	for {
		maxPos := i
		//left
		if i*2 <= h.count && h.arr[i] < h.arr[i*2] {
			maxPos = i * 2
		}
		//right
		if i*2+1 <= h.count && h.arr[i] < h.arr[i*2+1] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}
		//swap
		h.arr[i], h.arr[maxPos] = h.arr[maxPos], h.arr[i]
		i = maxPos
	}
}

func main() {
	h := newHeap()

	h.insert(5)
	h.insert(7)
	h.insert(9)
	h.insert(100)
	h.insert(2)
	h.insert(4)
	h.insert(19)
	h.insert(1)
	h.removeMax()
	fmt.Println(h.arr)
}
