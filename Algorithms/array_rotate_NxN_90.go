package main

import "fmt"

func PrintMtx(mtx [][]int) {
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

func rotateByNintyRight(mtx [][]int) {
	n := len(mtx) - 1

	for i := 0; i < n; i++ {
		nn := n
		j := i
		for ; j < n; j++ {
			temp := mtx[i][j] //top
			mtx[i][j] = mtx[nn][i] //top from left
			mtx[nn][i] = mtx[n][nn] // left from bottom
			mtx[n][nn] = mtx[j][n] // bottom from right
			mtx[j][n] = temp // right from top
			nn--
		}
		n--
	}
}

func main() {
	mtx := make([][]int, 9)

	k := 0
	for i := 0; i < 9;i++ { // make mtx
		mtx[i] = make([]int, 9)
		for j:=0;j<9;j++ {
			mtx[i][j] = k
			k++
		}
	}

	PrintMtx(mtx)
	fmt.Println()
	fmt.Println()

	rotateByNintyRight(mtx)
	PrintMtx(mtx)

}
