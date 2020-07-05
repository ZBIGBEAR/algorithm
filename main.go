package main

import (
	"algorithm/other"
	"algorithm/tree"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	TestBinaryTree()
}

func TestBinaryTree() {
	for i := 0; i < 10; i++ {
		arr := createArr(i)
		TestBinaryTreeByArr(arr)
		fmt.Println("================================")
	}
}

func TestBinaryTreeByArr(arr []int64) {
	fmt.Printf("输出数组:\t")
	fmt.Println(arr)
	head := tree.NewBinaryTree(arr)

	fmt.Printf("递归前序遍历:\t")
	tree.PreOrderTraverseBinaryTree(head)
	fmt.Println()

	fmt.Printf("方法一：循环遍历前序:\t")
	tree.PreOrderTraverseBinaryTreeWithLoop1(head)
	fmt.Println()

	fmt.Printf("方法二：循环遍历前序:\t")
	tree.PreOrderTraverseBinaryTreeWithLoop2(head)
	fmt.Println()

	fmt.Printf("递归中序遍历：\t")
	tree.MidOrderTraverseBinaryTree(head)
	fmt.Println()

	fmt.Printf("循环中序遍历：\t")
	tree.MidOrderTraverseBinaryTreeWithLoop(head)
	fmt.Println()

	fmt.Printf("递归后续遍历：\t")
	tree.PostOrderTraverseBinaryTree(head)
	fmt.Println()

	fmt.Printf("循环后续遍历 方法一：\t")
	tree.PostOrderTraverseBinaryTreeWithLoop1(head)
	fmt.Println()

	fmt.Printf("循环后续遍历 方法二：\t")
	tree.PostOrderTraverseBinaryTreeWithLoop2(head)
	fmt.Println()
}

// 创建一个二维数组
func createMatrix(n int) [][]int64 {
	rand.Seed(time.Now().UnixNano())
	arr := make([][]int64, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int64, n)
		for j := 0; j < n; j++ {
			arr[i][j] = rand.Int63n(100)
		}
	}
	return arr
}

// 输出二维数组
func PrintMatrix(arr [][]int64) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%d\t", arr[i][j])
		}
		fmt.Println()
	}
}

func createArr(n int) []int64 {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int64, n)
	for i := 0; i < n; i++ {
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
