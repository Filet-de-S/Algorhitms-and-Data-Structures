package main

import "fmt"

func fibo_for(n int) int {
	prev, next := 0, 1
	for i := 0; i < n; i++ {
		temp := next
		next = prev + next
		prev = temp
	}
	return prev
}

func main() {
	fmt.Println(fibo_for(10))
}
