package util

import (
	"math/rand"
	"time"
)

func SwapInt(arr []int, i int, j int) []int {
	var sortArr []int
	var p int

	sortArr = arr[:]
	p = sortArr[i]
	sortArr[i] = sortArr[j]
	sortArr[j] = p
	return sortArr
}

func MaxInt(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func MinInt(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func MakeArr(len int, min int, max int) []int {
	if min < 0 {
		return nil
	}
	var arr []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		arr = append(arr, rand.Intn(max-min)-min)
	}
	return arr
}
