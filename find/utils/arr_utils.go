package utils

// todo 数组旋转
func RotateN(arr []int64, n int) []int64 {
	first := 0
	cur := 0
	curVal := arr[cur]
	nextVal := arr[cur]
	l := len(arr)
	if n <= 0 {
		return arr
	}
	for {
		next := cur - n
		if next < 0 {
			next += l
		}
		nextVal = arr[next]
		arr[next] = curVal
		cur = next
		curVal = nextVal
		if cur == first {
			break
		}
	}
	return arr
}

// 交换数组两个数
func SwapArr(arr []int64, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
