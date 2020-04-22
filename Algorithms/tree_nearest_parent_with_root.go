package main

import (
	"fmt"
)

type n struct {
	val int //unique
	left *n
	right *n
	parent *n
}

func nearestParentWithRootOK(root *n, p *n, q *n) *n {
	val := false
	v := valid{entry:true, valid: &val}
	n := validParent(root, p, q, &v)
	if val {
		return n
	}
	fmt.Println(val, string(n.val))
	return nil
}

type valid struct {
	valid 			*bool
	entry 			bool
	oneIsFound		bool
	checkFirstIf	bool
}

func validParent(root, p, q *n, v *valid) *n {
	if root == nil {
		return nil
	}
	if v.entry {
		v.entry = false
	 	if p == root || q == root {
			if p == q {
				*v.valid = true
				return root
			}
			v.oneIsFound = true
		}
	}
	if v.checkFirstIf && (p == root || q == root) {
		return root
	}
	v.checkFirstIf = true

	left := validParent(root.left, p, q, v)
	right := validParent(root.right, p, q, v)
	if left != nil && right != nil {
		*v.valid = true
		return root
	}
	if left != nil {
		if v.oneIsFound {
			*v.valid = true
			return root
		}
		if left == p && left == q {
			*v.valid = true
			return left
		}
		v.checkFirstIf = false
		right = validParent(left, p, q, v)
		if right != nil {
			*v.valid = true
			return left
		}
		return left
	}
	if right != nil {
		if v.oneIsFound {
			*v.valid = true
			return root
		}
		if right == p && right == q {
			*v.valid = true
			return right
		}
		v.checkFirstIf = false
		left = validParent(right, p, q, v)
		if left != nil {
			*v.valid = true
			return right
		}
		return right
	}
	return nil
}


func initTree() *n {
	root := &n{val: 'a'}

	root.right = &n{val: 'c'}
	root.right.parent = root

	root.left = &n{val: 'b'}
	root.left.parent = root
	root.left.left = &n{val: 'd'}
	root.left.left.parent = root.left
	root.left.right = &n{val: 'e'}
	root.left.right.parent = root.left
	root.left.right.left = &n{val: 'g'}
	root.left.right.left.parent = root.left.right
	root.left.right.right = &n{val: 'f'}
	root.left.right.right.parent = root.left.right

	return root
}

func main() {

	root := initTree()
	wrong := &n{val:'x'}
	fmt.Println("ans", nearestParentWithRootOK(root, root.left, wrong), "== nil") //b, x)) = b

	root = initTree()
	fmt.Println("ans", string(nearestParentWithRootOK(root, root.left.left, root.left.right.right).val), "== b") //d, f)) = b

	root = initTree()
	fmt.Println("ans", string(nearestParentWithRootOK(root, root.left.right.right, root.left.left).val), "== b") //d, f)) = b

	root = initTree()
	fmt.Println("ans", string(nearestParentWithRootOK(root, root.right, root.left.right.left).val), "== a") //c, g)) = a

	root = initTree()
	fmt.Println("ans", string(nearestParentWithRootOK(root, root.left.right, root.left).val), "== b") //e, b)) = b

	root = initTree()
	fmt.Println("ans", string(nearestParentWithRootOK(root, root.left.right, root.right).val), "== a") //e, c)) = a

	root = initTree()
	fmt.Println("ans", string(nearestParentWithRootOK(root, root, root.left).val), "== a") // = a

	root = initTree()
	fmt.Println("ans", string(nearestParentWithRootOK(root, root.left, root.left).val), "== b") // = b

	root = initTree()
	fmt.Println("ans", string(nearestParentWithRootOK(root, root.right, root.left).val), "== a") // = a
}


func iterativeParent(root, one, two *n) *n {
	if root == nil || one == nil || two == nil {
		return nil
	}

	route := []*n{}
	route = append(route, root)
	oneFl := false
	twoFl := false
	parent := &n{}
	parent = nil

	route = append(route, root)
	for len(route) > 0 {
		if parent != nil {
			fmt.Println(string(parent.val))
		}
		if root == one {
			oneFl = true
			if parent == nil {
				//fmt.Println("parent __one")
				parent = root
			}
		}
		if root == two {
			twoFl = true
			if parent == nil {
				//fmt.Println("parent __two")
				parent = root
			}
		}
		if oneFl == true && twoFl == true {
			return parent
		}

		if root.left != nil {
			root = root.left
			route = append(route, root)
			continue
		}
		if root.right != nil {
			root = root.right
			route = append(route, root)
			continue
		}

		for len(route)-2 > -1 { // at least one child from root
			tmp := route[len(route)-2]
			if parent != nil && (root == parent || tmp == parent) {
				//parent = tmp
				parent = refreshParent(route, parent)
			}

			if tmp.right != nil && tmp.right != root {
				//if parent != root {
				//	parent = tmp
				//}
				//fmt.Println("parent __Verx Right")
				root = tmp.right
				route[len(route)-1] = root
				break
			}
			root = tmp
			route = route[:len(route)-1]
			//if parent != nil && parent == tmp {
			//}
			//parent = route[len(route)-1]
			//fmt.Println("parent __Verx")
		}
	}
	return nil
	//return nearestParentWithRootOK(root, left, right)
}

	func refreshParent(route []*n, parent *n) *n {
		for i := range route {
			if i > 0 && route[i] == parent {
				return route[i-1]
			}
		}
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