package main

import (
	"fmt"
	"strconv"
)

type channels struct {
	in, out chan int
	res     chan *string
}

type step func(ch channels)

var leftmost = make(chan int)

func infiniteSeq(data int, res chan *string, funcs ...step) {
	go func() {
		leftmost <- data
	}()
	ch := channels{leftmost, make(chan int), res}
	for i := 0; i < len(funcs); i++ {
		// add wg if wanna after {receive loop} function work
		ch.out = make(chan int)
		go func(fstep step, ch channels) {
			fstep(ch)
			close(ch.out)
		}(funcs[i], ch)
		ch.in = ch.out
	}
}

func main() {
	funcs := []step{
		func(ch channels) {
			for i := range ch.in {
				fmt.Println("LEFT", i)
				ch.out <- i + 1
			}
			fmt.Println("LEFT EXIT")
		},
		func(ch channels) {
			for i := range ch.in {
				fmt.Println("RIGHT", i)
				ch.out <- i + 1
			}
			fmt.Println("RIGHT EXIT")
		},
		func(ch channels) {
			for i := range ch.in {
				fmt.Println("RESULT", i)
				if i > 10 {
					res := strconv.Itoa(i)
					ch.res <- &res
					close(leftmost)
					break
				}
				leftmost <- i + 1
			}
			fmt.Println("RESULT EXIT")
		},
	}
	res := make(chan *string)
	data := 5
	infiniteSeq(data, res, funcs...)
	fmt.Println("\nRES MAIN", *<-res)
}
