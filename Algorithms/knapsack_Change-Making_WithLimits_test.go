package main

import (
	"errors"
	"math"
	"strconv"
	"testing"
)

const ATMlimit = 100000

const (
	_ = iota
	fiveThousands
	thousands
	fiveHundreds
	hundreds
	fifties
	thirties
)

type ATMwallet struct {
	fiveThousands, thousands int
	fiveHundreds, hundreds   int
	fifties, thirties        int
	totalMoneyInATM          int
	notesAmount              int
}
type withdrawMoney struct {
	fiveThousands, thousands int
	fiveHundreds, hundreds   int
	fifties, thirties        int
	amount                   int
}

type nomValue struct {
	nom        int
	amountLeft int
}


func TestRecursive(t *testing.T) {
	atm := encash()
	m := atm.giveMeMoneyRecursive(120)
	if m == nil || !(m.thirties == 4 && m.fifties == 0 && m.hundreds == 0 && m.fiveHundreds == 0 && m.thousands == 0 && m.fiveThousands == 0) {
		t.Fatal(120, m)
	} else if atm.thirties != (encash().thirties-4) || atm.fiveThousands != (encash().fiveThousands) {
		t.Log("Initially", encash())
		t.Fatal(120, "ATM", atm)
	}

	atm = encash()
	m = atm.giveMeMoneyRecursive(1190)
	if m == nil || !(m.thirties == 3 && m.fifties == 0 && m.hundreds == 1 && m.fiveHundreds == 0 && m.thousands == 1 && m.fiveThousands == 0) {
		t.Fatal(1190, m)
	} else if atm.fiveThousands != (encash().fiveThousands) || atm.thousands != (encash().thousands-1) || atm.fiveHundreds != (encash().fiveHundreds) {
		t.Fatal(1190, atm)
	}

	atm = encash()
	m = atm.giveMeMoneyRecursive(6620)
	if m == nil || !(m.thirties == 4 && m.fifties == 0 && m.hundreds == 0 && m.fiveHundreds == 1 && m.thousands == 1 && m.fiveThousands == 1) {
		t.Fatal(6620, m)
	}

	atm = encash()
	m = atm.giveMeMoneyRecursive(32160)
	if m == nil || !(m.thirties == 22 && m.fifties == 100 && m.hundreds == 5 && m.fiveHundreds == 2 && m.thousands == 5 && m.fiveThousands == 4) {
		t.Fatal(32160, m)
	}

	atm = encash()
	m = atm.giveMeMoneyRecursive(32190)
	if m == nil || !(m.thirties == 23 && m.fifties == 100 && m.hundreds == 5 && m.fiveHundreds == 2 && m.thousands == 5 && m.fiveThousands == 4) {
		t.Fatal(32190, m)
	} else if !(atm.thirties == 0 && atm.fifties == 0 && atm.hundreds == 0 && atm.fiveHundreds == 0 && atm.thousands == 0 && atm.fiveThousands == 0) {
		t.Fatal(32190, atm)
	}

	atm = encash()
	m = atm.giveMeMoneyRecursive(270)
	if m == nil || !(m.thirties == 4 && m.fifties == 1 && m.hundreds == 1 && m.fiveHundreds == 0 && m.thousands == 0 && m.fiveThousands == 0) {
		t.Fatal(270, m)
	} else if !(atm.thirties == 19 && atm.fifties == 99 && atm.hundreds == 4 && atm.fiveHundreds == 2 && atm.thousands == 5 && atm.fiveThousands == 4) {
		t.Fatal(270, atm)
	}

	atm = encash()
	m = atm.giveMeMoneyRecursive(100000)
	if m != nil {
		t.Fatal("overdraft 100000", m)
	}
}
func BenchmarkRecursiveBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atm := encash()
		atm.fiveThousands = 0
		atm.fiveHundreds += 0 //100 // 50000
		atm.hundreds += 4     //100  //+ 10000
		atm.fifties += 100
		atm.thirties += 100 //+ 3000
		//atm.totalMoneyInATM += 63000 //75190
		m := atm.giveMeMoneyRecursive(20560)
		if m == nil {
		}
	}
}

func (atm *ATMwallet) giveMeMoneyRecursive(sum int) *withdrawMoney {
	wm, err := atm.helperRec(withdrawMoney{}, 1, 0, sum)
	if err != nil {
		return nil
	}
	return &wm
}
func (atm *ATMwallet) helperRec(withdraw withdrawMoney, nom, notes, sum int) (withdrawMoney, error) {
	if nom > 6 {
		return withdrawMoney{}, errors.New("cant")
	}

	i := 0
	for sum > 0 && atm.isNomFitToSum(nom, sum) {
		value := atm.coinTo(&withdraw, nom)
		sum -= value
		notes++
		i++
	}
	if sum == 0 {
		return withdraw, nil
	}

	newWithdraw, err := atm.helperRec(withdraw, nom+1, notes, sum)
	if err == nil {
		return newWithdraw, nil
	}

	for i > 0 {
		sum += atm.moneyBackV2(nom, &withdraw, 1)
		notes -= 1
		i--

		newWithdraw, err = atm.helperRec(withdraw, nom+1, notes, sum)
		if err == nil {
			return newWithdraw, nil
		}
	}
	return withdrawMoney{}, errors.New("cantEND")
}

