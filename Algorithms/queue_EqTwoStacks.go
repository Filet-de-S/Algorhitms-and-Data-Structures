package main

import "fmt"

type stack struct { // last in, first out
	st []int
}

type myQueue struct { // first in, first out
	 one stack
	 two stack
}

func (s *stack) insert(key int) {
	s.st = append(s.st, key)
}

func (s *stack) pop() (r int) {
	r = s.st[len(s.st)-1]
	s.st = s.st[:len(s.st)-1]
	return r
}

func (q *myQueue) Insert(key int) {
	q.one.insert(key)
}

func (q *myQueue) Pop() int {
	if len(q.two.st) == 0 {
		for len(q.one.st) != 1 {
			q.two.insert( q.one.pop() )
		}
		return q.one.pop()
	}
	return q.two.pop()
}

func newMyQueue() *myQueue {
	return &myQueue{
		one: stack{},
		two: stack{},
	}
}

func main() {
	q := &myQueue{}
	for i := 0; i < 20; i++ {
		q.Insert(i)
		fmt.Println("inserted", i)
		if i == 5 {
			fmt.Println("popped", q.Pop())
			fmt.Println("popped", q.Pop())
		}
		if i == 10 {
			fmt.Println("popped", q.Pop())
			fmt.Println("popped", q.Pop())
		}
	}
}