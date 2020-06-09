package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
// https://contest.yandex.ru/contest/17160/problems/E/
type tree struct {
	val         float64
	left, right *tree
}

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	if r.Err() != nil {
		os.Exit(0)
	}

	bacteries, err := strconv.Atoi(r.Text())
	if err != nil || bacteries < 1 || bacteries > 1000 {
		os.Exit(0)
	}

	m := map[float64]*tree{}
	var root *tree
	for r.Scan() {
		txt := r.Text()
		if len(txt) == 0 {
			break
		}
		spl := strings.Split(txt, " ")
		if len(spl) != 4 {
			os.Exit(0)
		}
		bact := 0.0
		var t *tree
		for i := range spl {
			n, err := strconv.Atoi(spl[i])
			if err != nil || n < -1000 || n > 1000 {
				os.Exit(0)
			}
			if i == 0 {
				bact = float64(n)
			} else if i == 1 {
				t = &tree{val: float64(n)}
			} else if n != -1 {
				if i == 2 {
					t.left = &tree{val: float64(n)}
				} else {
					t.right = &tree{val: float64(n)}
				}
			}
		}
		if t.left != nil {
			m[t.left.val] = t
		}
		if t.right != nil {
			m[t.right.val] = t
		}

		if root == nil {
			root = t
		} else {
			tr := m[bact]
			if tr.left != nil && tr.left.val == bact {
				tr.left = t
			} else {
				tr.right = t
			}
		}
	}
	m = nil

	res := make([][]float64, 0)
	que := []*tree{root}
	for i := 0;; i++ {
		l := len(que)
		if l == 0 {
			break
		}
		if len(res) < i+1 {
			res = append(res, []float64{})
		}
		for j := range que {
			res[i] = append(res[i], que[j].val)
			if que[j].left != nil {
				que = append(que, que[j].left)
			}
			if que[j].right != nil {
				que = append(que, que[j].right)
			}
		}
		que = que[l:]
	}
	que = nil
	root = nil
	print := []byte{}
	for i := range res {
		var av float64 = 0.0
		for j := range res[i] {
			av += res[i][j]
		}
		print = append(print, []byte(fmt.Sprintf("%.2f", av/float64(len(res[i])))+" ")...)
	}
	os.Stdout.Write(print)
}
