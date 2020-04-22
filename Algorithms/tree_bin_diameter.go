package main

import (
	"fmt"
)

type node struct {
	left *node
	right *node
}

func main() {
	root := &node{}

	root.right = &node{}
	root.right.right = &node{}
	root.right.right.left = &node{}
	root.right.right.right = &node{}
	root.right.right.right.left = &node{}

	root.left = &node{}
	root.left.left = &node{}
	root.left.right = &node{}
	root.left.right.left = &node{}
	root.left.right.left.right = &node{}

	l := getDiameterOfTree(root)
	fmt.Println(l)
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

func getDiam(root *node, fin *int) int {
	if root == nil {
		return 0
	}

	lh := getDiam(root.left, fin) // 2
	rh := getDiam(root.right, fin) // 4

	*fin = max(*fin , lh + rh + 1)
	return max(lh, rh) + 1
}

func getDiameterOfTree(root *node) (fin int) {
	getDiam(root, &fin)
	return
}