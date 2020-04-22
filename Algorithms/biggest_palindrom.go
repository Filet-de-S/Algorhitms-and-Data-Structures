package main

import (
	"fmt"
)

func biggestPal(str string) string {
	if str == "" {
		return ""
	}
	start, end, length := 0, 0, 0
	l := len(str)
	for i := 0; i < l; i++ {
		odd := aroundCenter(str, i, i)
		even := aroundCenter(str, i, i + 1)
		if odd > even {
			length = odd
		} else {
			length = even
		}
		if length > end - start {
			start = i - (length - 1) / 2 // 1 5 1 =OK; 1 22 1 OK;
			end = i + length/ 2         // 1 5 1 =OK; 1 22 1 OK;
		}
	}
	return str[start: end+1]
}

func aroundCenter(str string, lhs int, rhs int) int {
	l := len(str)
	for lhs >= 0 && rhs < l && str[lhs] == str[rhs] {
		lhs--
		rhs++
	}
	return rhs - lhs - 1 // 1 5 1
}



func main() {
	cases := []string{
		//"onetwo", "abccba", "aaa aaa", "nag gan", "nag g gan", "nagggan", "nagdan", "nagrgan", "a", "ab", "aa"}
		//"ibvjkmpyzsifuxcabqqpahjdeuzaybqsrsmbfplxycsafogotliyvhxjtkrbzqxlyfwujzhkdafhebvsdhkkdbhlhmaoxmbkqiwiusngkbdhlvxdyvnjrzvxmukvdfobzlmvnbnilnsyrgoygfdzjlymhprcpxsnxpcafctikxxybcusgjwmfklkffehbvlhvxfiddznwumxosomfbgxoruoqrhezgsgidgcfzbtdftjxeahriirqgxbhicoxavquhbkaomrroghdnfkknyigsluqebaqrtcwgmlnvmxoagisdmsokeznjsnwpxygjjptvyjjkbmkxvlivinmpnpxgmmorkasebngirckqcawgevljplkkgextudqaodwqmfljljhrujoerycoojwwgtklypicgkyaboqjfivbeqdlonxeidgxsyzugkntoevwfuxovazcyayvwbcqswzhytlmtmrtwpikgacnpkbwgfmpavzyjoxughwhvlsxsgttbcyrlkaarngeoaldsdtjncivhcfsaohmdhgbwkuemcembmlwbwquxfaiukoqvzmgoeppieztdacvwngbkcxknbytvztodbfnjhbtwpjlzuajnlzfmmujhcggpdcwdquutdiubgcvnxvgspmfumeqrofewynizvynavjzkbpkuxxvkjujectdyfwygnfsukvzflcuxxzvxzravzznpxttduajhbsyiywpqunnarabcroljwcbdydagachbobkcvudkoddldaucwruobfylfhyvjuynjrosxczgjwudpxaqwnboxgxybnngxxhibesiaxkicinikzzmonftqkcudlzfzutplbycejmkpxcygsafzkgudy",
			"cbsbd"}
	for i, n := range cases {
		fmt.Println(n, biggestPal(cases[i]))
	}
}