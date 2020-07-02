package other

// 将数组按指定规则分成两类
func SplitArr(arr []int64, f func(int64) bool) []int64 {
	begin, end := 0, len(arr)-1
	for {
		if begin >= end {
			break
		}
		for {
			if f(arr[begin]) {
				begin++
			} else {
				break
			}
		}

		for {
			if !f(arr[end]) {
				end--
			} else {
				break
			}
		}
		if begin >= end {
			break
		}
		arr[begin], arr[end] = arr[end], arr[begin]
	}
	return arr
}

func IsOdd(i int64) bool {
	return i&1 == 1
}

func SmallTHanZero(i int64) bool {
	return i < 0
}

func IsDivisibleByThree(i int64) bool {
	return i%3 == 0
}

// 对于一个给定的、乱序数组，把奇数都放前面，偶数都放后面。
// 要求：1.奇数、偶数相对位置不变。2.时间复杂度O（n），空间复杂度O（1）
func SplitArrByOddEven(arr []int64) []int64 {
	if len(arr) < 2 {
		return arr
	}
	// 找出奇数个数
	oddCount := 0
	for idx := range arr {
		if arr[idx]%2 == 1 {
			oddCount += 1
		}
	}
	oddIdx := oddCount - 1
	evenIdx := oddCount
	if evenIdx >= len(arr) {
		return arr
	}
	for {
		oddI := oddIdx
		evenJ := evenIdx
		for {
			if oddI >= 0 && arr[oddI]%2 == 1 {
				oddI -= 1
			} else {
				break
			}
		}

		for {
			if evenJ < len(arr) && arr[evenJ]%2 == 0 {
				evenJ += 1
			} else {
				break
			}
		}
		if oddI < 0 {
			break
		}
		if evenJ >= len(arr) {
			break
		}
		// 保存arr[oddI]
		tmp := arr[oddI]
		// 将arr[oddI+1,oddIdx]往左移动一个位置
		for x := oddI; x < oddIdx; x++ {
			arr[x] = arr[x+1]
		}
		// 将arr[evenJ]移动到arr[oddIdx]
		arr[oddIdx] = arr[evenJ]
		// 将arr[evenIdx, evenJ-1]往右移动一个位置
		for x := evenJ; x > evenIdx; x-- {
			arr[x] = arr[x-1]
		}
		// 将tmp移动到arr[evenIdx]
		arr[evenIdx] = tmp
	}
	return arr
}
