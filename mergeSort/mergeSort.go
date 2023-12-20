package mergeSort

// [4,3,2,1]
func MergeSort(a []int, n int) {
	// MergeSort_c([4,3,2,1], 0, 3)
	MergeSort_c(a, 0, len(a))
}

func MergeSort_c(a []int, p, r int) {
	// 0 >= 3
	// 第二次运算: 0 >= 1
	if p >= r {
		return
	}
	// q:= 2 0+3/2 q = 1
	//第二次运算: 0 +1 / 2 = 0 到了第三就满足了 if p >= r出发return条件
	q := (p + r) / 2
	//分治递归
	// MergeSort_c([4,3,2,1], 0, 1)
	MergeSort_c(a, p, q) //左边

	MergeSort_c(a, q+1, r) //右边
	//合并 [p...q]和[q+1...r]合并为[p...r]
	Merge(a, p, q, r)
}

func Merge(a []int, p, q, r int) {
	//i 给左边数据当初始下标
	//j 给右边数据当初始下标
	//k 给tmp []int 当初始下标
	i, j, k := p, q+1, 0

	var tmp []int

	// i =0  q = 5
	for i <= q && j <= r {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
		k++
	}

}
