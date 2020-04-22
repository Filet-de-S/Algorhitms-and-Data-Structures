package main

import "fmt"

type lists struct {
	val int
	next *lists
}

func circleNode(head *lists) *lists {
	if head == nil {
		return nil
	}
	same := head
	tmp := head
	for same != nil && same.next != nil {
		same = same.next.next
		tmp = tmp.next
		if tmp == same {
			break
		}
	}
	if same == nil {
		return nil
	}
	tmp = head
	for same != tmp {
		same = same.next
		tmp = tmp.next
	}
	return tmp

}

func main() {
	a := &lists{val:1}
	b := &lists{val:2}
	c := &lists{val:3}
	d := &lists{val:4}
	e := &lists{val:5}

	a.next = b; b.next = c; c.next = d; d.next = e; e.next = b

	fmt.Println(circleNode(a))
}