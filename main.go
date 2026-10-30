package goqueue

import (
	"errors"
	"reflect"
	"strings"
	"sync"
)

type queue[T any] struct {
	sl   []T
	mut  sync.RWMutex
	cap  int
	cond sync.Cond
}

func New[T any](capacity int) *queue[T] {
	q := queue[T]{
		sl:  make([]T, capacity),
		cap: capacity,
	}
	q.cond = *sync.NewCond(&q.mut)
	return &q
}

func (q *queue[any]) Push(element any) {
	q.mut.Lock()
	defer q.mut.Unlock()
	if len(q.sl) > q.cap {

	}
	elemType := reflect.TypeOf(element)
	sliceType := reflect.TypeOf(q.sl).Elem()
	if elemType != sliceType {
		builder := strings.Builder{}
		builder.WriteString("[err] invalid types: ")
		builder.WriteString(elemType.Name())
		builder.WriteString(" and ")
		builder.WriteString(sliceType.Name())
		panic(errors.New(builder.String()))
	}
}

func (q *queue[T]) Pop() any {
	q.mut.Lock()
	defer q.mut.Unlock()
}
