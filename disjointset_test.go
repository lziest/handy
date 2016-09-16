package handy

import (
	"testing"
)

func TestNewDisjointSet(t *testing.T) {
	n := 100
	s := NewDisjointSet(n)
	if s.Size() != 100 {
		t.Fatal("bad NewDisjointSet")
	}

	for i := 0; i < n; i++ {
		if s.Find(i) != i {
			t.Fatal("bad NewDisjointSet")
		}
	}
}

func TestDisjointSetFindAndUnion(t *testing.T) {
	n := 100
	s := NewDisjointSet(n)
	// Use disjoint set as sieve to find primes
	for i := 0; i < n; i++ {
		if i > 1 {
			if s.Find(i) < i {
				continue
			}
			m := i
			for m < n {
				pm := s.Find(m)
				if pm == m {
					// m is either a prime, or it should be merged with i
					s.Merge(i, m)
				}
				t.Log(m, s.Find(m))
				m += i
			}
		}
	}

	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23,
		29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97,
	}

	for _, p := range primes {
		if s.Find(p) != p {
			t.Fatalf("Bad DisjointSet: failed to find prime %d", p)
		}
	}
}
