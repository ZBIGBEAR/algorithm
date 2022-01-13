package sort

// fast sort
func MySort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	begin, end := 0, len(data)-1

	i := begin
	j := end

	midIdx := i
	i += 1

	// todo 之前这里写成 i < j, 少了一个等于号，造成结果不对。排查了好久，当data只有2个元素的时候这里有问题，所以需要等于号
	for i <= j {
		for j > midIdx && data[j] >= data[midIdx] {
			j--
		}

		if j > midIdx {
			data[midIdx], data[j] = data[j], data[midIdx]
			midIdx = j
			j--
		}

		for i < midIdx && data[i] <= data[midIdx] {
			i++
		}

		if i < midIdx {
			data[midIdx], data[i] = data[i], data[midIdx]
			midIdx = i
			i++
		}
	}

	if midIdx-1 > begin {
		MySort(data[begin:midIdx])
	}

	if midIdx+1 < len(data) {
		MySort(data[midIdx+1 : end+1])
	}

	return data
}
