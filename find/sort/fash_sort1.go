package sort

// 快速排序两外一种解法
// 先将小于基准元素的放到一边，就可以得到一边是小于基准元素的，一边是大于基准元素的
func MySort1(data []int) []int {
	mySort1(data, 0, len(data)-1)
	return data
}

func mySort1(data []int, begin, end int) {
	if begin >= end {
		return
	}
	i := begin
	for j := i; j <= end; j++ {
		if data[begin] > data[j] {
			i++
			swap(data, i, j)
		}
	}

	swap(data, begin, i)
	mySort1(data, begin, i-1)
	mySort1(data, i+1, end)
}

func swap(data []int, i, j int) {
	if i != j {
		data[i], data[j] = data[j], data[i]
	}
}
