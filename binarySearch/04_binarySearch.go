package binarySearch

//变体三：查找第一个大于等于给定值的元素

//变体三：查找第一个大于等于给定值的元素

func BinarySearchFirstGT(arr []int, value int) int {
	if len(arr) == 0 {
		return -1
	}
	low := 0
	high := len(arr) - 1
	for low <= high {
		midd := low + (high-low)>>1
		if value > arr[midd] {
			low = midd + 1
		} else if value < arr[midd] {
			high = midd - 1
		} else {
			if midd != len(arr)-1 || arr[midd+1] > value {
				return midd + 1
			} else {
				low = midd + 1
			}
		}
	}

	return -1
}
