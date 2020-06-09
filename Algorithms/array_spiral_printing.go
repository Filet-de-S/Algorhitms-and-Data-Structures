package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)


func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	if r.Err() != nil {
		os.Exit(0)
	}

	mtxSize, err := strconv.Atoi(r.Text()); if err != nil || mtxSize <= 0 ||
		mtxSize > 1000 || mtxSize % 2 == 0 {
		os.Exit(0)
	}

	mtx := make([][]int, 0, mtxSize)
	for r.Scan() {
		txt := r.Text()
		if len(txt) == 0 {
			break
		}
		spl := strings.Split(txt, " ")
		if len(spl) != mtxSize {
			os.Exit(0)
		}

		nums := make([]int, mtxSize)
		for i := range spl {
			n, err := strconv.Atoi(spl[i])
			if err != nil || n < -1000 || n > 1000 {
				os.Exit(0)
			}
			nums[i] = n
		}
		mtx = append(mtx, nums)
	}
	if len(mtx) == 0 {
		os.Exit(0)
	}
	//5
	//4 10 7 10 9
	//5  9 0  9 8
	//8  3 6  0 2
	//8 10 3  0 0
	//0  9 0  7 4
	// 2,2
	//1,1 => 1,2 => 1,3 => 2,3 => 3,3 => 3,2 => 3,1 => 2,1 =>
	//0,1 => 0,2 => 0,3 => 0,4 => 1,4 => 2,4 => 3,4 => 4,4
	//4,3 => 4,2 => 4,1 => 4,0 => 3,0 => 2,0 => 1,0 => 0,0
	//if !up => right
	//if !right => down
	//if !down => left
	//if !left => up
	//if mtxSize > 50 {
	//	mtxSize = 50
	//}
	print := make([]byte, 0, 100)
	center := mtxSize/2
	y, x := center, center
	print = append(print, []byte(strconv.Itoa(mtx[y][x])+"\n")...)
	coef := 0
	for ; y != 0 && x != 0; {
		coef++
		//if coef > 5 {
		os.Stdout.Write(print)
		print = make([]byte, 0, 100)
		//}
		max := center+coef
		min := center-coef
		for {
			for y-1 >= min && y-1 >= 0 {
				y--
				//os.Stdout.Write([]byte(strconv.Itoa(mtx[y][x])+"\n"))
				print = append(print, []byte(strconv.Itoa(mtx[y][x])+"\n")...)
			}

			if x == min && y == min {
				break
			}

			for x+1 <= max && x+1 < mtxSize {
				x++
				//os.Stdout.Write([]byte(strconv.Itoa(mtx[y][x])+"\n"))
				print = append(print, []byte(strconv.Itoa(mtx[y][x])+"\n")...)
			}

			for y+1 <= max && y+1 < mtxSize {
				y++
				//os.Stdout.Write([]byte(strconv.Itoa(mtx[y][x])+"\n"))
				print = append(print, []byte(strconv.Itoa(mtx[y][x])+"\n")...)
			}

			for x-1 >= min && x-1 >= 0 {
				x--
				//os.Stdout.Write([]byte(strconv.Itoa(mtx[y][x])+"\n"))
				print = append(print, []byte(strconv.Itoa(mtx[y][x])+"\n")...)
			}
		}
	}
	os.Stdout.Write(print)
}