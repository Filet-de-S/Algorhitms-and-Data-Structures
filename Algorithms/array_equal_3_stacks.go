package main

import "fmt"

type tripleStack struct {
	stack *[]int
	firstNext int
	secNext int
	thirdNext int
}

func newTripleStack() *tripleStack {
	return &tripleStack{
		stack:     &[]int{-1},
		firstNext: 1,
		secNext:   2,
		thirdNext: 3,
	}
}

func (st *tripleStack) insert(key, stackNum int) {
	stN := 0
	switch stackNum {
	case 1: {
		stN = st.firstNext
		st.firstNext += 3
	}
	case 2: {
		stN = st.secNext
		st.secNext += 3
	}
	default: {
		stN = st.thirdNext
		st.thirdNext += 3
	}
	}

	if stN >= len(*st.stack) {
		*st.stack = append(*st.stack, []int{0,0,0}...)
	}
	(*st.stack)[stN] = key
}

func (st *tripleStack) del(stackNum int) {
	stN := 0
	switch stackNum {
	case 1: {
		stN = st.firstNext
		st.firstNext -= 3
	}
	case 2: {
		stN = st.secNext
		st.secNext -= 3
	}
	default: {
		stN = st.thirdNext
		st.thirdNext -= 3
	}
	}
}

func main() {
	arr := newTripleStack()

	arr.insert(2, 3)
	fmt.Println(arr.stack)
	arr.insert(1, 3)
	fmt.Println(arr.stack)
	arr.insert(1, 3)
	fmt.Println(arr.stack)
	arr.insert(123, 2)
	arr.insert(999, 1)
	fmt.Println(arr.stack)
	arr.insert(123, 1)
	arr.insert(999, 2)
	fmt.Println(arr.stack)

}
