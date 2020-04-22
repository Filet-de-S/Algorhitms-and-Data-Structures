package main

import (
	"fmt"
	"sort"
)

func main() {
	in := []int{3,2,1,4,3,2,1,4,1} // any item can be sorted
	//inPlaceDupRmv(&in)
	in = noMemoryDupRmv(in)
	fmt.Println(in)
}

func noMemoryDupRmv(arr []int) []int {
	st := 0
	sort.Ints(arr)
	for i := range arr {
		if i > 0 && arr[i-1] == arr[i] {
			continue
		}
		arr[st] = arr[i]
		st++
	}
	arr = arr[:st]
	return arr
}
