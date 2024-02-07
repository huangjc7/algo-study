package binarySearch

//变体三：查找第一个大于等于给定值的元素

//变体三：查找第一个大于等于给定值的元素

// 变体四: 查找最后一个小于等于给定值的元素
func BinarySearchLastLT(arr []int, value int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		midd := low + (high-low)>>1
		if value > arr[midd] {
			low = midd + 1
		} else if value < arr[midd] {
			high = midd - 1
		} else {
			if arr[midd-1] < value {
				return midd - 1
			} else {
				high = midd - 1
			}
		}
	}

	return -1
}
