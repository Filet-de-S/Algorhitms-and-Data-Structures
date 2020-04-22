package main

import (
	"fmt"
)

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	f := n
	for n > 2 {
		n--
		f *= n
	}
	return f
}

func main ()  {
	fmt.Println(factorial(5))
	fmt.Println(factorial( 19))
	fmt.Println(factorial(3))
	fmt.Println(factorial(2))
	fmt.Println(factorial(0))
}
