package utils

// 数组旋转。最简单的方法，空间复杂度为o(n)
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

// 递归的方法实现旋转数组。空间复杂度是O(1)
func RotateNByRecursion(arr []int64, n int) []int64 {
	rotateNByRecursion(arr, 0, len(arr), n)
	return arr
}

func rotateNByRecursion(arr []int64, begin, end, n int) []int64 {
	if end < 0 {
		return arr
	}

	if n == 0 {
		index := 0
		for i := 0; i < end; i++ {
			if i+begin >= end {
				arr[index] = -1
			} else {
				arr[index] = arr[i+begin]
			}
			index++
		}
		return arr
	}
	tmp := arr[begin]
	arr = rotateNByRecursion(arr, begin+1, end, n-1)
	for i := end - 1; i >= 0; i-- {
		if arr[i] == -1 {
			arr[i] = tmp
			return arr
		}
	}
	return arr
}
