package main

import (
	"strconv"
	"testing"
	"time"
)

type pointTo = int


type uf struct {
	id 		[]pointTo
	size 	[]int
	user 	map[string]pointTo
}

func new(ids []string) *uf {

	u := &uf{
		id: make([]pointTo, len(ids)),
		user: make(map[string]pointTo, len(ids)),
		size: make([]int, len(ids)),
	}
	for i := range ids {
		u.id[i] = i
		u.user[ids[i]] = i
		u.size[i] = 1
	}
	return u
}

func (u *uf) union(p, q string) {
	pid := u.getRoot(u.id[ u.user[p] ])
	qid := u.getRoot(u.id[ u.user[q] ])

	switch {
	case pid == qid:
		return
	case u.size[pid] < u.size[qid]: // if pid tree smaller
		u.id[pid] = qid
		u.size[qid] += u.size[pid]
	default:
		u.id[qid] = pid
		u.size[pid] += u.size[qid]
	}
}

func (u *uf) isConnected(p, q string) bool {
	return u.getRoot(u.id[ u.user[p] ]) ==
		   u.getRoot(u.id[ u.user[q] ])
}

func (u *uf) getRoot(node int) int {
	for node != u.id[node] {
		u.id[node] = u.id[ u.id[node] ]
		node = u.id[node]
	}
	return node
}

//func (u *uf) find(toFind string) int {
//
//}

//func (u *uf) count() int {
//
//}

func (u *uf) badUnion(p, q string) {
	pid := u.id[ u.user[p] ]
	qid := u.id[ u.user[q] ]
	for i := range u.id {
		if u.id[i] == pid {
			u.id[i] = qid
		}
	}
}

func (u *uf) isConnectedFromBad(p, q string) bool {
	return u.id[ u.user[p] ] == u.id[ u.user[q] ]
}

func newBad(ids []string) *uf {

	u := &uf{
		id: make([]pointTo, len(ids)),
		user: make(map[string]pointTo, len(ids)),
		//size: make([]int, len(ids)),
	}
	for i := range ids {
		u.id[i] = i
		u.user[ids[i]] = i
	}
	return u
}


func TestUnionFind(t *testing.T) {
	//relations := [][]string{
	//	{"Kusha", "Dasha"}, {"Dasha", "Leido"}, {"Zoom", "MrNobody"},
	//	{"Antio","Kusha"},  {"Pasha", "Sasha"},	{"Leido", "Antio"},
	//	{"MrNobody", "Masha"}, {"Saiko", "Pasha"}, {"Zoom", "Sasha"},
	//	{"Sasha", "Masha"}, {"Zoom", "Saiko"},
	//}
	//users := []string{"Masha", "Sasha", "Pasha", "Dasha", "Kusha",
	//				"MrNobody", "Zoom", "Saiko", "Leido", "Antio"}
	// can imagine like that are IDs (string in DB)
	// 0 - Masha
	// 1 - Sasha
	// 2 - Pasha
	// 3 - Dasha
	// 4 - Kusha
	// 5 - MrNobody
	// 6 - Zoom
	// 7 - Saiko
	// 8 - Leido
	// 9 - Antio
	users := genUsers()
	//relations := getRelations()
	//fmt.Println("got users", len(users))
	u := new(users)
	now := time.Now()
	for j := 0; j < len(users) - 1; j++ {
		//fmt.Println("boec n ", j)
		if j == 5000 {
			break
		}
		for i := range users {
			//if i < 1 && i == len(users)-2 {
			//	continue
			//}
			p := users[j]
			q := users[i]
			if !u.isConnected(p, q) {
				if j == 1 {
					t.Fatal("J == 1")
				}
				u.union(p, q)
				//fmt.Println("joined", p, "and", q)
			}
		}
	}
	since := time.Since(now)
	t.Log("ok", since)
}

