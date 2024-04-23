package main

import "fmt"

func buildHeap(a []int, n int) {
	for i := n >> 1; i > 0; i-- {
		heapify(a, n, i)
	}
}

func heapify(a []int, n, i int) {
	for {
		maxPos := i
		//left
		//    18
		//   /  \
		//  20  21
		if i<<1 <= n && a[i] < a[i<<1] {
			maxPos = i << 1
		}
		//这里使用maxPos是因为左边比较以后发生交换还需要跟右边比较一次
		if i<<1+1 <= n && a[maxPos] < a[i<<1+1] {
			maxPos = i<<1 + 1
		}
		//当上述条件都没有出发则会触发下面的if
		if maxPos == i {
			break
		}
		swap(a, i, maxPos)
		i = maxPos

	}
}

func swap(a []int, i, maxPos int) {
	a[i], a[maxPos] = a[maxPos], a[i]
}

func sort(a []int, n int) {
	buildHeap(a, n)
	k := n
	for k > 1 {
		swap(a, 1, k)
		k--
		heapify(a, k, 1)
	}

}

func main() {
	a := []int{0, 22, 6, 3, 1, 23, 65, 34, 21, 7, 3, 20}
	sort(a, len(a)-1)
	fmt.Println(a)
}
