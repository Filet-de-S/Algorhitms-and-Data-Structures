package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func compress(s string) string {
	var last rune
	str := []rune(s)

	c := 0
	res := bytes.Buffer{}
	for i := range str {
		if str[i] == last {
			c++
			if i == len(str) - 1 {
				res.WriteString(string(last))
				res.WriteString(strconv.Itoa(c))
			}
		} else if last != 0 {
			res.WriteString(string(last) + strconv.Itoa(c))
			last = str[i]
			c = 1
		} else {
			last = str[i]
			c++
		}
	}

	if res.Len() < len(s) {
		return res.String()
	}
	return s
}

func main() {
	vars := []string{"aabcccccaaa","abc","aabbccc", ""}

	for i:=range vars {
		fmt.Println(compress(vars[i]))
	}
}
