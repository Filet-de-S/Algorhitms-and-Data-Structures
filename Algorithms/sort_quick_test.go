package main

import (
	"encoding/gob"
	"fmt"
	"math/rand"
	"os"
	sort2 "sort"
	"sync"
	"testing"
	"time"
)

func qSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	lh, mid, rh, pivot := 0, len(arr) / 2, len(arr) - 1, 0
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
		if arr[i] < arr[rh] {
			arr[lh], arr[i] = arr[i], arr[lh] // 14232 5 768
			lh++
		}
	}
	arr[lh], arr[rh] = arr[rh], arr[lh]

	qSort(arr[:lh])
	if lh+1 < len(arr) {
		qSort(arr[lh+1:])
	}
}

func qSortGo(arr []int, wg *sync.WaitGroup) {
	if len(arr) < 2 {
		wg.Done()
		return
	}

	lh, mid, rh, pivot := 0, len(arr) / 2, len(arr) - 1, 0
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
		if arr[i] < arr[rh] {
			arr[lh], arr[i] = arr[i], arr[lh] // 14232 5 768
			lh++
		}
	}
	arr[lh], arr[rh] = arr[rh], arr[lh]

	wg.Add(1)
	go qSortGo(arr[:lh], wg)
	if lh+1 < len(arr) {
		wg.Add(1)
		go qSortGo(arr[lh+1:], wg)
	}
	wg.Done()
}
// todo try iterative

func BenchmarkQuickSortGo(b *testing.B)  {
	arr, err := readFromFile("testdataINTS")
	if err != nil {
		b.Fatal(err)
	}
	wg := &sync.WaitGroup{}
	now := time.Now()

	wg.Add(1)
	qSortGo(arr, wg)
	wg.Wait()

	s := time.Since(now)
	if b.N < 10000 {
		b.Log(b.N, "				", s / time.Duration(b.N))
	} else {
		b.Log(b.N, "			", s / time.Duration(b.N))
	}
	//for i := range arr {
	//	if i > 0 && arr[i] < arr[i-1] {
	//		b.Fatal("Not Sorted")
	//	}
	//}
}

func BenchmarkQuickSort(b *testing.B)  {
	arr, err := readFromFile("testdataINTS")
	if err != nil {
		b.Fatal(err)
	}
	now := time.Now()
	qSort(arr)
	s := time.Since(now)
	if b.N < 10000 {
		b.Log(b.N, "				", s / time.Duration(b.N))
	} else {
		b.Log(b.N, "			", s / time.Duration(b.N))
	}
	//for i := range arr {
	//	if i > 0 && arr[i] < arr[i-1] {
	//		b.Fatal("Not Sorted")
	//	}
	//}
}

func readFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dataDecoder := gob.NewDecoder(file)

	var data []int
	err = dataDecoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}


func TestQQQSort(t *testing.T) {
	arr := []int{}
	for i := 0; i < 1000000; i++ {
		arr = append(arr, rand.Intn(100))
	}
	dup := make([]int, len(arr))
	copy(dup, arr)
	start := time.Now()
	end := time.Since(start)

	start = time.Now()
	qSort(arr)
	end = time.Since(start)
	fmt.Println("qqq", end)

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
