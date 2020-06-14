package other

import (
	"fmt"
	"strconv"
)

// 打印从1到最大的n位十进制数。比如n=3,则打印1，2，3，...999
// 陷阱：如果n非常大，比如n=10000000,怎么办？

func Print1ToMaxN(n int) {
	N := maxN(n)
	next := "1"
	for {
		if CompareStrNumber(next, N) {
			break
		} else {
			printI(next, N)
			next = Incremental(next)
		}
	}
}

func printI(i string, max string) {
	// 只打印最后10个数字
	if len(i) != len(max) {
		return
	}
	for index := range i {
		if index == len(i)-1 {
			fmt.Printf("%s\n", i)
			return
		}
		if i[index] == max[index] {
			continue
		} else {
			return
		}
	}
}

func Incremental(n string) string {
	ns := Str2Strs(n)
	l := len(ns)
	index := l - 1
	carry := 1
	for {
		if index < 0 {
			break
		}
		tmp := Convert10(ns[index])
		if tmp == 9 {
			ns[index] = "0"
		} else {
			ns[index] = fmt.Sprintf("%d", tmp+carry)
			break
		}
		index--
	}
	result := Strs2Str(ns)
	if index < 0 {
		return fmt.Sprintf("%d%s", 1, result)
	} else {
		return result
	}
}

func Convert10(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Str2Strs(s string) []string {
	result := make([]string, len(s))
	for i, v := range s {
		result[i] = string(v)
	}
	return result
}

func Strs2Str(ss []string) string {
	result := ""
	for i := range ss {
		result += ss[i]
	}
	return result
}

func maxN(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += fmt.Sprintf("%d", 9)
	}
	return result
}

// 比较两个字符串代表的数字的大小
func CompareStrNumber(dest, src string) bool {
	lDesc := len(dest)
	lSrc := len(src)
	if lDesc > lSrc {
		return true
	} else if lDesc < lSrc {
		return false
	} else {
		descStrs := Str2Strs(dest)
		srcStrs := Str2Strs(src)
		for i := 0; i < lDesc; i++ {
			descStrsI := Convert10(descStrs[i])
			srcStrsI := Convert10(srcStrs[i])
			if descStrsI > srcStrsI {
				return true
			} else if descStrsI <= srcStrsI {
				return false
			}
		}
		return false
	}
}
