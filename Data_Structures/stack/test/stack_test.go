package test

import (
	st ".."
	"strconv"
	"testing"
)


type testStruct struct {
	str string
	i int
}

func TestStack(t *testing.T) {
	s := &st.Stack{}
	if _, err := s.Pop(); err == nil {
		t.Fatal()
	}

	for i := 0; i < 3; i++ {
		s.Insert(&testStruct{
			str: "ia"+strconv.Itoa(i),
			i:   i,
		})
	}
	for i := 2; i >= 0; i-- {
		v, err := s.Pop()
		if err != nil {
			t.Fatal()
		}
		val := v.(*testStruct)
		if val.i != i {
			t.Fatal()
		}
	}
	if _, err := s.Pop(); err == nil {
		t.Fatal()
	}

}
