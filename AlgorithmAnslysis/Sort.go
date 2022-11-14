package main

import (
	"fmt"
	"test/AlgorithmAnslysis/util"
	"time"
)

func InsertSortA(arr []int) []int {
	strarttime := time.Now().UnixNano()
	len := len(arr)
	for i := 1; i < len; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr = util.SwapInt(arr, j, j-1)
			}
		}
	}
	endtime := time.Now().UnixNano()
	fmt.Printf("InsertSortA arr cast %v Nanosecond.\n", endtime-strarttime)
	return arr
}

func BubbleSortA(arr []int) []int {
	strarttime := time.Now().UnixNano()
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := len - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				util.SwapInt(arr, j, j-1)
			}
		}
	}
	endtime := time.Now().UnixNano()
	fmt.Printf("BubbleSortA arr cast %v Nanosecond.\n", endtime-strarttime)
	return arr
}

func SelectionSortA(arr []int) []int {
	strarttime := time.Now().UnixNano()
	len := len(arr)
	for i := 0; i < len; i++ {
		lowindex := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[lowindex] {
				lowindex = j
			}
		}
		util.SwapInt(arr, i, lowindex)
	}
	endtime := time.Now().UnixNano()
	fmt.Printf("BubbleSortA arr cast %v Nanosecond.\n", endtime-strarttime)
	return arr
}

func ShellSortA(arr []int) []int {
	strarttime := time.Now().UnixNano()
	len := len(arr)
	for gap := len / 2; gap > 0; gap /= 2 {
		for i := gap; i < len; i++ {
			for j := i - gap; j >= 0; j -= gap {
				if arr[j] > arr[j+gap] {
					util.SwapInt(arr, j, j+gap)
				}
			}
		}
	}
	endtime := time.Now().UnixNano()
	fmt.Printf("InsertSortA arr cast %v Nanosecond.\n", endtime-strarttime)
	return arr
}

func main() {
	//arr := []int{5, 32, 55, 32, 1, 9, 0, 45, 67, 8, 234, 356, 33, 2, 0, 99, 34}
	//arr := []int{}
	//arr := []int{1}
	arr := util.MakeArr(100, 0, 1000)
	//fmt.Println(arr)
	//_ = InsertSortA(arr)
	//_ = BubbleSortA(arr)
	//_ = SelectionSortA(arr)
	sortArr := ShellSortA(arr)
	fmt.Println(sortArr)
	//fmt.Println(sortArr2)
}
