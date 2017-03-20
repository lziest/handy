package handy

type LinkedListNode struct {
	Value interface{}
	Next  *LinkedListNode
}

type LinkedList struct {
	Header LinkedListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Header: LinkedListNode{
			Value: struct{}{},
			Next:  nil,
		},
	}
}

func NewLinkedListNode(x interface{}) *LinkedListNode {
	return &LinkedListNode{
		Value: x,
		Next:  nil,
	}
}

func (l *LinkedList) Insert(x *LinkedListNode) {
	x.Next = l.Header.Next
	l.Header.Next = x
}

func (l *LinkedList) Reverse() {
	curr := l.Header.Next
	l.Header.Next = nil
	for curr != nil {
		next := curr.Next
		l.Insert(curr)
		curr = next
	}
}
