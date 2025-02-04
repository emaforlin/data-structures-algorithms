package DS

import (
	"fmt"
)

type SList[T comparable] interface {
	Print()
	Len() int
	Get(index int) (*T, error)
	Add(value T)
	Prepend(value T)
	Insert(index int, value T) error
	Find(value T) (*T, error)
	Remove(value T) error
	RemoveAt(index int) error
}

type cell[T comparable] struct {
	data *T
	next *cell[T]
}

type SListImpl[T comparable] struct {
	head *cell[T]
	size int
}

func indexOf[T comparable](node *cell[T], value T) (int, error) {
	indx := indexOfR(node, value, 0)
	if indx == -1 {
		return indx, ErrValueNotFound
	}
	return indx, nil
}

func indexOfR[T comparable](c *cell[T], value T, index int) int {
	if c == nil {
		return -1
	}
	if *c.data == value {
		return index
	}
	return indexOfR(c.next, value, index+1)
}

func (L *SListImpl[T]) advance(index int) *cell[T] {
	c := L.head
	for i := 0; i < index && c.next != nil; i++ {
		c = c.next
	}
	if c != nil {
		return c
	}
	return nil
}

func NewSList[T comparable]() SList[T] {
	return &SListImpl[T]{
		head: nil,
		size: 0,
	}
}

func (L *SListImpl[T]) Print() {
	fmt.Print("[ ")
	for c := L.head; c != nil; c = c.next {
		fmt.Printf("%#v ", *c.data)
	}
	fmt.Print("]")
}

func (L *SListImpl[T]) Len() int {
	return L.size
}

func (L *SListImpl[T]) Prepend(value T) {
	newHead := &cell[T]{
		data: &value,
		next: L.head,
	}
	L.head = newHead
	L.size++
}

func (L *SListImpl[T]) Add(value T) {
	newCell := &cell[T]{data: &value, next: nil}
	if L.head == nil {
		L.head = newCell
	} else {
		c := L.head
		for c.next != nil {
			c = c.next
		}
		c.next = newCell
	}
	L.size++
}

func (L *SListImpl[T]) Get(index int) (*T, error) {
	if index < 0 || index > (L.size-1) {
		return nil, ErrIndexOutOfBound
	}

	var c = L.head
	for i := 0; i < index; i++ {
		c = c.next
	}

	return c.data, nil
}

func (L *SListImpl[T]) Insert(index int, value T) error {
	if index < 0 || index > L.size {
		return ErrIndexOutOfBound
	}
	newCell := &cell[T]{
		data: &value,
		next: nil,
	}
	if index == 0 {
		newCell.next = L.head
		L.head = newCell
	} else {
		c := L.head
		for i := 0; i < index-1 && c.next != nil; i++ {
			c = c.next
		}
		newCell.next = c.next
		c.next = newCell
	}
	L.size++
	return nil
}

func (L *SListImpl[T]) Find(value T) (*T, error) {
	index, err := indexOf(L.head, value)
	if err != nil {
		return nil, err
	}
	found, err := L.Get(index)
	if err != nil {
		return nil, err
	}
	return found, nil
}

func (L *SListImpl[T]) RemoveAt(index int) error {
	if index < 0 || index > L.size {
		return ErrIndexOutOfBound
	}
	if index == 0 {
		L.head = L.head.next
	} else {
		cell := L.advance(index - 1)
		if cell == nil {
			return ErrValueNotFound
		}

		next := cell.next
		cell.next = next.next
	}

	L.size--
	return nil
}

func (L *SListImpl[T]) Remove(value T) error {
	index, err := indexOf(L.head, value)
	if err != nil {
		return err
	}
	return L.RemoveAt(index)
}
