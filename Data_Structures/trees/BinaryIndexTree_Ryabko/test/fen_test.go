package testing

import (
	Ryabko ".."
	"math/rand"
	"testing"
	"time"
)

const arraySize = 100000

func TestSumFromZero(t *testing.T) {
	arr, prefixSumArr := initArrPrefixSum()
	fen := Ryabko.Init(arr)

	l, r := 0, arraySize
	fSum := fen.RangeSumQuery(l, r)
	naive := prefixSum(prefixSumArr, l, r)
	if fSum != naive {
		t.Fatal()
	}

	for i := 0; i < 10000; i++ {
		l, r = 0, rand.Intn(arraySize)
		fSum = fen.RangeSumQuery(l, r)
		naive = prefixSum(prefixSumArr, l, r)
		if fSum != naive {
			t.Fatal()
		}
	}
}

func TestSumFromToRand(t *testing.T) {
	arr, prefixSumArr := initArrPrefixSum()
	fen := Ryabko.Init(arr)

	for i := 0; i < 100000; i++ {
		l, r := rand.Intn(1000), arraySize

		fSum := fen.RangeSumQuery(l, r)
		naive := prefixSum(prefixSumArr, l, r)

		if fSum != naive {
			t.Fatal()
		}
	}
}

func TestAddThenSum(t *testing.T) {
	arr, prefixSumArr := initArrPrefixSum()
	fen := Ryabko.Init(arr)

	for i := 0; i < 10000; i++ {
		l, r := rand.Intn(1000), arraySize
		pos, val := rand.Intn(1000), rand.Intn(1000000)
		if l > pos {
			l, pos = pos, l
		}

		fen.Add(pos, val)
		fSum := fen.RangeSumQuery(l, r)

		prefixAdd(prefixSumArr, pos, val)
		naive := prefixSum(prefixSumArr, l, r)

		if fSum != naive {
			t.Fatal()
		}
	}
}

func TestUpdateZeroThenSum(t *testing.T) {
	arr, prefixSumArr := initArrPrefixSum()
	fen := Ryabko.Init(arr)

	for i := 0; i < 10000; i++ {
		l, r := 0, arraySize
		pos, val := 0, rand.Intn(1000000)
		fen.Update(pos, val)
		fSum := fen.RangeSumQuery(l, r)

		prefixUpdate(prefixSumArr, pos, val)
		naive := prefixSum(prefixSumArr, l, r)

		if fSum != naive {
			t.Fatal()
		}
	}
}

func TestUpdateRandThenSum(t *testing.T) {
	arr, prefixSumArr := initArrPrefixSum()
	fen := Ryabko.Init(arr)

	for i := 0; i < 10000; i++ {
		l, r := rand.Intn(1000), arraySize
		pos, val := rand.Intn(1000), rand.Intn(1000000)
		if l > pos {
			l, pos = pos, l
		}

		fen.Update(pos, val)
		fSum := fen.RangeSumQuery(l, r)

		prefixUpdate(prefixSumArr, pos, val)
		naive := prefixSum(prefixSumArr, l, r)

		if fSum != naive {
			t.Fatal(fSum-naive)
		}
	}
}

func prefixUpdate(arr []int, pos int, val int) {
	if pos > 0 {
		// 6 2 = 4
		oldValue := arr[pos] - arr[pos-1]
		// 4 - 1 = 3 = -3
		diff := -(oldValue - val)
		prefixAdd(arr, pos, diff)
	} else {
		old := arr[pos]
		diff := -(old - val)
		prefixAdd(arr, pos, diff)
	}
}

func prefixSum(arr []int, l, r int) (sum int) {
	if l > 0 {
		return arr[r] - arr[l-1]
	} else {
		return arr[r]
	}
}

func prefixAdd(arr []int, pos, val int) {
	for i := pos; i < len(arr); i++ {
		arr[i] += val
	}
}

func initArrPrefixSum() (arr, prefixSumArr []int) {
	arr = make([]int, arraySize+1)
	prefixSumArr = make([]int, arraySize+1)
	rand.Seed(time.Now().UnixNano())
	arr[0] = rand.Intn(100000)
	prefixSumArr[0] = arr[0]
	for i := 1; i < arraySize+1; i++ {
		arr[i] = rand.Intn(100000)
		prefixSumArr[i] = arr[i] + prefixSumArr[i-1]
	}
	return arr, prefixSumArr
}
