package binarySearch

// 二分查找变种 ----查找第一个值等于给定值的元素
// 思路
// 1. 循环找到任意位置重复数值，通过缩紧high来查找到重复值的第一个
// 2. 在循环里面进行判断arr[midd-1] != value 前一个是否等于查找的值，如果不等于则打印当前midd值 以此盘判断是否为第一个值

func Bsearch(arr []int, value int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		midd := low + (high-low)>>1
		if value > arr[midd] {
			low = midd + 1
		} else if value <= arr[midd] {
			high = midd - 1
		}
	}
	// 其实可以直接retrun low的值
	// 为了严谨做了二次判断
	//low < len(arr)  防止low溢出
	if low < len(arr) && arr[low] == value {
		return low
	}
	return -1
}
