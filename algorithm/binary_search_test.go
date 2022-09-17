package algorithm

import (
	"fmt"
	"testing"
)

// 二分查找递归解法
func BinarySearch(array []int, target int, l, r int) int {

	if l > r {
		// 出界了 找不到
		return -1
	}

	//从中间开始找
	mid := (l + r) / 2
	middleNum := array[mid]
	if middleNum == target {
		return mid //找到了
	} else if middleNum > target {
		// 中间的数比目标还大，从左边找
		return BinarySearch(array, target, 0, mid-1)
	} else {
		// 中间的数比目标还小，从右边找
		return BinarySearch(array, target, mid+1, r)
	}
}

func BinarySearch2(array []int, target, l, r int) int {
	ltemp := l
	rtemp := r
	for {
		if ltemp > rtemp {
			return -1
		}

		//从中间开始找
		mid := (ltemp + rtemp) / 2
		middleNum := array[mid]
		if middleNum == target {
			return mid // 找到了
		} else if middleNum > target {
			// 中间的数比目标还大，从左边找
			rtemp = mid - 1
		} else {
			// 中间的数比目标还小，从右边找
			ltemp = mid + 1
		}
	}

}

func sum(n int) int {
	total := ((1 + n) * n) / 2
	return total

}

func TestSum(t *testing.T) {
	fmt.Println(sum(100))
}

// 二分数

func TestBinarySearch(t *testing.T) {
	array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	target := 500
	result := BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(target, result)

	target = 189
	result = BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(target, result)
	target = 500
	result1 := BinarySearch2(array, target, 0, len(array)-1)
	fmt.Println(target, result1)
	target = 189
	result2 := BinarySearch2(array, target, 0, len(array)-1)
	fmt.Println(target, result2)
}
