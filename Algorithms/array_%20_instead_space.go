package main

import "fmt"

//func run(s []byte, l int) {
//	sp := 0
//	for i:=0;i<l;i++{
//		if s[i] == ' ' {
//			sp++
//		}
//	}
//
//	in := l + sp*2
//
//	for i := 13 - 1; i >= 0; i-- {
//		if s[i] == ' ' {
//			s[in-1] = '0'
//			s[in-2] = '2'
//			s[in-3] = '%'
//			in -= 3
//		} else {
//			s[in - 1] = s[i]
//			in--
//		}
//	}
//}

func run(arr []rune) {
	truelen := 0
	for i := 0; i < len(arr); i++ {
		if i>0 && arr[i-1] == ' ' && arr[i] != ' '{
			truelen += 2
		} else if arr[i] != ' ' {
			truelen++
		}
	}
	sp := len(arr) - 1
	for i:=truelen-1; i >= 0;i-- {
		if arr[i] == ' ' {
			arr[sp] = '0'
			arr[sp-1] = '2'
			arr[sp-2] = '%'
			sp -= 3
		} else {
			arr[sp] = arr[i]
			sp--
		}
	}
}


func main() {
	s := []rune("Mв John ФЫВыы ш s        ")
	run(s)
	fmt.Println("***"+string(s) + "***")
}
