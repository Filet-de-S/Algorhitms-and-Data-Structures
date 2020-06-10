package test

import (
	queue ".."
	"errors"
	"strconv"
	"testing"
)

func TestMinPQ(t *testing.T) {
	minPQ := queue.New("min")

	for i := 0; i < 10; i++ {
		minPQ.Insert(&queue.Data{CmpData: i, Data:"I+"+strconv.Itoa(i)})
	}
	for i := 0; i < 11; i++ {
		if i == 10 {
			_, err := minPQ.GetRoot()
			var e queue.EmptyHeap
			if _, err1 := minPQ.Pop(); !errors.As(err, &e) || !errors.As(err1, &e) {
				t.Fatal()
			}
		} else {
			v, err := minPQ.GetRoot()
			if v1, err1 := minPQ.Pop(); err != nil || err1 != nil || v != *v1 || v.CmpData != i {
				t.Fatal()
			}
		}

	}
}

func TestMaxPQ(t *testing.T) {
	minPQ := queue.New("max")

	for i := 0; i < 10; i++ {
		minPQ.Insert(&queue.Data{CmpData: i, Data:"I+"+strconv.Itoa(i)})
	}
	for i := 10; i >= 0; i-- {
		if i == 0 {
			_, err := minPQ.GetRoot()
			var e queue.EmptyHeap
			if _, err1 := minPQ.Pop(); !errors.As(err, &e) || !errors.As(err1, &e) {
				t.Fatal()
			}
		} else {
			v, err := minPQ.GetRoot()
			if v1, err1 := minPQ.Pop(); err != nil || err1 != nil || v != *v1 || v.CmpData != i-1 {
				t.Fatal()
			}
		}

	}
}