package datastruct

import (
	"fmt"
)

// doubleNode is a struct that provide a node that can be concatenated to create a double linked list.
// It has a value of type any and a pointer to the next node and one to the previous.
type doubleNode struct {
	val  any
	next *doubleNode
	prev *doubleNode
}

// doubleLinkedList is a list that is double linked.
type DoubleLinkedList struct {
	head *doubleNode
	tail *doubleNode
	size int
}

// GetHead returns tail's value.
// Panics if the list is empty.
func (l *DoubleLinkedList) GetHead() any {
	return l.head.val
}

// GetTail returns tail's value.
// Panics if the list is empty.
func (l *DoubleLinkedList) GetTail() any {
	return l.tail.val
}

// GetElement returns the element at given index.
// Panics if index is out of bounds.
func (l *DoubleLinkedList) GetElement(index int) (val any) {
	if index >= l.size {
		panic("index out of bounds")
	}
	var node *doubleNode
	if index < l.size/2 {
		node = l.head
		for i := 0; i < index; i++ {
			node = node.next
		}

	} else {
		node = l.tail
		for i := l.size - 1; i > index; i-- {
			node = node.prev
		}
	}
	return node.val
}

// AddAtStart add the element passed as input at the start of the list
func (l *DoubleLinkedList) AddAtStart(e any) {
	newNode := &doubleNode{e, nil, nil}
	if l.head == nil {
		l.head, l.tail = newNode, newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
	l.size++
}

// AddAtEnd add the element passed as input at the start of the list
func (l *DoubleLinkedList) AddAtEnd(e any) {
	newNode := &doubleNode{e, nil, nil}
	if l.tail == nil {
		l.head, l.tail = newNode, newNode
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

// RemoveHead remove the head of the list
func (l *DoubleLinkedList) RemoveHead() {
	if l.size < 2 {
		l.head, l.tail = nil, nil
	} else {
		l.head = l.head.next
	}
}

// RemoveHead remove the tail of the list
func (l *DoubleLinkedList) RemoveTail() {
	if l.size < 2 {
		l.head, l.tail = nil, nil
	} else {
		l.tail = l.tail.prev
	}
}

// Size returns the len of the list
func (l *DoubleLinkedList) Size() int {
	return l.size
}

// String returns a string representation of the list
func (l *DoubleLinkedList) String() (s string) {
	node := l.head
	for node != nil {
		s += fmt.Sprint(node.val)
		if node.next != nil {
			s += " "
		}
		node = node.next
	}

	return
}
