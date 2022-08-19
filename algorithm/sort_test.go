package algorithm

import (
	"fmt"
	"testing"
)

func TestAlgorithm(t *testing.T) {

	selectSortValue := selectSort([]int{3, 5, 2, 1, 6})
	bubbleSortValue := bubbleSort([]int{3, 5, 2, 1, 6})
	insertSort1Value := insertSort1([]int{3, 5, 2, 1, 6})
	insertSort2Value := insertSort2([]int{3, 5, 2, 1, 6})

	fmt.Println("选择排序：", selectSortValue)
	fmt.Println("冒泡排序：", bubbleSortValue)
	fmt.Println("插入排序1：", insertSort1Value)
	fmt.Println("插入排序2：", insertSort2Value)
}

// 选择排序
func selectSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l; i++ {

		value := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[value] {
				value = j
			} else {
				value = value
			}

		}
		tmp := arr[value]
		arr[value] = arr[i]
		arr[i] = tmp

	}
	return arr
}

// 冒泡排序
func bubbleSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return nil
	}

	l := len(arr)
	for e := l - 1; e >= 0; e-- {
		for s := 1; s <= e; s++ {
			if arr[s-1] > arr[s] {
				tmp := arr[s-1]
				arr[s-1] = arr[s]
				arr[s] = tmp
				fmt.Println("==", tmp, e)

			}

		}
	}
	return arr
}

// 插入排序 1
func insertSort1(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return nil
	}
	l := len(arr)

	for i := 1; i < l; i++ {
		num := i
		for num-1 >= 0 && arr[num-1] > arr[num] {
			tmp := arr[num]
			arr[num] = arr[num-1]
			arr[num-1] = tmp
			num--
		}
	}
	return arr
}

// 插入排序 2 优化
func insertSort2(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return nil
	}
	l := len(arr)

	for i := 1; i < l; i++ {

		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			tmp := arr[j]
			arr[j] = arr[j+1]
			arr[j+1] = tmp
		}
	}
	return arr
}
