package BinaryIndexTree_Ryabko

type Tree struct {
	Tree []int
}

func Init(arr []int) *Tree {
	fen := &Tree{make([]int, len(arr)) }

	for i, l := 0, len(arr); i < l; i++ {
		fen.Add(i, arr[i])
	}
	return fen
}

func (t *Tree) Add(pos, val int) {
	for l := len(t.Tree); pos < l; pos |= pos+1 {
		t.Tree[pos] += val
	}
}

func (t *Tree) Update(pos, newValue int) {
	prevSum := t.rsq(pos-1)
	oldSum := t.rsq(pos)

	oldValue := oldSum - prevSum
	t.Add(pos, -(oldValue - newValue))
}

func (t *Tree) RangeSumQuery(l, r int) int {
	return t.rsq(r) - t.rsq(l-1)
}

func (t *Tree) rsq(i int) (sum int) {
	if i >= len(t.Tree) || i < 0 {
		return 0
	}

	for ; i >= 0; i = (i & (i +1)) - 1 {
		sum += t.Tree[i]
	}
	return sum
}

