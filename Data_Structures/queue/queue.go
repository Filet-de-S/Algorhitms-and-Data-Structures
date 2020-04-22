package queue

type Node struct {
	Data interface{}
	next *Node
}
type Queue struct {
	head *Node
	last *Node
	size int
}


func (q *Queue) Add(val interface{}) {
	if q.head == nil {
		q.last = &Node{Data: val}
		q.size = 1
		q.head = q.last
	} else {
		q.last.next = &Node{Data: val}
		q.last = q.last.next
		q.size++
	}
}

func (q *Queue) Pop() (data interface{}) {
	if q.head != nil {
		data = q.head.Data
		q.head = q.head.next
		q.size--
	}
	return data
}

func (q *Queue) Size() int {
	if q.head == nil {
		return 0
	}
	return q.size
}
