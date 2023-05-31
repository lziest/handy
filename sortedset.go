package handy

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

type ListNode[T constraints.Ordered] struct {
	Value T
	Next  *ListNode[T]
}

type SortedList[T constraints.Ordered] struct {
	Head *ListNode[T]
	Tail *ListNode[T]
	Size int
}

func NewListNode[T constraints.Ordered](x T) *ListNode[T] {
	return &ListNode[T]{
		Value: x,
		Next:  nil,
	}
}

func NewSortedList[T constraints.Ordered]() *SortedList[T] {
	header := NewListNode[T](*new(T))
	return &SortedList[T]{
		Head: header,
		Tail: header,
		Size: 0,
	}
}

func (s *SortedList[T]) Insert(x T) {
	nn := NewListNode[T](x)
	s.Size += 1
	// Shortcut: if x >= the tail node, insert to the end.
	if s.Tail.Next != nil && x >= s.Tail.Next.Value {
		s.Tail.Next.Next = nn
		s.Tail = s.Tail.Next
		return
	}
	// Insert into the right position by going through
	// the list.
	index := s.Head
	for index.Next != nil {
		if index.Next.Value > x {
			nn.Next = index.Next
			index.Next = nn
			return
		}
		index = index.Next
	}

	// Insertion doesn't happen, so append it to the end.
	index.Next = nn
	s.Tail = index
}

func (s *SortedList[T]) Remove(x T) bool {
	index := s.Head
	for index.Next != nil {
		if index.Next.Value == x {
			found := index.Next
			index.Next = found.Next
			s.Size -= 1
			found = nil
			return true
		}
		index = index.Next
	}
	return false
}

func (s *SortedList[T]) String() string {
	ret := "{"
	index := s.Head
	for index.Next != nil {
		if s, ok := any(index.Next.Value).(string); ok {
			ret += s
		}
		if i, ok := any(index.Next.Value).(int); ok {
			ret += strconv.Itoa(i)
		}
		if f, ok := any(index.Next.Value).(float64); ok {
			ret += strconv.FormatFloat(f, 'f', -1, 64)
		}
		index = index.Next
		if index.Next != nil {
			ret += " "
		}
	}

	ret += "}"
	return ret
}

// SortedList stores generic elements, assumed sorted.
type SortedSet[T constraints.Ordered] struct {
	list     *SortedList[T]
	capacity int
}

func NewOrderedSet[T constraints.Ordered](capacity int) *SortedSet[T] {
	return &SortedSet[T]{
		list:     NewSortedList[T](),
		capacity: capacity,
	}
}

func (s *SortedSet[T]) Insert(x T) bool {
	if s.list.Size >= s.capacity {
		return false
	}
	s.list.Insert(x)
	return true
}

func (s *SortedSet[T]) Remove(x T) bool {
	return s.list.Remove(x)
}

func (s *SortedSet[T]) Union(o *SortedSet[T]) *SortedSet[T] {
	ret := NewOrderedSet[T](s.capacity + o.capacity)
	sIndex := s.list.Head
	oIndex := o.list.Head
	for sIndex.Next != nil || oIndex.Next != nil {
		switch {
		case sIndex.Next == nil:
			{
				ret.Insert(oIndex.Next.Value)
				oIndex = oIndex.Next
				continue
			}
		case oIndex.Next == nil:
			{
				ret.Insert(sIndex.Next.Value)
				sIndex = sIndex.Next
				continue
			}
		case sIndex.Next.Value == oIndex.Next.Value:
			{
				ret.Insert(sIndex.Next.Value)
				sIndex = sIndex.Next
				oIndex = oIndex.Next
				continue
			}
		case sIndex.Next.Value > oIndex.Next.Value:
			{
				ret.Insert(oIndex.Next.Value)
				oIndex = oIndex.Next
				continue
			}
		case sIndex.Next.Value < oIndex.Next.Value:
			{
				ret.Insert(sIndex.Next.Value)
				sIndex = sIndex.Next
				continue
			}
		}

	}
	return ret
}

func (s *SortedSet[T]) Intersect(o *SortedSet[T]) *SortedSet[T] {
	ret := NewOrderedSet[T](s.capacity + o.capacity)
	sIndex := s.list.Head
	oIndex := o.list.Head
	for sIndex.Next != nil || oIndex.Next != nil {
		switch {
		case sIndex.Next == nil:
			{
				oIndex = oIndex.Next
				continue
			}
		case oIndex.Next == nil:
			{
				sIndex = sIndex.Next
				continue
			}
		case sIndex.Next.Value == oIndex.Next.Value:
			{
				ret.Insert(sIndex.Next.Value)
				sIndex = sIndex.Next
				oIndex = oIndex.Next
				continue
			}
		case sIndex.Next.Value > oIndex.Next.Value:
			{
				oIndex = oIndex.Next
				continue
			}
		case sIndex.Next.Value < oIndex.Next.Value:
			{
				sIndex = sIndex.Next
				continue
			}
		}

	}
	return ret
}
