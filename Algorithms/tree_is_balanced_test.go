package main

import (
	"testing"
)

type node struct {
	t *tree
	n *node
}
type queue struct {
	head *node
	last *node
	size int
}


func (q *queue) add(val *tree) {
	if q.head == nil {
		q.last = &node{t:val}
		q.size = 1
		q.head = q.last
	} else {
		q.last.n = &node{t:val}
		q.last = q.last.n
		q.size++
	}
}

func (q *queue) pop() *tree {
	var tmp *tree
	if q.head != nil {
		tmp = q.head.t
		q.head = q.head.n
	}
	if q != nil {
		q.size--
	}
	return tmp
}

func (q *queue) siz() int {
	if q == nil || q.head == nil {
		return 0
	}
	return q.size
}

type tree struct {
	l *tree
	r *tree
	val int
}

func treeHeight(t *tree) int {
	if t == nil {
		return 0
	}
	q := &queue{last:&node{t:t}, size:1}
	q.head = q.last
	//fmt.Println(string(t.val))
	h := 0
	for {
		s := q.siz()
		if s == 0 {
			return h
		}
		h++
		//fmt.Println("+1")
		for s > 0 {
			tmp := q.pop()
			if tmp.l != nil {
				q.add(tmp.l)
				//fmt.Println(string(tmp.l.val))
			}
			if tmp.r != nil {
				q.add(tmp.r)
				//fmt.Println(string(tmp.r.val))
			}
			s--
		}
	}
}


func isBalFull(t *tree) int {
	if t == nil {
		return 0
	}
	//fmt.Println(string(t.val))
	qL := &queue{}
	qR := &queue{}
	if t.l != nil {
		qL.add(t.l)
	}
	if t.r != nil {
		qR.add(t.r)
	}

	left, right := false, false
	caution := false
	lastL, lastR := 0,0
	lastLc, lastRc := 0,0

	for {
		lastL, left = checkSide(qL)
		lastR, right = checkSide(qR)
		switch {
		case left == false && right == false:
			return 0
		case caution == true:
			return -1
		case left && right && (lastLc == 1 || lastRc == 1):
			return -1
		case left == false || right == false:
			caution = true
		}
		lastLc, lastRc = lastL, lastR
	}
}

func checkSide(q *queue) (l int, side bool) {
	s := q.siz()
	l = s
	for s > 0 {
		tmp := q.pop()
		if tmp != nil {
			if tmp.l != nil {
				q.add(tmp.l)
			}
			if tmp.r != nil {
				q.add(tmp.r)
			}
			side = true
		}
		s--
	}
	return l, side
}

func treeIsBal(t *tree) int {
	l, r := 0, 0
	if t != nil {
		l = treeHeight(t.l)
		r = treeHeight(t.r)
	}
	dif := l - r
	if dif > 1 || dif < -1 {
		return -1
	}
	return 0
}

func TestTreeIsBalancedInHeight(test *testing.T) {
	t := &tree{val:'a'}
	t.l = &tree{val:'b',l: &tree{val:'c',l: &tree{val:'d'}}} // 3
	t.r = &tree{val:'z',l: &tree{val:'y',r: &tree{val:'t',r: &tree{val:'e',  }}}}
	t.r.r = &tree{val:'k', r:&tree{val:'r'}}
	if treeIsBal(t) != 0 {
		test.Fatal("expected 0")
	}
	//fmt.Println(isBalFull(t))

	//	    a
	//	   / \
	//	  b   z
	//	  /   /\
	//	  c  y	k
	//	 /	  \	 \
	//  d	   t  r
	//			e
}

func TestTreeIsBalancedInHeightFalse(test *testing.T) {
	t := &tree{val:'a'}
	t.l = &tree{val:'b',l: &tree{val:'c',l: &tree{val:'d'}}} // 3
	t.r = &tree{val:'z',l: &tree{val:'y',r: &tree{val:'t',r: &tree{val:'e', l:&tree{} }}}}
	t.r.r = &tree{val:'k', r:&tree{val:'r'}}
	if treeIsBal(t) != -1 {
		test.Fatal("expected -1")
	}
	//fmt.Println(isBalFull(t))

	//	    a
	//	   / \
	//	  b   z
	//	  /   /\
	//	  c  y	k
	//	 /	  \	 \
	//  d	   t  r
	//			e
	//			/
	//			*
}


func TestTreeIsBalFullFalse(test *testing.T) {
	t := &tree{val:'a'}
	t.l = &tree{val:'b',l: &tree{val:'c',l: &tree{val:'d'}}} // 3
	t.r = &tree{val:'z',l: &tree{val:'y',r: &tree{val:'t',r: &tree{val:'e',  }}}}
	t.r.r = &tree{val:'k', r:&tree{val:'r'}}
	if isBalFull(t) != -1 {
		test.Fatal("expected -1")
	}

	//	    a
	//	   / \
	//	  b   z
	//	  /   /\
	//	  c  y	k
	//	 /	  \	 \
	//  d	   t  r
	//			e
}

func TestTreeIsBalFull(test *testing.T) {
	t := &tree{val:'a'}
	t.l = &tree{val:'b',r: &tree{val:'s'},l: &tree{val:'c',l: &tree{val:'d'}}} // 3
	t.r = &tree{val:'z',l: &tree{val:'y',r: &tree{val:'t',}}}
	t.r.r = &tree{val:'k', r:&tree{val:'r'}}
	if isBalFull(t) != -1 {
		test.Fatal("expected -1")
	}

	//	    a
	//	   / \
	//	  b   z
	//	  /\  /\
	//	  c s y	k
	//	 /	  \	 \
	//  d	   t  r
	//			e
}