package mergeSort

// [4,3,2,1]
func MergeSort(a []int, n int) {
	// MergeSort_c([4,3,2,1], 0, 3)
	MergeSort_c(a, 0, len(a)-1)
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

func Merge(a []int, p, q, r int) []int {
	//i 给左边数据当初始下标
	//j 给右边数据当初始下标
	//k 给tmp []int 当初始下标
	i, j, k := p, q+1, 0

	tmp := make([]int, r-p+1)

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

	// 复制 left 切片中剩余的元素
	for i <= q {
		tmp[k] = a[i]
		i++
		k++
	}
	// 复制 right 切片中剩余的元素

	//存在问题
	for j <= p {
		tmp[k] = a[j]
		j++
		k++
	}

	//数据拷贝回原数组
	for i := 0; i <= r-p; i++ {
		a[p+i] = tmp[i]
	}
	return a
}
