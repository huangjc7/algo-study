package binarySearch

// binary search(二分查找)
func BinarySearch(arr []int, value int) int {
	////找到中位数
	low := 0
	high := len(arr) - 1

	for low <= high {
		midd := (low + high) / 2
		if arr[midd] == value {
			return midd
		} else if arr[midd] < value {
			low = midd + 1
		} else {
			high = midd - 1
		}
	}
	//表示未找到
	return -1
}
