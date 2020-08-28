package other

import (
	//"fmt"
	"github.com/pkg/errors"
)

// 找出数组中出现次数超过一半的数字
func FindMoreThanHalf(arr []int64) (int64, error) {
	if len(arr) == 0 {
		return 0, errors.New("input array is empty")
	}
	
	result := arr[0]
	count :=1
	for idx :=1;idx<len(arr);idx++ {
		if arr[idx] == result {
			count ++
		}else{
			count --
		}
		if count == 0 {
			result = arr[idx]
			count = 1
		}
	}
	return result,nil
}

// 找出数组第k小的数字
// 思路：参考快速排序
func FindKSmall(arr []int64, k int) (int64, error) {
	if len(arr) < k {
		return 0, errors.Errorf("arr count less than k%d",k)
	}
	begin := 0
	end := len(arr)-1
	midIndex := fastSort(arr,begin,end)
	for {
		if midIndex>k{
			end = midIndex-1
		}else if midIndex<k{
			begin = midIndex+1
		}else{
			break
		}
		midIndex = fastSort(arr,begin,end)
	}
	return arr[midIndex],nil
}

func fastSort(arr []int64,begin,end int) int{
	if begin>=end{
		return end
	}
	midIndex := begin
	i,j := begin+1,end
	// defer func(){
	// 	fmt.Println("=============i,j:",i,j)
	// }()
	for{
		for {
			if j > midIndex && arr[j]>=arr[midIndex]{
				j--
				continue
			}
			break
		}
		if j<=midIndex{
			return midIndex
		}
		swap(arr,j,midIndex)
		midIndex=j

		for{
			if i<midIndex && arr[i]<=arr[midIndex]{
				i++
				continue
			}
			break
		}
		if i>=midIndex{
			return midIndex
		}
		swap(arr,i,midIndex)
		midIndex=i
	}
	return midIndex
}

func swap(arr []int64, i,j int){
	tmp := arr[i]
	arr[i]=arr[j]
	arr[j]=tmp
}