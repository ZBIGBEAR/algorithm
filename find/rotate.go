package find

/*
求旋转数组的最小数字
*/
func RotateArrayFindSmall(arr []int64) int64 {
	count := len(arr)
	var index1, index2 int = 0, count - 1
	var mid int = 0
	for {
		if arr[index1] < arr[index2] {
			break
		}
		if index1+1 == index2 {
			mid = index2
			break
		}
		mid = (index1 + index2) / 2
		if arr[mid] >= arr[index1] {
			index1 = mid
		} else if arr[mid] <= arr[index2] {
			index2 = mid
		}

	}
	return arr[mid]
}
