package main

import "fmt"

type somest struct {
	ne *somest
}

func f() bool {
	return false || false
}

func main() {
	n := 50
	for n > 0 {
		fmt.Println(n % 2)
		n /= 2
	}
}