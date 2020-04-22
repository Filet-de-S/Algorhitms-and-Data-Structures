package main

import "fmt"

func main() {
	i := [][]int{
		{3,4,5,6,7,1,2,2,2}, // 0
		{3,4,5,6,7,1,2,2,},
		{3,4,5,6,7,1,2}, // 2
		{1,2,3,4,5,6,7,},
		{6,7,8,9,10,11,12,2,3,4,5}, // 4
		{2, 2, 2, 3, 4, 2},
		{1, 2, 2, 2, 3, 4, 2},
		{6,7,8,9,10,11,12,1,2,3,4,5}, // 6
		{8,8,8,8,8},
		{1,8,8,8,8}, //8
		{8,1,8,8,8},
		{8,8,1,8,8}, // 10
		{8,8,8,8,1},
		{18,19,20,21,22,22,1,2,3,4,5,6,7,8,9,10,11,12,13,14,16,17}, // 12
		{5, 1, 2},
		{0}, // 14
		{4,2},
		{5,10}, //16
		{1,2,3},
		{1,2,3,4,5,6,7,8,9,10,11,12,13,14,16,17, 18,19,20,21,22,22,}, // 18

	}
	for j := range i {
		fmt.Println("case", i[j])
		fmt.Println(findMinBinRot( i[j], 0, len(i[j]) - 1 ), j)
		fmt.Println(findMinBinRotIter( i[j]))
	}
}

func findMinBinRotIter(arr []int) int {
	l, h := 0, len(arr) - 1
	for {
		if l > h {
			return -1
		}
		mid := (l + h) / 2
		if mid == 0 {
			if arr[l] < arr[h] {
				return arr[l]
			}
			return arr[h]
		}
		if mid > 0 && arr[mid] < arr[mid-1] {
			return arr[mid]
		}

		if arr[mid] <= arr[h] { //sorted
			h = mid - 1
		} else { // if arr[mid] >= arr[h] {
			l = mid + 1
		}
	}
}

func findMinBinRot(arr []int, l int, h int ) int {
	if l > h {
		return -1
	}
	mid := (l+h) / 2
	if mid == 0 {
		if arr[l] < arr[h] {
			return arr[l]
		}
		return arr[h]
	}
	if mid > 0 && arr[mid] < arr[mid - 1] {
		return arr[mid]
	}

	if arr[mid] <= arr[h] { //sorted
		return findMinBinRot(arr, l, mid - 1)
	} else {// if arr[mid] >= arr[h] {
		return findMinBinRot(arr, mid + 1, h)
	}

}
//
//func findMinBinRot(arr []int) int {
//	right := len(arr) - 1
//	mid := right / 2
//	switch {
//	case right == 0:
//		return arr[0]
//	case right < 0: // err
// 		return -1
//	case right == 1: // 2elem
//		if arr[0] < arr[1] {
//			return arr[0]
//		} else {
//			return arr[1]
//		}
//	}
//
//	for {
//		if mid > 0 && arr[mid] < arr[mid-1] { // we found it in RIGHT
//			return arr[mid]
//		} else if right == 0 { // in left, FULLY SORTED
//			return arr[0]
//		}
//		if arr[mid] < arr[right] {
//			right = mid
//			mid /= 2
//		} else if arr[mid] > arr[right] { // 5 1 2
//			mid += mid / 2
//			if mid > right {
//				mid = right
//			}
//		} else {
//			right--
//		}
//	}
//}