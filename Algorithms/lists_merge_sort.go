package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	list := &ListNode{
		Val:4, Next:&ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		}}
	fmt.Println("unsorted")
	printList(list)

	list = sortList(list)

	fmt.Println("sorted")
	printList(list)
}

func listLen(head *ListNode) int {
	tmp := head
	i := 0
	for ;tmp != nil; i++ {
		tmp = tmp.Next
	}
	return i
}

func getSides(head, lh, rh **ListNode) {
	l := listLen(*head)
	if l < 2 {
		*lh = *head
		*rh = nil
	} else {
		mid := (l-1)/2
		tmp := *head
		for i:=0; i<mid; i++ {
			tmp = tmp.Next
		}
		*lh = *head
		*rh = tmp.Next
		tmp.Next = nil
	}
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	lh, rh := &ListNode{}, &ListNode{}
	lh, rh = nil, nil
	getSides(&head, &lh, &rh)
	lh = sortList(lh)
	rh = sortList(rh)
	head = mergeSorted(lh, rh)
	return head
}

func mergeSorted(lh, rh *ListNode) *ListNode {
	if lh == nil {
		return rh
	} else if rh == nil {
		return lh
	}
	res := &ListNode{}
	if lh.Val < rh.Val {
		res = lh
		res.Next = mergeSorted(lh.Next, rh)
	} else {
		res = rh
		res.Next = mergeSorted(lh, rh.Next)
	}
	return res
}

func printList(list *ListNode) {
	tmp := list
	for tmp != nil {
		fmt.Print(tmp.Val, " ")
		tmp = tmp.Next
	}
	fmt.Println()
}