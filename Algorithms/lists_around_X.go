package main

import "fmt"

type lists struct {
	val int
	next *lists
}

func main () {
	l := &lists{70, &lists{
		val:  1,
		next: &lists{
			val:  50,
			next: &lists{
				val:  2,
				next: &lists{
					val:  5,
					next: &lists{
						val:  900,
						next: nil,
					},
				},
			},
		},
	}}
	l = aroundX(l, 5)
	for l != nil {
		fmt.Println(l.val)
		l = l.next
	}
}

func aroundX(l *lists, x int) *lists {
	if l == nil {
		return nil
	}
	before, after := &lists{}, &lists{}
	var befHead *lists
	var afterHead *lists

	for l != nil {
		if l.val < x {
			if befHead == nil {
				before = &lists{val: l.val}
				befHead = before
			} else {
				before.next = &lists{val: l.val}
				before = before.next
			}
		} else {
			if afterHead == nil {
				after = &lists{val: l.val}
				afterHead = after
			} else {
				after.next = &lists{val: l.val}
				after = after.next
			}
		}
		l = l.next
	}
	if afterHead != nil {
		before.next = &lists{
			val:  afterHead.val,
			next: afterHead.next,
		}
	}
	return befHead
}

