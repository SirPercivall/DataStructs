package datastructs

// Stack is an interface that provides a stack
type Stack interface {
	IsEmpty() bool
	Top() any
	Pop() any
	Push(any)
	Size() int
}

// LinkedLinkStack is an implementation of Stack based on LinkedList.
type LinkedListStack struct {
	LinkedList
}

// IsEmpty returns true if stack is empty false otherwise.
func (s *LinkedListStack) IsEmpty() bool {
	return s.Size() == 0
}

// Top returns the element at the top of the stack
func (s *LinkedListStack) Top() any {
	return s.GetElement(0)
}

// Pop returns the element at the top of the stack and removes it.
func (s *LinkedListStack) Pop() (res any) {
	res = s.GetElement(0)
	s.RemoveHead()
	return
}

// Push add the parameter val at the top of the stack.
func (s *LinkedListStack) Push(val any) {
	s.AddAtStart(val)
}
