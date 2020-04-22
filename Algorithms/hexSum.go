package main

import (
	"fmt"
	"math"
)

func main() {
	s := sumHexDummy("ff", "10a") // ff = 15 * 16pow1 + 15 * 16pow0 = 255; + (1 * 16pow2, 0*16pow1, +10 = 266)
	fmt.Println(s)
	//s = sumHexClever("ff", "10a")
	//fmt.Println(s)
}
//
//func sumHexClever(s1 string, s2 string) string {
//	if s1 == "0" {
//		return s2
//	} else if s2 == "0" {
//		return s1
//	}
//
//	n1 := &s1
//	n2 := &s2
//	if len(s2) > len(s1) {
//		n1 = &s2
//		n2 = &s1
//	}
//
//	res := []byte{}
//	rmb := 0 // if digit doesn't suit in limit, we remember +1 to next digit
//	for i := len(*n1) - 1; i >= 0; i-- {
//		for j := len(*n2) - 1; j >= 0; j-- {
//			if (*n2)[j] > '9' && // to much options: 'f' >> +1 rmb++ || 'a' >> +5 rmb++ ...
//
//			res = append([]byte{}, res...)
//		}
//	}
//
//
//	return "ad"
//}



func getHex(s string) (n int64) {
	j := 0.0
	for i := len(s)-1; i >= 0; i-- {
		if s[i] > '9' {
			n += int64(s[i] - 'a' + 10) * int64(math.Pow(16.0, j))
		} else {
			n += int64(s[i] - '0') * int64(math.Pow(16, j))
		}
		j++
	}
	return n
}

func sumHexDummy(s1 string, s2 string) string {
	n1 := getHex(s1)
	n2 := getHex(s2)

	return toHexString(n1+n2)
}

func toHexString(n int64) string {
	if n == 0 {
		return "0"
	}
	nDup := n
	i := 0
	for ; nDup != 0; i++ {
		nDup /= 10
	}
	neg := 0
	if n < 0 {
		neg = 1
		i++
	}

	s := make([]byte, i)
	i--
	for ; i > 0; i-- {
		if n % 16 > 9 {
			s[i] = byte(n % 16) + 'a' - 10
		} else {
			s[i] = byte(n % 16) + '0'
		}
		n /= 16
	}

	if neg == 1 {
		s[0] = '-'
	} else {
		if n % 16 > 9 {
			s[i] = byte(n % 16) + 'a' - 10
		} else {
			s[i] = byte(n % 16) + '0'
		}
	}
	return string(s)
}
