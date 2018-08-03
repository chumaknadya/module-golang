package stack

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value int
	next  *Element
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value int) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value int) {
	if s.size > 0 {
		value = s.top.value
		s.top = s.top.next
		s.size--
		return value
	}
	return -1
}
