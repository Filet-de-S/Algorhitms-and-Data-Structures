package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkQuickSortRev(b *testing.B)  {
	arr := make([]int, 0, 10)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i < 10; i++{
		arr = append(arr,i)// rand.Intn(1000000))
	}
	now := time.Now()
	qSortRev(arr)
	s := time.Since(now)
	b.Log(arr)
	if b.N < 10000 {
		b.Log(b.N, "				", s / time.Duration(b.N))
	} else {
		b.Log(b.N, "			", s / time.Duration(b.N))
	}
	for i := range arr {
		if i > 0 && arr[i] > arr[i-1] {
			b.Log("Not Sorted")
		}
	}
}

func qSortRev(arr []int) {
	if len(arr) < 2 {
		return
	}

	lh, mid, rh, pivot := 0, len(arr)/2, len(arr)-1, 0
	switch {
	case arr[lh] > arr[rh] && arr[lh] > arr[mid]:
		pivot = lh
	case arr[rh] > arr[lh] && arr[rh] > arr[mid]:
		pivot = rh
	default:
		pivot = mid
	}
	arr[rh], arr[pivot] = arr[pivot], arr[rh]

	for i := range arr {
		if arr[i] > arr[rh] {
			arr[lh], arr[i] = arr[i], arr[lh]
			lh++
		}
	}
	arr[lh], arr[rh] = arr[rh], arr[lh]

	qSortRev(arr[:lh])
	qSortRev(arr[lh+1:])
	//fmt.Println(arr)
}

// todo try iterative