package handy

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	l := NewLinkedList()
	if l == nil {
		t.Fatal("bad NewLinkedList")
	}
}

func TestLinkedListInsert(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		n := NewLinkedListNode(i)
		l.Insert(n)
	}

	next := l.Header.Next
	for i := 0; i < 10; i++ {
		if next.Value != 10-i-1 {
			t.Fatal("Bad Insert")
		}
		next = next.Next
	}
}

func TestLinkedListReverse(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		n := NewLinkedListNode(i)
		l.Insert(n)
	}
	l.Reverse()

	next := l.Header.Next
	for i := 0; i < 10; i++ {
		if next.Value != i {
			t.Fatal("Bad Insert")
		}
		next = next.Next
	}
}
