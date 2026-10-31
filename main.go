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
	q := Queue[T]{}
	q.cap = cap
	if cap > 0 {
		q.sl = make([]T, cap)
	} else {
		panic("[err] capacity must be > 0")
	}
	q.cond = *sync.NewCond(&q.mut)
	return &q
}

// Push element into the queue
// panics if types dont match
func (q *Queue[T]) Push(element T) {
	q.mut.Lock()
	defer q.mut.Unlock()
	for q.lElemPtr == q.cap {
		q.cond.Wait()
	}
	q.sl[q.lElemPtr] = element
	q.lElemPtr++
}

func (q *Queue[T]) Pop() any {
	q.mut.Lock()
	defer q.mut.Unlock()
	defer q.cond.Signal()
	if q.lElemPtr == 0 {
		return nil
	}
	q.lElemPtr--
	return q.sl[q.lElemPtr+1]
}

func (q *Queue[T]) Pull() any {
	q.mut.Lock()
	defer q.mut.Unlock()
	defer q.cond.Signal()
	if q.lElemPtr == 0 {
		return nil
	}
	a := q.sl[0]
	for k := range len(q.sl) - 1 {
		q.sl[k] = q.sl[k+1]
	}
	q.lElemPtr--
	return a

}

func (q *Queue[T]) Close() {
	q = nil
}
