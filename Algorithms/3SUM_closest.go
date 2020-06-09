//Гарри Поттер и Гермиона играют в игру. У Гермионы есть фишки с указанным на них количеством очков.
//	На каждой фишке — k очков, k находится в промежутке от −10^5 до 10^5.
//	Гарри называет число, а Гермиона должна выбрать три фишки, сумма очков на которых наиболее близка с заданному числу.
//	Долго же будет маленькая волшебница искать нужные фишки! Помогите ей справиться с задачей скорее.
//
//Формат ввода
//	Первая строка содержит целое число n в диапазоне от −10^5 до 10^5, число, названное первым участником.
//	Во второй строке через пробел заданы целые числа в диапазоне от −10^5 до 10^5, — очки для фишек второго участника.
//	Их количество может быть от 3 до 10^4.
//
//Формат вывода
//Нужно вывести целое число — сумму очков трёх фишек, наиболее близкую к n.
//
//Пример 1
//Ввод	Вывод
//6		1
//-1 -1 -9 -7 3 -6
//
//Пример 2
//Ввод	Вывод
//5		6
//7 -8 2 -8 -3
//
//Пример 3
//Ввод	Вывод
//8		8
//6 2 8 -3 1 1 6 10

// Решение должно работать не дольше, чем за O(n^2). Ограничение по памяти — O(1).
// copyright to https://contest.yandex.ru/contest/17160/problems/F/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	lowBoundNum     = int(-math.Pow10(5))
	upperBoundNum   = int(math.Pow10(5))
	lowBoundChips   = 3
	upperBoundChips = int(math.Pow10(4))
	knapsackWeight  = 3
	harrysNum       = 0
	err             error
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	if r.Err() != nil {
		panic(r.Err())
	}

	txt := r.Text()
	harrysNum, err = strconv.Atoi(txt)
	if err != nil || harrysNum < lowBoundNum || harrysNum > upperBoundNum {
		panic("Harry! Thinking in Snake is okay.. but why not in numbers between −10^5 and 10^5?")
	}

	r.Scan()
	if r.Err() != nil {
		panic(r.Err())
	}

	chipsSliced := strings.Split(r.Text(), " ")
	if len(chipsSliced) > upperBoundChips || len(chipsSliced) < lowBoundChips {
		panic("Ron! Hand out chips beetwen 3 and 10^4, please")
	}

	chips := make([]int, len(chipsSliced))
	for i := range chipsSliced {
		num, err := strconv.Atoi(chipsSliced[i])
		if err != nil || num < lowBoundNum || num > upperBoundNum {
			panic("Hermione! I trust you can conjure not only numbers on chips or go outside the bounds...\nWhat about playing by rules? " +
				"Keep your chips between −10^5 and 10^5!")
		}
		chips[i] = num
	}
	sort.Ints(chips)
	os.Stdout.Write(closest(chips)) // ok
	// []byte(strconv.Itoa(knapsack(knapsackWeight, len(chips), &chips))))
}

func closest(chips []int) []byte {
	l := len(chips) - 1
	closestSum := chips[0] + chips[1] + chips[l]

	for i := 0; i < l-1; i++ {
		start, end := i+1, l
		for start < end {
			curSum := chips[i] + chips[start] + chips[end]
			if abs(harrysNum-curSum) < abs(harrysNum-closestSum) {
				closestSum = curSum
			}
			if curSum > harrysNum {
				end--
			} else {
				start++
			}
		}
	}
	return []byte(strconv.Itoa(closestSum))
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func knapsack(W, n int, chips *[]int) int { // wrong, cause  we need all combinations, not sequentially
	W--; n--

	lh := (*chips)[n]
	if W > 0 && n > 0 {
		lh += knapsack(W, n, chips)
	}
	if n > 0 {
		rh := knapsack(W+1, n, chips)
		return nearest(lh, rh)
	}
	return lh
	//return nearest((*chips)[n-1]+knapsack(W-1, chips, n-1), knapsack(W, chips, n-1))
}

func nearest(lh, rh int) int {
	fmt.Println("NEAR", lh, rh)
	lhs := abs(lh - harrysNum)
	rhs := abs(rh - harrysNum)
	if lhs < rhs {
		return lh
	}
	return rh
}
