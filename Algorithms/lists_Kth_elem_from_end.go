package main

import "fmt"

type lists struct {
	val int
	next *lists
}

func main () {
	l := &lists{1, &lists{
		val:  2,
		next: nil,
	}}
	fmt.Println(findK(l, 0))
	fmt.Println(findK(l, 1))
	fmt.Println(findK(l, 2))
	fmt.Println(findK(l, 3))
}

func findK(head *lists, k int) *lists {
	if head == nil || k < 1 {
		return nil
	}
	a, b := head, head
	for i:=0; i < k-1; i++ { // 2
		if a == nil {
			return nil
		}
		a = a.next
	}
	if a == nil {
		return nil
	}
	for a.next != nil {
		a, b = a.next, b.next
	}
	return b
}