func (atm *ATMwallet) isNomFitToSum(nom int, sum int) bool {
	nomVal := new(nomValue)

	switch nom {
	case fiveThousands:
		nomVal.amountLeft = atm.fiveThousands
		nomVal.nom = 5000
	case thousands:
		nomVal.amountLeft = atm.thousands
		nomVal.nom = 1000
	case fiveHundreds:
		nomVal.amountLeft = atm.fiveHundreds
		nomVal.nom = 500
	case hundreds:
		nomVal.amountLeft = atm.hundreds
		nomVal.nom = 100
	case fifties:
		nomVal.amountLeft = atm.fifties
		nomVal.nom = 50
	case thirties:
		nomVal.amountLeft = atm.thirties
		nomVal.nom = 30
	default:
		panic("WHAT")
	}

	if nomVal.amountLeft == 0 || nomVal.nom > sum {
		return false
	}
	return true
}
func (atm *ATMwallet) coinTo(withdraw *withdrawMoney, nom int) int {
	switch nom {
	case fiveThousands:
		atm.fiveThousands--
		withdraw.fiveThousands++
		return 5000
	case thousands:
		atm.thousands--
		withdraw.thousands++
		return 1000
	case fiveHundreds:
		atm.fiveHundreds--
		withdraw.fiveHundreds++
		return 500
	case hundreds:
		atm.hundreds--
		withdraw.hundreds++
		return 100
	case fifties:
		atm.fifties--
		withdraw.fifties++
		return 50
	case thirties:
		atm.thirties--
		withdraw.thirties++
		return 30
	default:
		panic("WHAT")
	}
}
func (atm *ATMwallet) moneyBack(nom int, money *withdrawMoney) int {
	switch nom {
	case fiveThousands:
		atm.fiveThousands++
		money.fiveThousands--
		return 5000
	case thousands:
		atm.thousands++
		money.thousands--
		return 1000
	case fiveHundreds:
		atm.fiveHundreds++
		money.fiveHundreds--
		return 500
	case hundreds:
		atm.hundreds++
		money.hundreds--
		return 100
	case fifties:
		atm.fifties++
		money.fifties--
		return 50
	case thirties:
		atm.thirties++
		money.thirties--
		return 30
	default:
		panic("WHAT")
	}
}
func (atm *ATMwallet) moneyBackV2(nom int, money *withdrawMoney, times int) int {
	switch nom {
	case fiveThousands:
		atm.fiveThousands += times
		money.fiveThousands -= times
		return 5000 * times
	case thousands:
		atm.thousands += times
		money.thousands -= times
		return 1000 * times
	case fiveHundreds:
		atm.fiveHundreds += times
		money.fiveHundreds -= times
		return 500 * times
	case hundreds:
		atm.hundreds += times
		money.hundreds -= times
		return 100 * times
	case fifties:
		atm.fifties += times
		money.fifties -= times
		return 50 * times
	case thirties:
		atm.thirties += times
		money.thirties -= times
		return 30 * times
	default:
		panic("WHAT")
	}
}


