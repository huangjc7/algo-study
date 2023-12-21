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
	//Merge([4, 3, 2, 1], 0, 1, r)
	Merge(a, p, q, r)
}

func Merge(arr []int, p, q, r int) []int {
	// i是左边起始下标 j是中位下标 k是临时数组的起始下标
	i, j, k := 0, q+1, 0
	tmp := make([]int, r-p)

	for i <= q && j <= p {
		//左右对比
		if arr[i] < arr[j] {
			tmp[k] = arr[i]
			i++
			k++
		} else {
			tmp[k] = arr[j]
			j++
			k++
		}
	}
	// 左边是否有多余数据
	if i <

	// 右边是否有多余数据
	// 合并到原来数组
}
