package main

import (
	"math"
	"testing"
)

type bst struct {
	data int
	left *bst
	right *bst
}

func checkIsBST(b *bst) bool {
	if b == nil {
		return true
	}
	if !checkIsBSTSide(b, math.MinInt32, math.MaxInt32) {
		return false
	}
	return true
}

func checkIsBSTSide(head *bst, min, max int) bool {
	if head == nil {
		return true
	}
	if head.data <= min || head.data > max {
		return false
	}
	if !checkIsBSTSide(head.left, min, head.data) ||
		!checkIsBSTSide(head.right, head.data, max) {
		return false
	}
	return true
}

func bstCreateFromArr(arr []int) *bst {//, start, end int) *bst {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	mid := len(arr) / 2
	bst := &bst{}
	bst.data = arr[mid]
	bst.left = bstCreateFromArr(arr[:mid])
	bst.right = bstCreateFromArr(arr[mid+1:])
	return bst
}


func TestCheckIsBST(t *testing.T) {
	arr := []int{}
	for i := 0; i<8; i++ {
		arr = append(arr, i)
	}

	bst := bstCreateFromArr(arr)
	bst.left.right.data = 5
	if checkIsBST(bst) {
		t.Fatal()
	}
}

func TestCreateBSTfromArr(t *testing.T) {
	arr := []int{}
	for i := 0; i<8; i++ {
		arr = append(arr, i)
	}

	bst := bstCreateFromArr(arr)
	if !checkIsBST(bst) {
		t.Fatal()
	}
}