func TestDynamic(t *testing.T) {
	atm := encash()
	m := atm.giveMeMoneyDynamic(120)
	if m == nil || !(m.thirties == 4 && m.fifties == 0 && m.hundreds == 0 && m.fiveHundreds == 0 && m.thousands == 0 && m.fiveThousands == 0) {
		t.Fatal(120, m)
	} else if atm.thirties != (encash().thirties-4) || atm.fiveThousands != (encash().fiveThousands) {
		t.Log("Initially", encash())
		t.Fatal(120, "ATM", atm)
	}

	atm = encash()
	m = atm.giveMeMoneyDynamic(1190)
	if m == nil || !(m.thirties == 3 && m.fifties == 0 && m.hundreds == 1 && m.fiveHundreds == 0 && m.thousands == 1 && m.fiveThousands == 0) {
		t.Fatal(1190, m)
	} else if atm.fiveThousands != (encash().fiveThousands) || atm.thousands != (encash().thousands-1) || atm.fiveHundreds != (encash().fiveHundreds) {
		t.Fatal(1190, atm)
	}

	atm = encash()
	m = atm.giveMeMoneyDynamic(6620)
	if m == nil || !(m.thirties == 4 && m.fifties == 0 && m.hundreds == 0 && m.fiveHundreds == 1 && m.thousands == 1 && m.fiveThousands == 1) {
		t.Fatal(6620, m)
	}

	atm = encash()
	m = atm.giveMeMoneyDynamic(32160)
	if m == nil || !(m.thirties == 22 && m.fifties == 100 && m.hundreds == 5 && m.fiveHundreds == 2 && m.thousands == 5 && m.fiveThousands == 4) {
		t.Fatal(32160, m)
	}

	atm = encash()
	m = atm.giveMeMoneyDynamic(32190)
	if m == nil || !(m.thirties == 23 && m.fifties == 100 && m.hundreds == 5 && m.fiveHundreds == 2 && m.thousands == 5 && m.fiveThousands == 4) {
		t.Fatal(32190, m)
	} else if !(atm.thirties == 0 && atm.fifties == 0 && atm.hundreds == 0 && atm.fiveHundreds == 0 && atm.thousands == 0 && atm.fiveThousands == 0) {
		t.Fatal(32190, atm)
	}

	atm = encash()
	m = atm.giveMeMoneyDynamic(270)
	if m == nil || !(m.thirties == 4 && m.fifties == 1 && m.hundreds == 1 && m.fiveHundreds == 0 && m.thousands == 0 && m.fiveThousands == 0) {
		t.Fatal(270, m)
	} else if !(atm.thirties == 19 && atm.fifties == 99 && atm.hundreds == 4 && atm.fiveHundreds == 2 && atm.thousands == 5 && atm.fiveThousands == 4) {
		t.Fatal(270, atm)
	}

	atm = encash()
	m = atm.giveMeMoneyDynamic(100000)
	if m != nil {
		t.Fatal("overdraft 100000", m)
	}
}

type dynTable struct {
	sum   int
	notes int
}

func (atm *ATMwallet) giveMeMoneyDynamic(sum int) *withdrawMoney {
	if sum == 0 || sum > atm.totalMoneyInATM || sum > ATMlimit { //|| sum%30 != 0 && sum%50 != 0 {
		return nil
	}

	numLen := 1
	for i := sum; i >= 10; i /= 10 {
		numLen++
	}
	zeros := 0
	for i := sum; i > 0 && i%10 == 0; i /= 10 {
		sum /= 10
		zeros++
	}

	T := make([][]dynTable, 6+1)
	for i := range T {
		T[i] = make([]dynTable, sum+1)
	}
	for i := range T[0] {
		T[0][i].notes = 1 << 31
	}

	ansVal := 0
	for val := 1; val <= 6; val++ {
		for sumNeeded := 1; sumNeeded <= sum; sumNeeded++ {
			if nom, ok := atm.fit(val, sumNeeded, zeros); !ok { //todo: count availability of num
				for ; nom > sumNeeded && sumNeeded <= sum; sumNeeded++ {
					T[val][sumNeeded] = T[val-1][sumNeeded]
				}
			}

			valAmount, notesAmount := atm.getNomMax(val, sumNeeded, zeros)
			if T[val-1][sumNeeded].notes < (notesAmount + T[val-1][sumNeeded-valAmount].notes) {
				T[val][sumNeeded] = T[val-1][sumNeeded]
			} else {
				if val-1 == 0 {
					T[val][sumNeeded].notes = notesAmount
					T[val][sumNeeded].sum = valAmount
				} else {
					T[val][sumNeeded].notes = T[val-1][sumNeeded-valAmount].notes + notesAmount
					T[val][sumNeeded].sum = T[val-1][sumNeeded-valAmount].sum + valAmount
				}
			}
			//T[val][sumNeeded].notes = min(T[val-1][sumNeeded].notes, notesAmount + T[val-1][sumNeeded-valAmount].notes)

			if sumNeeded == sum && T[val][sumNeeded].sum == sum {
				ansVal = val
			}
		}
	}

	withdraw := &withdrawMoney{}
	return atm.findAns(T, ansVal, sum, withdraw)
} // todo

func (atm *ATMwallet) getNomMax(nom, sumNeeded, zeros int) (valAmount, notesAmount int) {
	nomVal := new(nomValue)

	switch nom {
	case fiveThousands:
		nomVal.amountLeft = atm.fiveThousands
		nomVal.nom = 5000
	case thousands:
		nomVal.amountLeft = atm.thousands
		nomVal.nom = 1000
	case fiveHundreds:
		nomVal.amountLeft = atm.fiveHundreds
		nomVal.nom = 500
	case hundreds:
		nomVal.amountLeft = atm.hundreds
		nomVal.nom = 100
	case fifties:
		nomVal.amountLeft = atm.fifties
		nomVal.nom = 50
	case thirties:
		nomVal.amountLeft = atm.thirties
		nomVal.nom = 30
	default:
		panic("WHAT")
	}
	nomVal.nom /= int(math.Pow10(zeros))

	canGiveYou := 0
	notes := 0
	valCap := nomVal.nom * nomVal.amountLeft
	for nomVal.nom <= sumNeeded && valCap > 0 {
		canGiveYou += nomVal.nom
		notes++

		sumNeeded -= nomVal.nom
		nomVal.amountLeft--
		valCap = nomVal.nom * nomVal.amountLeft
	}
	return canGiveYou, notes
}

