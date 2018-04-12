package main

import (
	"fmt"
)

type maxBinHeap struct {
	baseArr []int
}

func (mbh *maxBinHeap) getFatherValue(index int) int {
	if index != 0 {
		return mbh.baseArr[int((index+1)/2)-1]
	} else {
		return -1
	}
}

func (mbh *maxBinHeap) getLChildValue(index int) int {
	if 2*index+1 < len(mbh.baseArr) {
		return mbh.baseArr[2*index+1]
	} else {
		return -1
	}
}

func (mbh *maxBinHeap) getRChildValue(index int) int {
	if 2*index+2 < len(mbh.baseArr) {
		return mbh.baseArr[2*index+2]
	} else {
		return -1
	}
}

func (mbh *maxBinHeap) getFatherIndex(index int) int {
	if index != 0 {
		return int((index+1)/2) - 1
	} else {
		return -1
	}
}

func (mbh *maxBinHeap) getLChildIndex(index int) int {
	if 2*index+1 < len(mbh.baseArr) {
		return 2*index + 1
	} else {
		return -1
	}
}

func (mbh *maxBinHeap) getRChildIndex(index int) int {
	if 2*index+2 < len(mbh.baseArr) {
		return 2*index + 2
	} else {
		return -1
	}
}

func (mbh *maxBinHeap) maxHeapify(indexNeed2Fix int) {
	lChildIndex := mbh.getLChildIndex(indexNeed2Fix)
	rChildIndex := mbh.getRChildIndex(indexNeed2Fix)
	largestIndex := indexNeed2Fix

	if lChildIndex != -1 && mbh.baseArr[lChildIndex] > mbh.baseArr[indexNeed2Fix] {
		largestIndex = lChildIndex
	}

	if rChildIndex != -1 && mbh.baseArr[rChildIndex] > mbh.baseArr[largestIndex] {
		largestIndex = rChildIndex
	}

	if largestIndex != indexNeed2Fix {
		key := mbh.baseArr[largestIndex]
		mbh.baseArr[largestIndex] = mbh.baseArr[indexNeed2Fix]
		mbh.baseArr[indexNeed2Fix] = key

		mbh.maxHeapify(largestIndex)
	}
}

func (mbh *maxBinHeap) heapSort() []int {
	var result []int
	for len(mbh.baseArr) > 2 {
		tempAddOn := []int{mbh.baseArr[0]}
		result = append(tempAddOn, result...)
		mbh.baseArr[0] = mbh.baseArr[len(mbh.baseArr)-1]
		mbh.baseArr = mbh.baseArr[0 : len(mbh.baseArr)-1]
		mbh.maxHeapify(0)
	}

	tr0 := []int{mbh.baseArr[0]}
	tr1 := []int{mbh.baseArr[1]}
	result = append(tr0, result...)
	result = append(tr1, result...)

	return result
}

func makeAMaxBinHeapFromArr(arr []int) maxBinHeap {
	result := maxBinHeap{arr}
	lastFatherIndex := result.getFatherIndex(len(result.baseArr) - 1)

	for i := lastFatherIndex; i >= 0; i-- {
		result.maxHeapify(i)
	}

	return result
}

func main() {
	arr := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

	fmt.Println("at first we have a arr in a mess.")
	fmt.Println("arr:", arr)

	result := makeAMaxBinHeapFromArr(arr)
	arr = result.heapSort()

	fmt.Println("now we have a sorted arr.")
	fmt.Println("arr:", arr)
}
