package other

import "fmt"

// 输出旋转矩阵
func PrintRorateMatrix(arr [][]int64) {
	PrintRorateMatrixWithLen(arr, 0 , len(arr)-1, 0)
}

// 根接每一行/列的<开始，结束>递归输出矩阵
func PrintRorateMatrixWithLen(arr [][]int64, begin, end, dep int) {
	if begin > end{
		return
	}
	if begin == end {
		fmt.Printf("%d\t", arr[begin][begin])
	}
	// 输出上边的行
	for i:= begin; i < end; i++ {
		fmt.Printf("%d\t", arr[dep][i])
	}
	
	l := len(arr)
	// 输出右边的列
	for i := begin; i < end; i++ {
		fmt.Printf("%d\t", arr[i][l-1-dep])
	}
	// 输出下边的行
	for i := end; i> begin; i-- {
		fmt.Printf("%d\t", arr[l-1-dep][i])
	}
	// 输出左边的列
	for i:= end; i> begin; i--{
		fmt.Printf("%d\t", arr[i][dep])
	}
	newDep := dep + 1
	PrintRorateMatrixWithLen(arr, begin+1, end-1, newDep)
}