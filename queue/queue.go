package queues

type Node struct {
	data rune
	next *Node
}

func newNode(data rune) *Node {
	return &Node{data: data}
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func (q Queue) NewQueue() Queue {
	return Queue{}
}

func (q *Queue) EnQueue(data rune) {
	node := newNode(data)
	if q.head == nil {
		q.head = node
		q.tail = q.head
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.length++
}

func (q *Queue) DeQueue() rune {
	if q.head == nil {
		return 0
	}
	node := q.head

	q.head = q.head.next
	q.length--

	return node.data
}

func (q Queue) Len() int {
	return q.length
}
