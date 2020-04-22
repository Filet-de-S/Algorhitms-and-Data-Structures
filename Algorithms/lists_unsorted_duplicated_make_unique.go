package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
     Val int
     Next *ListNode
}

func delDups(list *ListNode) {
	m := make(map[int]struct{})
	tmp := list
	prev := &ListNode{}
	for tmp != nil {
		if _, ok := m[tmp.Val]; ok {
			//exists
			prev.Next = tmp.Next
		} else {
			m[tmp.Val] = struct{}{}
			prev = tmp
		}
		tmp = tmp.Next
	}
}

func delDupsUnsorted(list *ListNode) *ListNode {
	if list == nil || list.Next == nil{
		return list
	}
	arr := []int{}
	tmp := list
	for tmp != nil {
		arr = append(arr, tmp.Val)
		tmp = tmp.Next
	}
	sort.Ints(arr) 	// of cos we could sort lists, but lets be easy, if task fits )))
	// we could del dups in arr, but lets do SOMETHING

	new := &ListNode{Val: arr[0]}
	list = new
	for i:=1; i<len(arr); i++ {
		new.Next = &ListNode{Val:arr[i]}
		new = new.Next
	}
	//fmt.Println(list, list.Next, list.Next.Next, list.Next.Next.Next, list.Next.Next.Next.Next)
	delDupsSorted(list)
	fmt.Println(list, list.Next, list.Next.Next)

	return new
}

func delDupsSorted(head *ListNode) {
	if head == nil {
		return
	}

	tmp := head
	for tmp != nil && tmp.Next != nil {
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
}

func delDupsIfUniqueKeep(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	tmp := &ListNode{}
	tmp.Next = head

	p := tmp
	for p.Next != nil && p.Next.Next != nil {
		if p.Next.Val == p.Next.Next.Val {
			dup := p.Next.Val
			for p.Next != nil && p.Next.Val == dup {
				p.Next = p.Next.Next
			}
		} else {
			p = p.Next
		}
	}
	return tmp.Next
}

func delDupsN2Memory1(list *ListNode) {
	check := list
	tmp := list
	for tmp != nil {
		check = tmp
		for check.Next != nil {
			if check.Next.Val == tmp.Val {
				check.Next = check.Next.Next
			} else {
				check = check.Next
			}
		}
		tmp = tmp.Next
	}
}

func main() {
	list := &ListNode{
			Val: 4, Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: -1,
								Next: &ListNode{
									Val: 10,
									Next: &ListNode{
										Val:  4,
										Next: nil,
									},
								},
							},
						},
					},
				},
			}}}
	printList(list)
	delDupsN2Memory1(list)
	printList(list)

	//list = delDupsUnsorted(list)
}

func printList(list *ListNode) {
	tmp := list
	for tmp != nil {
		fmt.Print(tmp.Val, " ")
		tmp = tmp.Next
	}
	fmt.Println()
}
