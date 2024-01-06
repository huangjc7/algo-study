package quickSort

func partition(A []int, p int, r int) int {
	pivot := A[r]
	i := p
	for j := p; j < r; j++ {
		//i会一直保持在0
		//for循环负责持续获取数组的数字来和枢轴数进行比较 如果比枢轴数就保存在A[i]中
		if A[j] < pivot {
			//发现小于枢轴数就进行交换
			A[i], A[j] = A[j], A[i]
			i++
		}
	}
	//在一系列位置交换以后，数组中会有大于或等于枢轴数的树脂存在，无法进入if进行交换
	//所以该操作直接进行最后一次交换 这也使得快排是一种不稳定排序算法
	A[i], A[r] = A[r], A[i]
	//这个i的返回的是右边最后元素的下标 然后给上层递归右边开始计算
	return i
}

func quickSort(A []int, p int, r int) {
	if p < r {
		q := partition(A, p, r)
		quickSort(A, p, q-1)
		quickSort(A, q+1, r)
	}
}
