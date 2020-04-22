package main

import (
	"fmt"
)

func genMap(lhs string) map[rune]int {
	lha := map[rune]int{}
	for i:=0; i< len(lhs);i++ {
		lha[rune(lhs[i])]++
	}
	return lha
}

func isAnagram(lhs string, rhs string) bool {
	if len(lhs) != len (rhs) {
		return false
	}
	lha := genMap(lhs)
	rha := genMap(rhs)

	for i := range lha {
	fmt.Println("AAAA!", string(i))
		if lha[i] != rha[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("йухх", "хуйх"))
}
