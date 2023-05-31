package handy

import (
	"testing"
)

func TestSortedIntList(t *testing.T) {
	cases := []struct {
		in   []int
		want string
	}{
		{[]int{1}, "{1}"},
		{[]int{1, 2, 1}, "{1 1 2}"},
		{[]int{1, 2, 3}, "{1 2 3}"},
		{[]int{3, 2, 1}, "{1 2 3}"},
	}

	for _, c := range cases {
		l := NewSortedList[int]()
		for _, x := range c.in {
			l.Insert(x)
		}
		got := l.String()
		if got != c.want {
			t.Errorf("Want: %s, got: %s", c.want, got)
		}
	}
}

func TestSortedStringList(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"ab"}, `{ab}`},
		{[]string{"ab", "cd"}, `{ab cd}`},
		{[]string{"ab", "aa"}, `{aa ab}`},
	}

	for _, c := range cases {
		l := NewSortedList[string]()
		for _, x := range c.in {
			l.Insert(x)
		}
		got := l.String()
		if got != c.want {
			t.Errorf("Want: %s, got: %s", c.want, got)
		}
	}
}
