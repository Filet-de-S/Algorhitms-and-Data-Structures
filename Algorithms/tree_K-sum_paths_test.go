package main

import (
	"fmt"
	"testing"
)

type tree struct {
	data int
	left *tree
	right *tree
}

func TestK_sumTree(t *testing.T) {
	tree := getKsumTree()
	treeKsum(tree, 5)
}

func treeKsum(t *tree, sum int)  {
	deep := treeDeep(t)
	path := make([]int, deep)
	treeKsumFind(t, sum, path, 0)
}

func treeKsumFind(t *tree, sum int, path []int, level int) {
	if t == nil {
		return
	}
	path[level] = t.data
	curSum := 0
	for i := level; i >= 0; i-- {
		curSum += path[i]
		if curSum == sum {
			printPath(path, i, level)
		}
	}
	treeKsumFind(t.left, sum, path, level + 1)
	treeKsumFind(t.right, sum, path, level + 1)
}

func printPath(path[]int, i int, level int) {
	if path == nil {
		return
	}
	for ; i <= level; i++ {
		fmt.Print(path[i], " ")
	}
	fmt.Println()
}

func treeDeep(t *tree) int {
	if t == nil {
		return 0
	}
	return 1 + max(treeDeep(t.left), treeDeep(t.right))
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

func getKsumTree() *tree {
	return &tree{
		data:  1,
		left:  &tree{data: 3,
			left: &tree{data: 2},
			right: &tree{data:1,
				left: &tree{data: 1}}},
		right: &tree{data: -1,
			left: &tree{data: 4,
				left: &tree{data:1},
				right: &tree{data:2}},
			right: &tree{data: 5, right:
				&tree{data: 6}}}}
}