package main

import (
	"fmt"
)

type heap struct {
	count int
	a     []int
}

func newHeap(cap int) *heap {
	return &heap{
		a: make([]int, cap+1),
	}
}

func (h *heap) insert(data int) {
	h.count++
	h.a[h.count] = data
	i := h.count

	// from down to up heap
	for h.a[i] > h.a[i>>1] && i>>1 > 0 {
		h.a[i], h.a[i>>1] = h.a[i>>1], h.a[i]
		i = i >> 1
	}
}

func (h *heap) removeMax() {
	h.a[1] = h.a[h.count]
	h.count--
	i := 1
	//from up to down heap
	for {
		//largest = 2 i = 2
		largest := i
		left := i << 1
		right := i<<1 + 1

		// compare left
		if left <= h.count && h.a[largest] < h.a[left] {
			// largest = 4 i = 2
			largest = left
		}
		// compare right
		if right <= h.count && h.a[largest] < h.a[right] {
			largest = right
		}
		if largest == i {
			break
		}
		// a[2],a[4] = a[4],a[2]
		h.a[i], h.a[largest] = h.a[largest], h.a[i]
		// i = 4 2 = 4
		i = largest
	}
}

func main() {
	h := newHeap(9)
	h.insert(4)
	h.insert(10)
	h.insert(3)
	h.insert(5)
	h.insert(1)
	h.removeMax()
	fmt.Println(h.a)
}
