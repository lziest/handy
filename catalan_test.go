package handy

import (
	"math/big"
	"testing"
)

func TestCalatan(t *testing.T) {
	type testcase struct {
		n int
		c *big.Int
	}

	testcases := []testcase{
		{
			n: 0,
			c: big.NewInt(1),
		},
		{
			n: 1,
			c: big.NewInt(1),
		},
		{
			n: 2,
			c: big.NewInt(2),
		},
		{
			n: 3,
			c: big.NewInt(5),
		},
		{
			n: 4,
			c: big.NewInt(14),
		},
		{
			n: 20,
			c: big.NewInt(6564120420),
		},
	}

	for _, c := range testcases {
		r := CatalanNumber(c.n)
		if r.Cmp(c.c) != 0 {
			t.Fatalf("bad Catalan number: %v, got %v, expected %v", c.n, r, c.c)
		}

	}
}
