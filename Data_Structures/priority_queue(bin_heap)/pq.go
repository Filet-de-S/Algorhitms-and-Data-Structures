package queue

import "errors"

type EmptyHeap error

type Data struct {
	CmpData int         // add comparable field and change type if need
	Data    interface{} // leftover
}

type Heap struct {
	arr []*Data
	heapType byte
	size int
}

func New(heapType string) *Heap {
	if heapType == "min" {
		return &Heap{
			arr:     make([]*Data, 1),
			heapType: 'n',
			size:     0,
		}
	}
	return &Heap{
		arr:     make([]*Data, 1),
		heapType: 'x',
		size:     0,
	}
}

func (h *Heap) Insert(data *Data) {
	h.arr = append(h.arr, data)
	h.size++
	h.bubbleUp(h.size)
}

func (h *Heap) Pop() (data *Data, err error) {
	if h.size > 0 {
		data, err = h.arr[1], nil
		h.arr[1] = h.arr[h.size]
		h.arr = h.arr[:h.size]
		h.size--
		h.bubbleDown(1)
		return data, err
	}
	return nil, EmptyHeap(errors.New(""))
}

func (h *Heap) GetRoot() (data Data, err error) {
	if h.size > 0 {
		return *h.arr[1], nil
	}
	return Data{}, EmptyHeap(errors.New(""))
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) bubbleUp(child int) {
	for child > 1 && h.cmp(child, child/2) {
		h.swap(child, child/2)
		child = child/2
	}
}

func (h *Heap) bubbleDown(parent int) {
	for parent*2 <= h.size {
		child := parent*2
		if child < h.size && h.cmp(child+1, child) {
			child++
		}
		if h.cmp(parent, child) {
			break
		}
		h.swap(parent, child)
		parent = child
	}
}

func (h *Heap) cmp(child, parent int) bool {
	if h.heapType == 'x' {
		// MAX: if child > parent
		return h.arr[child].CmpData > h.arr[parent].CmpData
	}
	// MIN: if child < parent
	return h.arr[child].CmpData < h.arr[parent].CmpData
}

func (h *Heap) swap(child, parent int) {
	h.arr[child], h.arr[parent] = h.arr[parent], h.arr[child]
}