package datastructs

// Queue is an interface for queue data types.
type Queue interface {
	IsEmpty() bool
	Enqueue(any)
	Dequeue() any
	First() any
	Size() any
}

// DoubleLinkeListQueue is a Queue based on double linked lists
type DoubleLinkedListQueue struct {
	DoubleLinkedList
}

// IsEmpty returns true if the queue is empty, false otherwise.
func (q *DoubleLinkedListQueue) IsEmpty() bool {
	return q.Size() == 0
}

// Enqueue enqueue an element passed as input in the queue.
func (q *DoubleLinkedListQueue) Enqueue(val any) {
	q.AddAtEnd(val)
}

// Dequeue returns the first element of the queue and removes it from the queue.
func (q *DoubleLinkedListQueue) Dequeue() (res any) {
	res = q.GetHead()
	q.RemoveHead()
	return
}

// First reads the first element of the queue.
func (q *DoubleLinkedListQueue) First() any {
	return q.GetHead()
}
