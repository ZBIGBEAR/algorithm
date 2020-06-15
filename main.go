package main

import (
	"algorithm/other"
	"fmt"
	"time"
)

func main() {
	TestSplitArrByThree()
}

func TestSplitArrByOddEven() {
	arr := []int64{4, 8, 6, 4, 3, 76, 8, 9, 77, 44, 3, 6, 4, 434, 567, 223, 668, 4348, 5, 7}
	arr = other.SplitArr(arr, other.IsOdd)
	fmt.Println(arr)
}

func TestSplitArrByZero() {
	arr := []int64{4, -8, 6, 4, 3, -76, 8, -9, -77, -44, 3, 6, 4, 6, 7, -98, 3, 21, 435, 67, 23}
	arr = other.SplitArr(arr, other.SmallTHanZero)
	fmt.Println(arr)
}

func TestSplitArrByThree() {
	arr := []int64{4, -8, 6, 4, 3, -76, 8, -9, -77, -44, 3, 6, 4, 6, 7, -98, 3, 21, 435, 67, 23}
	arr = other.SplitArr(arr, other.IsDivisibleByThree)
	fmt.Println(arr)
}

func TestPrint1ToMaxN() {
	bTime := time.Now()
	other.Print1ToMaxN(7)
	fmt.Printf("%d\n", time.Now().Sub(bTime).Seconds())
}

func TestTimeCost() {
	// 3.5s
	TimeCost(TestOddEvenNumber, 100, 10000000000)

	// 6.9s
	TimeCost(TestOddEvenNumber1, 100, 10000000000)

	// 奇怪：为什么用与运算耗时比取模运算耗时还多？
}

func TestOddEvenNumber(n, count int) {
	for i := 0; i < count; i++ {
		other.IsEvenNumber(n)
	}
}

func TestOddEvenNumber1(n, count int) {
	for i := 0; i < count; i++ {
		other.IsEvenNumber1(n)
	}
}

func TimeCost(f interface{}, n, count int) {
	bTime := time.Now()
	f.(func(int, int))(n, count)
	cost := time.Now().Sub(bTime).Seconds()
	fmt.Printf("func:%s, %v\n", f, cost)
}
