package main

import (
	"fmt"
)

func main() {

	str := []string{ "abc"}

	for i := range str {
		fmt.Println("case", str[i], "\nresult", permuteChars(str[i]))
	}
}

func permuteChars(str string) (res []string) {
	if str == "" {
		return []string{""}
	}
	first := string(str[0])
	rem := str[1:]
	words := permuteChars(rem)
	for i := range words {
		for j := 0; j <= len(words[i]); j++ {
			fmt.Println(words[i])
			s := words[i][0:j] + first + words[i][j:] // abc
			res = append(res, s)
		}
	}
	return res
}
