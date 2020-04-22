package main

import (
	"errors"
	"fmt"
)

const maxStackSize = 2

type stack []int

type setOfStacks struct {
	set []stack
}

func (s *setOfStacks) insert(key int) {
	last := len(s.set) - 1
	if last < 0 {
		s.set = make([]stack, 1)
		s.set[0] = make(stack, 0, maxStackSize)
		s.set[0] = append(s.set[0], key)
	} else {
		if len(s.set[last]) == maxStackSize {
			s.set = append(s.set, make(stack, 0, maxStackSize))
			s.set[last+1] = append(s.set[last+1], key)
		} else {
			s.set[last] = append(s.set[last], key)
		}
	}
}

func (s *setOfStacks) pop() (int, error) {
	last := len(s.set) - 1
	if last < 0 {
		return 0, errors.New("empty")
	}
	l := len(s.set[last])-1
	ret := s.set[last][l]
	if l == 0 {
		s.set = s.set[:last]
	} else {
		s.set[last] = s.set[last][:l]
	}

	return ret, nil
}

func main()  {
	s := &setOfStacks{}
	_, err := s.pop()
	fmt.Println(err)

	for i := 0; i < 20; i++ {
		s.insert(i)
		fmt.Println("inserted", i)
		if i == 5 {
			_,_ = s.pop()
			fmt.Println("pop", s.set)
			_,_ = s.pop()
			fmt.Println("pop", s.set)
		}
		if i == 6 {
			fmt.Println(s.set)
		}
	}
}
