package stacks

type Node struct {
	data rune
	next *Node
}

func newNode(data rune) *Node {
	return &Node{data: data}
}

type Stack struct {
	top  *Node
	size int
}

func (s Stack) NewStack() Stack {
	return Stack{}
}

func (s Stack) Len() int {
	return s.size
}

func (s *Stack) Push(data rune) {
	node := newNode(data)
	node.next = s.top
	s.top = node
	s.size++
}

func (s *Stack) Pop() rune {
	if s.top == nil {
		return 0
	}
	node := s.top

	s.top = node.next
	s.size--
	return node.data
}
