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
func (s *StackElem) Pup(elem interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.top++
	s.elems = append(s.elems, elem)
}

// 出栈
func (s *StackElem) Pop() interface{} {
	s.lock
	if s.top >= 0 {
		var elem interface{}
		copy(elem, s.elems[s.top])
		s.top--
		return elem
	}
	return nil
	defer s.lock.Unlock()
}
