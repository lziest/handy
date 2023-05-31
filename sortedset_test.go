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

func TestSortedIntListRemoval(t *testing.T) {
	cases := []struct {
		in   []int
		out  []int
		want string
	}{
		{[]int{1}, []int{1}, "{}"},
		{[]int{1, 2, 1}, []int{1, 1}, "{2}"},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4}, "{}"},
		{[]int{3, 2, 1}, []int{4, 5, 6}, "{1 2 3}"},
	}

	for _, c := range cases {
		l := NewSortedList[int]()
		for _, x := range c.in {
			l.Insert(x)
		}
		for _, x := range c.out {
			l.Remove(x)
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

func TestSortedSet(t *testing.T) {
	cases := []struct {
		set_a     []int
		set_b     []int
		union     string
		intersect string
	}{
		{[]int{1}, []int{}, "{1}", "{}"},
		{[]int{1}, []int{2}, "{1 2}", "{}"},
		{[]int{1, 1, 2}, []int{1, 2, 3}, "{1 1 2 3}", "{1 2}"},
	}

	for _, c := range cases {
		sa := NewSortedSet[int](10)
		for _, x := range c.set_a {
			sa.Insert(x)
		}
		sb := NewSortedSet[int](10)
		for _, x := range c.set_b {
			sb.Insert(x)
		}
		u := sa.Union(sb)
		i := sa.Intersect(sb)
		if u.String() != c.union {
			t.Errorf("Want: %s, got: %s", c.union, u)
		}
		if i.String() != c.intersect {
			t.Errorf("Want: %s, got: %s", c.intersect, i)
		}
	}
}
