package list

import (
	"errors"
	"sync"

	"github.com/gohiweeds/go-datastructure/common"
)

type List struct {
	common.Operater
	head    *node
	size    uint64
	elemnum uint64
	m       *sync.Mutex
}

//Double Linked List
type node struct {
	prev *node
	next *node
	data interface{}
}

func NewList(size uint64) *List {
	return &List{
		head:    new(node),
		size:    size,
		elemnum: 0,
		m:       new(sync.Mutex),
	}
}

func (l *List) Add(value interface{}) (bool, error) {
	l.m.Lock()
	defer l.m.Unlock()

	if l.elemnum >= l.size {
		return false, errors.New("Linked List is Full")
	}
	n := new(node)
	n.data = value

	next := l.head
	for next.next != nil {
		next = next.next
	}
	next.next = n
	n.prev = next
	n.next = nil

	l.elemnum++

	return true, nil
}

func (l *List) Delete(value interface{}) (bool, error) {
	l.m.Lock()
	defer l.m.Unlock()

	if l.elemnum == 0 {
		return false, errors.New("Linked List is Empty")
	}

	next := l.head
	for next.next != nil {
		if next.data == value {
			next.next.prev = next.prev
			next.prev.next = next.next
			l.elemnum--
			return true, nil
		}
		next = next.next
	}
	return false, errors.New("No such value exist")
}

func (l *List) Get(value interface{}) (bool, error) {
	l.m.Lock()
	defer l.m.Unlock()

	if l.elemnum == 0 {
		return false, errors.New("Linked List is Empty")
	}

	next := l.head
	for next.next != nil {
		if next.data == value {
			return true, nil
		}
		next = next.next
	}
	return false, errors.New("Not exist such value")
}

func (l *List) Modify(value, update interface{}) (bool, error) {
	l.m.Lock()
	defer l.m.Unlock()

	if l.elemnum == 0 {
		return false, errors.New("Linked List is Empty")
	}

	next := l.head
	for next.next != nil {
		if next.data == value {
			return true, nil
		}
		next = next.next
	}
	return false, errors.New("Not exist such value")
}

func (l *List) Len() uint64 {
	return l.elemnum
}
