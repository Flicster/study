package algorithms

import (
	"encoding/json"
)

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func NewLinkedListFromArray[T any](data []T) *LinkedList[T] {
	ll := &LinkedList[T]{}
	for _, d := range data {
		ll.Append(d)
	}
	return ll
}

func (ll *LinkedList[T]) Append(data T) {
	newNode := &Node[T]{
		data: data,
	}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList[T]) AsSlice() []T {
	data := make([]T, 0)
	current := ll.head
	if current == nil {
		return data
	}
	data = append(data, current.data)
	for current.next != nil {
		current = current.next
		data = append(data, current.data)
	}
	return data
}

func (ll *LinkedList[T]) MarshalJSON() ([]byte, error) {
	data := ll.AsSlice()
	return json.Marshal(data)
}

func (ll *LinkedList[T]) UnmarshalJSON(data []byte) error {
	var dataSlice []T
	err := json.Unmarshal(data, &dataSlice)
	if err != nil {
		return err
	}
	*ll = *NewLinkedListFromArray(dataSlice)
	return nil
}
