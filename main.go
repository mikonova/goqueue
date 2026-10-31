package goqueue

import (
	"sync"
)

// A closed type for LIFO/FIFO queues
type Queue[T any] struct {
	sl       []T
	mut      sync.RWMutex
	cap      int
	cond     sync.Cond
	lElemPtr int
}

// Creates a new queue of type T and with a max capacity of cap. Panics if capacity is 0 or less.
// To use it like an unbuffered channel just use cap = 1
func New[T any](cap int) *Queue[T] {
	q := Queue[T]{lElemPtr: -1}
	q.cap = cap
	if cap > 0 {
		q.sl = make([]T, cap)
	} else {
		panic("[err] capacity must be > 0")
	}
	q.cond = *sync.NewCond(&q.mut)
	return &q
}

// Push element into the queue (this will be the last element).
func (q *Queue[T]) Push(element T) {
	q.mut.Lock()
	defer q.mut.Unlock()
	defer q.cond.Signal()
	for q.lElemPtr >= q.cap-1 {
		q.cond.Wait()
	}
	q.lElemPtr++
	q.sl[q.lElemPtr] = element
}

// Pop the last element from the queue. This also removes the last element
func (q *Queue[T]) Pop() T {
	q.mut.Lock()
	defer q.mut.Unlock()
	defer q.cond.Signal()
	for q.lElemPtr < 0 {
		q.cond.Wait()
	}
	elem := q.sl[q.lElemPtr]
	q.lElemPtr--
	return elem
}

// Pull the first element from the queue. This also removes the first element
func (q *Queue[T]) Pull() T {
	q.mut.Lock()
	defer q.mut.Unlock()
	defer q.cond.Signal()
	for q.lElemPtr < 0 {
		q.cond.Wait()
	}
	elem := q.sl[0]
	for k := range len(q.sl) - 2 {
		q.sl[k] = q.sl[k+1]
	}
	q.lElemPtr--
	return elem

}

// Close the queue. The queue can no longer be used and any usage will lead to a runtime panic
func (q *Queue[T]) Close() {
	q = nil
}

// Set the pointer to the first element and make it elem which virtually renders all values owerritable and inaccessible
func (q *Queue[T]) Drain(elem T) {
	q.mut.Lock()
	defer q.mut.Unlock()
	defer q.cond.Signal()
	q.lElemPtr = 0
	q.sl[q.lElemPtr] = elem
}