func (atm *ATMwallet) fit(nom int, sumNeeded int, zeros int) (int, bool) {
	nomVal := new(nomValue)

	switch nom {
	case fiveThousands:
		nomVal.amountLeft = atm.fiveThousands
		nomVal.nom = 5000
	case thousands:
		nomVal.amountLeft = atm.thousands
		nomVal.nom = 1000
	case fiveHundreds:
		nomVal.amountLeft = atm.fiveHundreds
		nomVal.nom = 500
	case hundreds:
		nomVal.amountLeft = atm.hundreds
		nomVal.nom = 100
	case fifties:
		nomVal.amountLeft = atm.fifties
		nomVal.nom = 50
	case thirties:
		nomVal.amountLeft = atm.thirties
		nomVal.nom = 30
	default:
		panic("WHAT")
	}
	nomVal.nom /= int(math.Pow10(zeros))

	if nomVal.nom == 0 || nomVal.amountLeft == 0 || nomVal.nom > sumNeeded {
		return nomVal.nom, false
	}
	return nomVal.nom, true
}

func (atm *ATMwallet) findAns(T [][]dynTable, val, notes int, w *withdrawMoney) *withdrawMoney {
	if val == 0 || notes == 0 {
		return w
	}

	if T[val-1][notes].sum == T[val][notes].sum {
		return atm.findAns(T, val-1, notes, w)
	} else {
		atm.coinToDyn(w, T[val][notes].notes, val)
		return atm.findAns(T, val-1, notes-T[val][notes].notes, w)
	}
}

//d(i−1,c)				 	for c=0,…,wi−1;
//d(i,c)=
//		min(d(i−1,c),d(i,c−wi)+1)	for c=wi,…,W;
//for k = 1 to n
//	for s = 1 to w                                            //Перебираем для каждого k все вместимости
//		if s >= w[k]                                            //Если текущий предмет вмещается в рюкзак
//			A[k][s] = max(A[k - 1][s], A[k - 1][s - w[k]] + p[k]) //Выбираем класть его или нет
//		else
//			A[k][s] = A[k - 1][s]

func (atm *ATMwallet) coinToDyn(w *withdrawMoney, notes, nom int) {
	switch nom {
	case fiveThousands:
		atm.fiveThousands -= notes
		w.fiveThousands += notes
	case thousands:
		atm.thousands -= notes
		w.thousands += notes
	case fiveHundreds:
		atm.fiveHundreds -= notes
		w.fiveHundreds += notes
	case hundreds:
		atm.hundreds -= notes
		w.hundreds += notes
	case fifties:
		atm.fifties -= notes
		w.fifties += notes
	case thirties:
		atm.thirties -= notes
		w.thirties += notes
	default:
		panic("WHAT")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func (m *withdrawMoney) String() string {
	str := "\nWallet:\n" + `5000: ` + strconv.Itoa(m.fiveThousands) + "\t"
	str += `1000: ` + strconv.Itoa(m.thousands) + "\t"
	str += `500: ` + strconv.Itoa(m.fiveHundreds) + "\t"
	str += `100: ` + strconv.Itoa(m.hundreds) + "\t"
	str += `50: ` + strconv.Itoa(m.fifties) + "\t"
	str += `30: ` + strconv.Itoa(m.thirties)

	return str
}
func (m *ATMwallet) String() string {
	str := "\nATM:\n" + `5000: ` + strconv.Itoa(m.fiveThousands) + "\t"
	str += `1000: ` + strconv.Itoa(m.thousands) + "\t"
	str += `500: ` + strconv.Itoa(m.fiveHundreds) + "\t"
	str += `100: ` + strconv.Itoa(m.hundreds) + "\t"
	str += `50: ` + strconv.Itoa(m.fifties) + "\t"
	str += `30: ` + strconv.Itoa(m.thirties)

	return str
}

func encash() *ATMwallet {
	return &ATMwallet{
		fiveThousands:   4,
		thousands:       5,
		fiveHundreds:    2,
		hundreds:        5,
		fifties:         100,
		thirties:        23,
		totalMoneyInATM: 32190,
		notesAmount:     139,
	}
}

