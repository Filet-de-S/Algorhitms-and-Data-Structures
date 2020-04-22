package main

import "fmt"

type N struct {
	val    int //unique
	left   *N
	right  *N
	parent *N
}

func InitTree() *N {
	root := &N{val: 'a'}

	root.right = &N{val: 'c'}
	root.right.parent = root

	root.left = &N{val: 'b'}
	root.left.parent = root
	root.left.left = &N{val: 'd'}
	root.left.left.parent = root.left
	root.left.right = &N{val: 'e'}
	root.left.right.parent = root.left
	root.left.right.left = &N{val: 'g'}
	root.left.right.left.parent = root.left.right
	root.left.right.right = &N{val: 'f'}
	root.left.right.right.parent = root.left.right

	return root
}

func main() {
	root := InitTree()
	fmt.Println(string(nearestParentWithoutRoot(root.left.left, root.left.right.right).val)) //d, f)) = b

	root = InitTree()
	fmt.Println(string(nearestParentWithoutRoot(root.right, root.left.right.left).val)) //c, g)) = a

	root = InitTree()
	fmt.Println(string(nearestParentWithoutRoot(root.left.right, root.left).val)) //e, b)) = b

	root = InitTree()
	fmt.Println(string(nearestParentWithoutRoot(root.left, root.left.right).val)) //e, b)) = broot := InitTree()

	root = InitTree()
	fmt.Println(string(nearestParentWithoutRoot2(root.left.left, root.left.right.right).val)) //d, f)) = b

	root = InitTree()
	fmt.Println(string(nearestParentWithoutRoot2(root.right, root.left.right.left).val)) //c, g)) = a

	root = InitTree()
	fmt.Println(string(nearestParentWithoutRoot2(root.left.right, root.left).val)) //e, b)) = b

	root = InitTree()
	fmt.Println(string(nearestParentWithoutRoot2(root.left, root.left.right).val)) //e, b)) = b
}

func nearestParentWithoutRoot2(one, two *N) *N {
	roots1 := []*N{}
	roots2 := []*N{}
	for one != nil {
		roots1 = append(roots1, one)
		one = one.parent
	}
	for two != nil {
		roots2 = append(roots2, two)
		two = two.parent
	}

	i1, i2 := len(roots1) - 1, len(roots2) - 1
 	for ; i1 >= 0 && i2 >= 0 && roots1[i1] == roots2[i2]; i1-- {
		i2--
	}
	if i1+1 == len(roots1) {
		return nil
	}
	if i1 < 0 {
		return roots1[0]
	} else if i2 < 0 {
		return roots2[0]
	}
	// if nil is possible, continue below, else return last of arr[maxBy_i]
	return roots1[i1+1]
}

func nearestParentWithoutRoot(one, two *N) *N {
	roots := map[*N]struct{}{}
	if one == two {
		return one
	}

	for one != nil {
		roots[one] = struct{}{}
		one = one.parent
	}
	for two != nil {
		if _, ok := roots[two]; ok {
			return two
		}
		two = two.parent
	}

	//check remains of ONE or TWO
	return nil
}

/*
 * Given two nodes of a tree,
 * method should return the deepest common ancestor of those nodes.
 *
 *          A
 *         / \
 *        B   C
 *       / \
 *      D   E
 *         / \
 *        G   F
 *
 *  (D, F) = B
 *  (C, G) = A
 *  (E, B) = B
 */