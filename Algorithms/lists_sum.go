package main

import "fmt"

type lists struct {
	val int
	next *lists
}

func main() {
	a := &lists{
		8, &lists{
			val:  1,
			next: nil,
	}}
	b := &lists{
		9, &lists{
			val:  9,
			next: &lists{
				val:  9,
				next: nil,
			},
		}}
	s := list_sum(a,b)
	for s != nil {
		fmt.Println(s.val)
		s = s.next
	}
}

func list_sum(a *lists, b *lists) *lists {
	rem := false
	res := []int{}
	for a != nil && b != nil {
		r := a.val + b.val
		if rem == true {
			r++
			rem = false
		}
		if r > 9 {
			rem = true
			r %= 10
		}
		res = append([]int{r}, res...)
		a, b = a.next, b.next
	}
	for a != nil {
		r := a.val
		if rem == true {
			r++
			rem = false
		}
		if r > 9 {
			rem = true
			r %= 10
		}
		res = append([]int{r}, res...)
		a = a.next
	}
	for b != nil {
		r := b.val
		if rem == true {
			r++
			rem = false
		}
		if r > 9 {
			rem = true
			r %= 10
		}
		res = append([]int{r}, res...)
		b = b.next
	}
	if rem == true {
		res = append([]int{1}, res...)
	}

	fmt.Println(res)
	s := &lists{}
	head := s
	s.val = res[len(res)-1]
	for i := len(res)-2; i >= 0; i-- {
		s.next = &lists{val:res[i]}
		s = s.next
	}
	return head
}
