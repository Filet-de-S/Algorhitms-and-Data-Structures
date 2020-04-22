package main

import (
	"fmt"
	"math/rand"
	sort2 "sort"
	"testing"
	"time"
)

func HeapSort(arr []int) {
	for k := len(arr)/2; k >= 0; k-- {
		bubbleDown(arr, k, len(arr)-1)
	}
	l := len(arr) - 1
	for l >= 0 {
		arr[0], arr[l] = arr[l], arr[0]
		l--
		bubbleDown(arr, 0, l)
	}
}

func bubbleDown(arr []int, k, l int) {
	for (k == 0 && 1 <= l) || (k != 0 && k*2 <= l) {
		child := k*2
		if k == 0 {
			child = 1
		}
		if child < l && arr[child+1] > arr[child] {
			child++
		}
		if arr[k] > arr[child] {
			break
		}
		arr[k], arr[child] = arr[child], arr[k]
		k = child
	}
}

func TestHeapSort(t *testing.T) {
	arr := []int{}
	for i := 0; i < 1000000; i++ {
		arr = append(arr, rand.Intn(100))
	}
	dup := make([]int, len(arr))
	copy(dup, arr)
	start := time.Now()
	end := time.Since(start)

	start = time.Now()
	HeapSort(arr)
	end = time.Since(start)
	fmt.Println("heap", end)

	start = time.Now()
	sort2.Ints(dup)
	end = time.Since(start)
	fmt.Println("def", end)

	for i := range dup {
		if dup[i] != arr[i] {
			t.Fatal(dup[i], arr[i], i)
		}
	}
}
