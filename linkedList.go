package datastruct

import (
	"fmt"
)

// node is a struct that is used to create a linkedList.
// It has a value of any type and a pointer to the next node.
type node struct {
	val  any
	next *node
}

// LinkedList is a struct that provides a list based on linked nodes.
// head is the first node of the list
// size is the length of the list
type LinkedList struct {
	head *node
	size int
}

// GetElement returns the element at given index.
// It panics if index is out of bounds
func (l *LinkedList) GetElement(index int) (val any) {
	if index >= l.size {
		panic("index out of bounds")
	}

	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.val
}

// AddElement add the element passed as input at the start of the list
func (l *LinkedList) AddAtStart(e any) {
	newNode := &node{e, l.head}
	l.head = newNode
	l.size++
}

// RemoveHead remove the head of the list
func (l *LinkedList) RemoveHead() {
	// one element case
	if l.head != nil {
		l.head = l.head.next
	}
}

// Size returns the len of the list
func (l *LinkedList) Size() int {
	return l.size
}

// String returns a string representation of the list
func (l *LinkedList) String() (s string) {
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
