package handy

// DisjointSet represents disjoint set.
type DisjointSet struct {
	parent []int
	size   int
}

// NewDisjointSet returns a new disjoint set.
func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DisjointSet{
		parent: parent,
		size:   n,
	}
}

// Size returns the size of DisjointSet.
func (s *DisjointSet) Size() int {
	return s.size
}

// Find returns the root of x.
func (s *DisjointSet) Find(x int) int {
	if s.parent[x] == x {
		return x
	}
	s.parent[x] = s.Find(s.parent[x])
	return s.parent[x]
}

// Merge merges two set that contains x and y respectively.
func (s *DisjointSet) Merge(x, y int) {
	px := s.Find(x)
	py := s.Find(y)

	if px < py {
		s.parent[py] = px
	} else {
		s.parent[px] = py
	}
}
