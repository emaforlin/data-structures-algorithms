package DS

import "fmt"

/*
c: Cell
L: list struct
e: element
------------------------
/*
TODO: Implement Simple Linked List Methods

[x] NewSList(): Creates a new empty linked list.
[x] Len(): Returns the number of elements in the list.
[x] Prepend(value): Adds a new element at the beginning of the list.
[x] Add(value): Adds a new element to the end of the list.
[x] Get(index): Retrieves an element at a given index in the list.
[x] Insert(index, value): Inserts a value at a specific position in the list.

[ ] Remove(value): Removes the first occurrence of a value from the list.

[ ] RemoveAt(index): Removes an element at a specific index and returns it.

[ ] Find(value): Searches for a value and returns its index if found.


[ ] IsEmpty(): Checks if the list is empty.

[ ] Print()
    Prints the elements of the list for debugging.
*/

type SList[T any] interface {
	Print()
	Len() int
	Get(index int) (T, error)
	Add(value T)
	Prepend(value T)
	Insert(index int, value T) error
}

type cell[T any] struct {
	data *T
	next *cell[T]
}

type SListImpl[T any] struct {
	head *cell[T]
	size int
}

func NewSList[T any]() SList[T] {
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

func (L *SListImpl[T]) Get(index int) (T, error) {
	var zeroValue T

	if index < 0 || index > (L.size-1) {
		return zeroValue, ErrIndexOutOfBound
	}

	var c = L.head
	for i := 0; i < index; i++ {
		c = c.next
	}

	return *c.data, nil
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
