package main

import (
	"math/rand"
	"testing"
)

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	m := len(arr) / 2
	//l := make([]int, m)
	//copy(l, arr[:m])
	//r := make([]int, len(arr)-m)
	//copy(r, arr[m:])
	//mergeSort(l)
	//mergeSort(r)
	l := mergeSort(arr[:m])
	r := mergeSort(arr[m:])
	i, il, ir := 0, 0, 0
	for il < len(l) && ir < len(r) {
		if l[il] < r[ir] {
			arr[i] = l[il]
			il++
		} else {
			arr[i] = r[ir]
			ir++
		}
		i++
	}
	for il < len(l) {
		arr[i] = l[il]
		i++
		il++
	}
	for ir < len(r) {
		arr[i] = r[ir]
		i++
		ir++
	}
	return arr
}
//22550215585253
//16537119794285
func genArr() []int {
	arr := make([]int, 50000)
	for i:=0; i<50000;i++{
		arr[i] = rand.Int()
	}
	return arr
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := genArr()
		mergeSort(arr)
	}
}

func TestMerge(t *testing.T) {
	arr := genArr()
	s := mergeSort(arr)
	for i := range s {
		if i > 0 && s[i-1] > s[i] {
			t.Fatal()
		}
	}
}

//func BenchmarkSortN2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		arr := genArr()
//		n2Sort(arr)
//	}
//}
//

//func main() {
//	arr := genArr()
//	//arr := []int{-356, 328, 705, -199, -373, 108, -377, -362, 128, 98, 1, -9, -500, -607, 387, 12, 210, -600, -351, 432}
//	//fmt.Println(arr)
//	mergeSort(arr)
//	//fmt.Println(arr)
//	fmt.Println("start")
//	for i := range arr {
//		if i>1 && arr[i] < arr[i-1] {
//			fmt.Println(arr[i], arr[i-1])
//			fmt.Println("BREAk")
//			break
//		}
//	}
//}

func n2Sort(arr []int) {
	for i := 0; i < len(arr); i++ {

		if i > 0 && arr[i-1] > arr[i] {
			for j := i; j < len(arr); j++ {
				if arr[j-1] > arr[j] {
					arr[j-1], arr[j] = arr[j], arr[j-1]
				}
			}
			i = -1
		}

	}
}