func TestUnionFindBad(t *testing.T) {
	//relations := [][]string{
	//	{"Kusha", "Dasha"}, {"Dasha", "Leido"}, {"Zoom", "MrNobody"},
	//	{"Antio","Kusha"},  {"Pasha", "Sasha"},	{"Leido", "Antio"},
	//	{"MrNobody", "Masha"}, {"Saiko", "Pasha"}, {"Zoom", "Sasha"},
	//	{"Sasha", "Masha"}, {"Zoom", "Saiko"},
	//}
	//users := []string{"Masha", "Sasha", "Pasha", "Dasha", "Kusha",
	//	"MrNobody", "Zoom", "Saiko", "Leido", "Antio"}
	// can imagine like that are IDs (string in DB)
	// 0 - Masha
	// 1 - Sasha
	// 2 - Pasha
	// 3 - Dasha
	// 4 - Kusha
	// 5 - MrNobody
	// 6 - Zoom
	// 7 - Saiko
	// 8 - Leido
	// 9 - Antio
	users := genUsers()
	//relations := getRelations()
	u := newBad(users)
	//fmt.Println("got users")
	now := time.Now()
	for j := 0; j < len(users) - 1; j++ {
		if j == 5000 {
			break
		}
		for i := range users {
			//if i < 1 && i == len(users)-2 {
			//	continue
			//}
			p := users[j]
			q := users[i]
			if !u.isConnectedFromBad(p, q) {
				if j == 1 {
					t.Fatal("J == 1")
				}
				u.badUnion(p, q)
				//t.Log("joined", p, "and", q)
			}
		}
	}
	since := time.Since(now)
	t.Log("bad", since)
}

func genUsers() []string {
	users := make([]string, 25000)
	for i := 0; i < 25000; i++ {
		users[i] = strconv.Itoa(i)
	}
	return users
}


type ufInt struct {
	id 		[]pointTo
	size 	[]int
	max 	int
}

func newInt(ids []int) *ufInt {

	u := &ufInt{
		id: make([]pointTo, len(ids)),
		size: make([]int, len(ids)),
	}
	for i := range ids {
		u.id[i] = i
		u.size[i] = 1
	}
	return u
}

func (u *ufInt) unionInt(p, q int) {
	pid := u.getRootInt(u.id[ p ])
	qid := u.getRootInt(u.id[ q ])

	switch {
	case pid == qid:
		return
	case u.size[pid] < u.size[qid]: // if pid tree smaller
		u.id[pid] = qid
		u.size[qid] += u.size[pid]
		if u.size[qid] > u.max {
			u.max = u.size[qid]
		}
	default:
		u.id[qid] = pid
		u.size[pid] += u.size[qid]
		if u.size[pid] > u.max {
			u.max = u.size[pid]
		}
	}
}

func (u *ufInt) isConnectedInt(p, q int) bool {
	return u.getRootInt(u.id[ p ]) ==
		u.getRootInt(u.id[ q ])
}

func (u *ufInt) getRootInt(node int) int {
	for node != u.id[node] {
		u.id[node] = u.id[u.id[node]]
		node = u.id[node]
	}
	return node
}

func TestConvertInt(t *testing.T) {
	users := genUsers()
	usersInt := make([]int, len(users))
	for u := range users {
		v, err := strconv.Atoi(users[u]); if err == nil {
			usersInt[u] = v
		} else {
			t.Fatal("ATOI")
		}
	}
	u := newInt(usersInt)
	now := time.Now()
	for j := 0; j < len(users) - 1; j++ {
		if j == 5000 {
			break
		}
		for i := range users {
			//if i < 1 && i == len(users)-2 {
			//	continue
			//}
			p, err1 := strconv.Atoi(users[j])
			q, err2 := strconv.Atoi(users[i])
			if err1 != nil || err2 != nil {
				t.Fatal("ATOI")
			}
			if j == 1 && u.max != len(users) {
				t.Fatal(u.max, len(users))
			}
			if !u.isConnectedInt(p, q) {
				if j == 1 {
					t.Fatal("J == 1")
				}
				u.unionInt(p, q)
				//t.Log("joined", p, "and", q)
			}
		}
	}
	since := time.Since(now)
	t.Log("int", since)
}