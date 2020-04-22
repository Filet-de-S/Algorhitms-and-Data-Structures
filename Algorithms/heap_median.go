package main

import (
	"fmt"
	"sync"
)

type Heap interface {
	Insert(int)
	GetMedian() int
}

// max is 1234 [max] [median] [min] 6789
type heaps struct {
	max []int // start with [1] value of root;
	maxLen int // len have to be always == len(arr) -1
	min []int
	minLen int
}

func HeapInit() Heap {
	return &heaps{
		max:    make([]int, 1),
		maxLen: 0,
		min:    make([]int, 1),
		minLen: 0,
	}
}

func (h *heaps) Insert(n int) {
	m := &sync.Mutex{}
	m.Lock()
	if h.maxLen == h.minLen {
		if h.minRoot() != 0 && n > h.minRoot() { // n > median
			newMaxRoot := h.minRootExtract()
			h.maxInsert(newMaxRoot)
			h.minInsert(n)
		} else {
			h.maxInsert(n)
		}
	} else if n < h.maxRoot() {
		newMinRoot := h.maxRootExtract()
		h.minInsert(newMinRoot)
		h.maxInsert(n)
	} else { // n > h.minRoot()
		h.minInsert(n)
	}
	m.Unlock()
}

func (h *heaps) GetMedian() int {
	if h.maxLen < 1 {
		return 0
	} else if h.maxLen == h.minLen {
		return (h.maxRoot() + h.minRoot()) / 2
	} else {
		return h.maxRoot()
	}
}

func (h *heaps) bubbleUp(i int, flag byte) { // flag: [x] for max; [n] for min;
											// (0 1 2 [max] [median] [min] 5 6 7)
	if i == 1 {
		return
	}
	var parent, val *int
	if flag == 'n' { // min
		val = &h.min[i]
		parent = &(h.min)[i/2]
		if *val < *parent {
			*val, *parent = *parent, *val
		}
	} else { // max
		val = &h.max[i]
		parent = &(h.max)[i/2]
		if *parent < *val {
			*parent, *val = *val, *parent
		}
	}
	h.bubbleUp(i/2, flag)

}

func (h *heaps) bubbleDown(i int, flag byte) {	// flag: [x] for max; [n] for min;
												// (0 1 2 [max] [median] [min] 5 6 7)
	val := i
	lhs := i*2
	var arr *[]int

	if flag == 'n' {
		for j := 0; j < 2; j++ {
			side := lhs+j
			if side <= h.minLen && h.min[val] > h.min[side] {
				val = side
			}
		}
		arr = &h.min
	} else {
		for j := 0; j < 2; j++ {
			side := lhs+j
			if side <= h.maxLen && h.max[val] < h.max[side] {
				val = side
			}
		}
		arr = &h.max
	}
	if val != i {
		(*arr)[val], (*arr)[i] = (*arr)[i], (*arr)[val]
		h.bubbleDown(val, flag)
	}
}

func (h *heaps) maxRoot() int {
	if h.maxLen > 0 {
		return h.max[1]
	}
	return 0
}

func (h *heaps) maxInsert(n int) {
	h.max = append(h.max, n)
	h.maxLen++
	h.bubbleUp(h.maxLen, 1)
}

func (h *heaps) maxRootExtract() (r int) {
	r = h.maxRoot()
	h.max[1] = h.max[h.maxLen]
	h.max = h.max[:h.maxLen]
	h.maxLen--
	h.bubbleDown(1, 'x')
	return r
}

func (h *heaps) minRoot() int {
	if h.minLen > 0 {
		return h.min[1]
	}
	return 0
}

func (h *heaps) minInsert(n int) {
	h.min = append(h.min, n)
	h.minLen++
	h.bubbleUp(h.minLen, 'n')
}

func (h *heaps) minRootExtract() (r int) {
	r = h.minRoot()
	h.min[1] = h.min[h.minLen]
	h.min = h.min[:h.minLen]
	h.minLen--
	h.bubbleDown(1, 'n')
	return r
}

func main() {
 h:= HeapInit()
 data := []int{7, 5, 1, 9, 0, 2, 3, 6, 4, 8}

 for i := range data {
 	h.Insert(data[i])
 	fmt.Println("cycle n:",i, "; inserted new value:", data[i], "; median is:", h.GetMedian())
 }
}
