package main

import "fmt"

// [4,3,2,1]
func MergeSort(a []int, n int) {
	// MergeSort_c([4,3,2,1], 0, 3)
	MergeSort_c(a, 0, n-1)
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

func Merge(arr []int, p, q, r int) {
	// i是左边起始下标 j是中位下标 k是临时数组的起始下标 r是末尾的下标
	i, j, k := p, q+1, 0
	tmp := make([]int, r-p+1)

	for i <= q && j <= r {
		//左右对比
		if arr[i] < arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++

		}
		k++
	}
	//[2,3]
	// 判断子数组的数据
	// 如果左侧子数组还有剩余，将其复制到tmp
	for i <= q {
		tmp[k] = arr[i]
		i++
		k++
	}

	// 如果右侧子数组还有剩余，将其复制到tmp
	for j <= r {
		tmp[k] = arr[j]
		j++
		k++
	}
	// 合并到原来数组
	//将tmp拷贝回arr
	for i := 0; i < k; i++ {
		arr[p+i] = tmp[i] //p+i p等于初始下标，p接收的是初始下标也会接收q+1的下标 所以p+i是左右起始的下标加上i新增的下标来赋值
	}
}

func main() {
	arr := []int{4, 3, 2, 1}
	MergeSort(arr, len(arr))
	fmt.Println(arr)
}
