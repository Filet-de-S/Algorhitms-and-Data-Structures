package main

import (
	"fmt"
)
func main() {
	var s = []string{"1", "2", "9"}
	fmt.Println("init:",s, cap(s))

	modifySlice(s)
	fmt.Println("after modifySlice func:",s, cap(s), "\n")

	secondTest(&s)
	fmt.Println("after secondTest func:",s, cap(s), "\n")

	thirdTest(&s)
	fmt.Println("after thirdTest func:",s, cap(s), "\n")

	fourthTest(s)
	fmt.Println("after fourthTest func:",s, cap(s), "\n")

	s = append(s, "4", "5", "6")
	fmt.Println("before fifthTest func:",s, cap(s))
	fifthTest(&s)
	fmt.Println("after fifthTest func:",s, cap(s), "\n")


	s =[]string{"1","2","3","4", "5", "6"}
	fmt.Println("before sixthTest func:",s, cap(s))
	sixthTest(&s)
	fmt.Println("after sixthTest func:",s, cap(s))
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "11"
	fmt.Println("modifySlice:",i, cap(i))
}

func secondTest(i *[]string) {
	tmp := append(*i, "22")
	fmt.Println("secondTest:",*i, cap(*i), "tmp:", tmp)
}

func thirdTest(i *[]string) {
	tmp := append((*i)[:1], "3333")
	fmt.Println("thirdTest:",*i, cap(*i), "tmp:", tmp)
}

func fourthTest(i []string) {
	tmp := append(i[:0], "4444")
	fmt.Println("fourthTest:",i, cap(i), "tmp:", tmp)
}

func fifthTest(i *[]string) {
	*i = append((*i)[:2], (*i)[1:]...)
	fmt.Println("fifthTest 1:",*i, cap(*i))

	tmp := append((*i)[:1], "555", "55555") // случай, который нужен был.
													// при том tmp не нужен, а-ля "_ = append(*slice, sl...)
	fmt.Println("fifthTest 2:",*i, cap(*i), "tmp:", tmp)

	*i = append((*i)[:1], "999", "99999")
	fmt.Println("fifthTest 3:",*i, cap(*i))
}

func sixthTest(i *[]string) {
	_ = append((*i)[:2], (*i)[1:]...) // ???? ???????
	fmt.Println("sixthTest 1:",*i, cap(*i)) // ???????????? *exitITwhereIsTheDoor*
	_ = append((*i)[:1], "555", "55555")
	fmt.Println("sixthTest 2:",*i, cap(*i))
}
