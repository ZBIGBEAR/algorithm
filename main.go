package main

import (
	"algorithm/other"
	"math/rand"
	"algorithm/tree"
	"fmt"
	"time"
)

func main() {
	arr1 := []int64{1,2,3,4,5,6,7,4,3}
	head1 := tree.NewBinaryTree(arr1)
	
	arr2 := []int64{2,4,5}
	head2 := tree.NewBinaryTree(arr2)

	if (tree.IsSubTree(head1, head2)) {
		fmt.Println("是子树")
	}else{
		fmt.Println("不是子树")
	}
	
}

func createArr(n int) []int64 {
	arr := make([]int64,n)
	for i:=0;i<n;i++{
		arr[i] = rand.Int63n(100)
	}
	return arr
}

func TestRevertListN(n int) {
	var head, p *other.Node
	for i := 0; i < n; i++ {
		node := &other.Node{V: int64(i)}
		if i == 0 {
			head = node
			p = node
		} else {
			p.Next = node
			p = node
		}
	}
	TestReverseList(head)
	fmt.Println("\n")
}

func TestReverseList(head *other.Node) {
	p := other.RevreseList(head)
	q := p
	for {
		if q == nil {
			break
		}
		fmt.Print(q.V, ",")
		q = q.Next
	}
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
