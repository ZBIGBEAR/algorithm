package mystack

import "sync"

type StackElem struct {
	elems []interface{}
	top   int
	lock  sync.Mutex
}

func NewStack() *StackElem {
	return &StackElem{
		elems: make([]interface{}, 0),
		top:   -1,
	}
}

// 进栈
func (s *StackElem) Push(elem interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.top++
	s.elems = append(s.elems, elem)
}

// 出栈
func (s *StackElem) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.top >= 0 {
		elem := s.elems[s.top]
		s.top--
		return elem
	}
	return nil
}
