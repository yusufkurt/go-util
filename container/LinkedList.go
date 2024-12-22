package container

type Node[T any] struct {
	Next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	head    *Node[T]
	tail    *Node[T]
	current *Node[T]
}

func (l *LinkedList[T]) Add(item T) {
	newItem := &Node[T]{Value: item}
	if l.head == nil {
		l.head = newItem
		l.tail = newItem
		l.current = newItem
		return
	}

	l.tail.Next = newItem
	l.tail = newItem
}

func (l *LinkedList[T]) Next() bool {
	return l.current != nil
}

func (l *LinkedList[T]) Get() *T {
	v := &l.current.Value
	l.current = l.current.Next
	return v
}
