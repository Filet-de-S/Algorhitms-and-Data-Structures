package main

import (
	"fmt"
	"time"
)

func main() {

	i := []map[int][]int{
		//{6:{2,4,5,7,8,10,20,25,61,124,0,1,2}}, // 0
		//{2:{3,4,5,6,7,1,2,2,2}}, // 0
		//{3:{3,4,5,6,7,1,2,2,}},
		//{1:{3,4,5,6,7,1,2}}, // 2
		//{7:{1,2,3,4,5,6,7,}},
		//{11:{6,7,8,9,10,11,12,2,3,4,5}}, // 4
		//{3:{2, 2, 2, 3, 4, 2}},
		//{1:{1, 2, 2, 2, 3, 4, 2}},
		//{9:{6,7,8,9,10,11,12,1,2,3,4,5}}, // 6
		//{2:{8,8,8,8,8}},
		//{1:{1,8,8,8,8}}, //8
		//{9:{18,19,20,21,22,22,1,2,3,4,5,6,7,8,9,10,11,12,13,14,16,17}}, // 12
		//{2:{5, 1, 2}},
		//{0:{0}}, // 14
		//{4:{4,2}},
		//{10:{5,10}}, //16
		//{2:{1,2,3}},
		//{1:{1,2,3}}, // 18
		{7:{4,4,4,5,7,1,2,3}},
	}
	for j := range i {
		for s := range i[j] {
			//now := time.Now()
			//v := searchElBinRotRecursive( i[j][s],0, len(i[j][s]) - 1, s)
			//since := time.Since(now)
			//fmt.Println("find", s ,"in", i[j], v, "\ntime_RECURSIVE:", since)

			now := time.Now()
			v := searchElBinRotIterative( i[j][s], s )
			since := time.Since(now)
			fmt.Println("find", s ,"in", i[j], v, "\ntime_ITERATIVE:", since)
		}
	}
}

func searchElBinRotIterative(arr []int, key int) int {
	l, h := 0, len(arr) - 1
	for {
		if l > h {
			return -1
		}
		mid := (l+h) / 2
		if arr[mid] == key {
			return mid
		}
		if arr[l] <= arr[mid] {
			if key < arr[mid] && key >= arr[l] {
				h = mid - 1
				continue
			}
			l = mid + 1
		} else if key > arr[mid] && key <= arr[h] {
			l = mid + 1
		} else { // 4445 7 123
			h = mid - 1
		}
	}
}

func searchElBinRotRecursive(arr []int, l int, h int, key int) int {
	if l > h {
		return -1
	}
	mid := (l+h) / 2
	if arr[mid] == key {
		return mid
	}

	if arr[l] <= arr[mid] {
		if key < arr[mid] && key >= arr[l] {
			return searchElBinRotRecursive(arr, l, mid - 1, key)
		}
		return searchElBinRotRecursive(arr, mid + 1, h, key)
	} else if key > arr[mid] && key <= arr[h] {
		return searchElBinRotRecursive(arr, mid + 1, h, key)
	}
	return searchElBinRotRecursive(arr, l, mid - 1, key)
}
