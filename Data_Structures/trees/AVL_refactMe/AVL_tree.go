package AVL_refactMe

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type Data struct {
	Key int
	ID string
}

type Node struct {
	Data *Data
	Height int8
	Left *Node
	Right *Node
}

type AVL struct {
	root *Node
	//lock sync.RWMutex
}

func makeAlphabet() []string {
	alpha := make([]string, 0, 52)
	for i := 65; i< 91; i++ {
		alpha = append(alpha, string(i))
	}
	for i := 97; i < 123; i++ {
		alpha = append(alpha, string(i))
	}
	return alpha
}

func main() {
	//alpha := makeAlphabet()

	m := &AVL{}
	//for i := 0; i < 20; i++ {
	//	m.Insert(&Data{Key:i})
	//}
	m.Insert(&Data{Key:17})
	m.Insert(&Data{Key:10})
	m.Insert(&Data{Key:20})
	m.Insert(&Data{Key:50})
	m.Insert(&Data{Key:5})
	m.Insert(&Data{Key:18})
	m.Insert(&Data{Key:15})
	m.Insert(&Data{Key:8})
	m.Insert(&Data{Key:19})
	m.Insert(&Data{Key:100})
	m.Insert(&Data{Key:45})

	ch := make(chan *Node, 0)
	datas := make([]string, 0, 0)

	go m.WalkLive(ch)
	for c := range ch {
		datas = append(datas, strconv.Itoa(c.Data.Key) + " " + strconv.Itoa(int(c.Height)) + "; ")
		//fmt.Print(c, " ")
	}
	fmt.Println(datas)

	m.Delete(6)
	m.Delete(3)

	datas = make([]string, 0, 0)
	ch = make(chan *Node, 0)
	go m.WalkLive(ch)
	for c := range ch {
		datas = append(datas, strconv.Itoa(c.Data.Key) + " " + strconv.Itoa(int(c.Height)) + "; ")
		//fmt.Print(c, " ")
	}
	fmt.Println(datas)


	fmt.Println("\n", "tree capacity without extending", math.Pow(2.0, float64(m.root.Height)))
	fmt.Println("height = ", m.root.Height)
	fmt.Println(len(datas))

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO make iterative
}


//func (y *Node) del(key int) {
//	if key < y.Data.Key {
//		y.Left = nil
//	} else {
//		y.Right = nil
//	}
//}
//
//func (n *AVL_refactMe) Pop(key int) (*Node, error) {
//	if n.root == nil {
//		return nil, errors.New("404 not found")
//	}
//
//	tmp := n.root
//	y := tmp
//	heights := []func(){}
//	for tmp != nil && tmp.Data.Key != key {
//		y = tmp
//		heights = append(heights, tmp.fixHeight)
//		if key < tmp.Data.Key {
//			tmp = tmp.Left
//		} else {
//			tmp = tmp.Right
//		}
//	}
//	//todo switch && balance!
//	if tmp == nil || tmp.Data.Key != key {
//		return nil, errors.New("404 not found")
//	}
//
//	poped := tmp
//	if tmp.Left == nil && tmp.Right == nil { //// DEF DEL
//		y.del(key)
//		//tmp = nil // doen't work: address is like "ABC"
//		//y.Right = nil // work: address is like y.Right "ABC"
//	} else if tmp.Right == nil { //
//		if key < y.Data.Key {
//			y.Left = tmp.Left //y.Left.Left
//		} else {
//			y.Right = tmp.Left //y.Right.Left
//		}
//		//tmp = tmp.Left
//	} else if tmp.Left == nil {
//		if key < y.Data.Key {
//			y.Left = tmp.Right//y.Left.Right
//		} else {
//			y.Right = tmp.Right //y.Right.Right
//		}
//		//tmp = tmp.Right
//	} else { // left and right 			 // 	  p
//		// find most left in right				q   s
//		leftestFromRight := tmp.Right
//		heights := make([]func(), 1)
//
//		x := &Node{}
//		for leftestFromRight.Left != nil {
//			x = leftestFromRight
//			heights = append(heights, leftestFromRight.fixHeight)
//			leftestFromRight = leftestFromRight.Left
//		}
//		//toDel := leftestFromRight
//		leftestFromRight.Left = tmp.Left 		//	 tmp
//		if tmp.Right != leftestFromRight {	//		l   r
//			leftestFromRight.Right = tmp.Right
//		}
//		if key < y.Data.Key {
//			y.Left = leftestFromRight
//			//tmp = leftestFromRight
//		} else {
//			y.Right = leftestFromRight
//			//tmp = leftestFromRight
//		}
//		if x.Data != nil {
//			x.Left = nil
//		} else {
//			tmp.Right = nil
//		}
//		//leftestFromRight = nil // check
//		//heights[len(heights)-1] = nil
//		//heights = heights[:len(heights) - 1]
//		heights[0] = tmp.fixHeight
//		for i := len(heights) - 2; i >= 0; i-- {
//			heights[i]()
//		}
//	}
//	for i := len(heights) - 1; i >= 0; i-- {
//		heights[i]()
//	}
//	return poped, nil
//}


