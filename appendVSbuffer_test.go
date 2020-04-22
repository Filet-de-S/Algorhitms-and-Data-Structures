package main

import (
	"bytes"
	"testing"
)

var str = "STROKAASKDJASDJASTROKAASKDJASDJASTROKAASKDJASDJASTROKAASKDJASDJASTROKAASKDJASDJASTROKAASKDJASDJASTROKAASKDJASDJA"

func BenchmarkAppend(b *testing.B) {
	sl := []byte{}

	for i := 0; i<10000000;i++ {
		sl = append(sl, []byte(str)...)
	}
	s := string(sl)
	if s != "" {
	}
}
func BenchmarkBuffer(b *testing.B) {

	buf := bytes.Buffer{}
	for i := 0; i<10000000;i++ {
		buf.WriteString(str)
	}
	sl := buf.String()
	if sl != "" {
	}
}