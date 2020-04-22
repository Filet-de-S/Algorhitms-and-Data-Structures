package main

import "fmt"

func printNMmtx(mtx [][]int) {
	l := len(mtx)
	for i:=0; i<l; i++ {
		ll:= len(mtx[i])
		for j:=0;j<ll;j++ {
			fmt.Print(mtx[i][j], " ")
			if mtx[i][j] < 10 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func fillByZero(mtx [][]int) {
	y := len(mtx)
	if y < 1 {
		return
	}
	x := len(mtx[0])
	zeros := [][]int{}

	for i := 0; i<y; i++ {
		for j:=0; j<x; j++ {
			if mtx[i][j] == 0 {
				zeros = append(zeros, []int{i,j})
			}
		}
	}

	for i := range zeros {
		y1 := zeros[i][0]
		x1 := zeros[i][1]

		// po x
		for i:=0; i < x; i++ {
			mtx[y1][i] = 0
		}
		// po y
		for i:=0; i < y; i++ {
			mtx[i][x1] = 0
		}
	}
}

func main() {
	mtx := make([][]int, 6)

	k := 1
	for i := 0; i < 6;i++ { // make mtx
		mtx[i] = make([]int, 4)
		for j:=0;j<4;j++ {
			mtx[i][j] = k
			k++
		}
	}
	mtx[4][2] = 0
	mtx[1][1] = 0
	printNMmtx(mtx)
	fillByZero(mtx)
	fmt.Println()
	fmt.Println()
	printNMmtx(mtx)
}