func (n *Node) findMin() *Node {
	if n.Left != nil {
		return n.Left.findMin()
	}
	return n
}

func (n *Node) removeMin() *Node {
	if n.Left == nil {
		return n.Right
	}
	n.Left = n.Left.removeMin()
	return n.balance()
}

func (a *AVL) Delete(key int) {
	a.root = del(a.root, key)
}

func del(n *Node, key int) *Node {
	if n == nil {
		return nil
	}

	if key < n.Data.Key {
		n.Left = del(n.Left, key)
	} else if key > n.Data.Key {
		n.Right = del(n.Right, key)
	} else { //  key == n.Data.Key
		lhs := n.Left
		rhs := n.Right
		n = nil
		if rhs == nil  {
			return lhs
		}
		min := rhs.findMin()
		min.Right = rhs.removeMin()
		min.Left = lhs
		return min.balance()
	}
	return n.balance()
}

func (n *AVL) LookFor(key int) (*Data, error) {
	tmp := n.root
	for tmp.Data.Key != key {
		if key < tmp.Data.Key {
			tmp = tmp.Left
		} else {
			tmp = tmp.Right
		}
	}
	if tmp.Data.Key != key {
		return nil, errors.New("node not found")
	}
	return tmp.Data, nil
}


func (n *Node) Walk(ch chan<- *Node) {
	if n == nil {
		return
	}
	if n.Left != nil {
		n.Left.Walk(ch)
	}
	ch <- n
	if n.Right != nil {
		n.Right.Walk(ch)
	}
}

func (n *AVL) WalkLive(ch chan<- *Node) {
	if n.root == nil {
		return
	}
	n.root.Walk(ch)
	close(ch)
}

//func (n *AVL_refactMe) InsertIT(data string) {
//	if n.root == nil {
//		n.root = n.root.newNode(data)
//		return
//	}
//	x := n.root
//	y := &Node{}
//
//	//defer func() {
//	//	n.root = n.root.balance()
//	//}()
//	//fmt.Println(n.root.Height+1)
//	fixHeights := [50]func(){}
//	i := 0
//
//	for x != nil {
//		fixHeights[i] = x.fixHeight
//		y = x
//		if data < x.Data {
//			x = x.Left
//		} else {
//			x = x.Right
//		}
//		i++
//	}
//	if data < y.Data {
//		y.Left = y.Left.newNode(data)
//	} else {
//		y.Right = y.Right.newNode(data)
//	}
//
//	for l := i - 1; l >= 0; l-- {
//		fixHeights[l]()
//	}
//	n.root = n.root.balance()
//}

func (n *Node) insertNode(data *Data) *Node {
	if n == nil {
		return n.newNode(data)
	}

	if data.Key < n.Data.Key {
		n.Left = n.Left.insertNode(data)
	} else {
		n.Right = n.Right.insertNode(data)
	}
	return n.balance()
}

func (n *AVL) Insert(data *Data) {
	if n.root == nil {
		n.root = n.root.newNode(data)
	} else {
		n.root = n.root.insertNode(data)
	}
}

func (n *Node) balance() *Node {
	n.fixHeight()
	// 														p
	//													   / \
	//						 							  a   q
	if n.balanceFactor() == 2 { //							 / \
		if n.Right.balanceFactor() < 0 { //					s	d
			n.Right = n.Right.rotateRight() //			   / \
		} 									//			  b	  c
		return n.rotateLeft() // height(q) - height(a) = 2
	} else if n.balanceFactor() == -2 {
		if n.Left.balanceFactor() > 0 {
			n.Left = n.Left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

func (n *Node) newNode(data *Data) *Node {
	return &Node{
		Data:   data,
		Height: 1,
	}
}

func (n *Node) height() int8 {
	if n != nil {
		return n.Height
	}
	return 0
}

func (n *Node) balanceFactor() int8 {
	return n.Right.height() - n.Left.height()
}
//
//void fixheight(node* p)
//{
//unsigned char hl = height(p->left);
//unsigned char hr = height(p->right);
//p->height = (hl>hr?hl:hr)+1;
//}

func (n *Node) fixHeight() {
	if n.Left.height() > n.Right.height() {
		n.Height = n.Left.height() + 1
	} else {
		n.Height = n.Right.height()  + 1
	}
}

func (n *Node) rotateRight() *Node { // правый поворот вокруг n    ––>    q
	//   													/	\		/	\
	//													   q	 c	   a	 n
	//													  / \				/ \
	//													 a	 b			   b   c
	q := n.Left
	n.Left = q.Right
	q.Right = n
	n.fixHeight()
	q.fixHeight()
	return q
}

func (q *Node) rotateLeft() *Node {  //  левый поворот вокруг q:	  n   <–––    q
	//			   													/	\		/	\
	//															   q	 c	   a	 n
	//															  / \				/ \
	//															 a	 b			   b   c
	n := q.Right
	q.Right = n.Left
	n.Left = q
	q.fixHeight()
	n.fixHeight()
	return n
}


