package my_queue

import "sync"

// 定义一个队列
type MyQueue struct {
	elems []interface{}
	front,tail int
	lock sync.Mutex
}

func NewQueue() *MyQueue {
	return &MyQueue{
		elems: make([]interface{},2),
		front: -1, // front 指向第一个元素
		tail: -1, // tail指向最后一个元素的后面一个元素
	}
}

// 进队列
func (q *MyQueue) In(elem interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.front == -1 {
		// 队列为空
		q.elems = make([]interface{}, 2)
		q.front = 0
		q.tail = 1
		q.elems[q.front] = elem
	}else{
		if len(q.elems) < q.tail+1 {
			// 需要扩容
			q.extendElems()
		} 
		q.elems[q.tail] = elem
		q.tail += 1
	}
}

// 出队列
func (q *MyQueue) Out() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	// 队列为空
	if q.front < 0 {
		return nil
	}

	if q.front < q.tail {
		// 队列不为空
		elem := q.elems[q.front]
		q.front += 1
		// 检查队列是否为空
		q.check()
		return elem
	}else{
		return nil
	}
}

func (q *MyQueue) Len() int {
	return q.tail-q.front
}

// 扩容队列
func (q *MyQueue) extendElems(){
	elems := make([]interface{}, q.tail*2)
	idx := 0
	for i := q.front; i < q.tail; i++ {
		elems[idx] = q.elems[i]
		idx += 1
	}
	q.front = 0
	q.tail = idx
	q.elems = elems
}

// 检查队列是否为空
func (q *MyQueue) check() {
	if q.front >= q.tail {
		// 队列为空
		q.front = -1
		q.tail = -1
		q.elems = make([]interface{}, 2)
		return
	}

	// 如果队列利用率小于1/3则优化
	if (q.tail - q.front ) * 3 < len(q.elems) {
		// 缩容
		q.subElems()
	}
	return
}

// 缩容
func (q *MyQueue) subElems() {
	elems := make([]interface{}, (q.tail - q.front) * 2)
	idx := 0 
	for i := q.front; i<q.tail; i++ {
		elems[idx] = q.elems[i]
		idx += 1
	}
	q.front = 0 
	q.tail = idx
	q.elems = elems
}