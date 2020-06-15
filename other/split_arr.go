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
