package utils

// todo 数组旋转
func RotateN(arr []int64, n int) []int64 {
	count := len(arr)
	newArr := make([]int64, count)
	index := 0
	for i := n; i < count; i++ {
		newArr[index] = arr[i]
		index++
	}
	for j := 0; j < n; j++ {
		newArr[index] = arr[j]
		index++
	}
	return newArr
}
