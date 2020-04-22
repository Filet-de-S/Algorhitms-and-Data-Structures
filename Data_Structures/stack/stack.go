package stack

import "errors"

type node struct {
	data interface{}
	next *node
}

type Stack struct {
	head *node
}

func (s *Stack) Insert(val interface{}) {
	if s.head == nil {
		s.head = &node{data: val}
	} else {
		n := &node{
			data: val,
			next: s.head}
		s.head = n
	}
}

func (s *Stack) Pop() (interface{}, error) {
	if s.head != nil {
		val := s.head.data
		s.head = s.head.next
		return val, nil
	}
	return nil, EmptyStack(errors.New(""))
}

type EmptyStack error