package stack

import "container/list"

// Stack is LIFO Data Structure
type Stack struct {
	v *list.List
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	result := s.v.Back()

	if result != nil {
		return s.v.Remove(result)
	}

	return nil
}

func New() *Stack {
	return &Stack{list.New()}
}
