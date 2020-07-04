package mystack

import "sync"

type MyStack struct {
	elems []interface{}
	top   int
	lock  sync.Mutex
}

func NewStack() *MyStack {
	return &MyStack{
		elems: make([]interface{}, 2),
		top:   -1,
	}
}

// 进栈
func (s *MyStack) Push(elem interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.top++
	if len(s.elems) < s.top*2 {
		s.extendElems()
	}
	s.elems[s.top] = elem
}

// 出栈
func (s *MyStack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.top >= 0 {
		elem := s.elems[s.top]
		s.elems[s.top] = nil
		s.top--
		return elem
	}
	return nil
}

// 扩容
func (s *MyStack) extendElems() {
	// s.lock.Lock()
	// defer s.lock.Unlock()

	originElems := s.elems
	s.elems = make([]interface{}, s.top*2)
	for idx := range originElems {
		s.elems[idx] = originElems[idx]
	}
}
