package handy

import (
	"math/big"
)

// CatalanNumber returns the n-th Catalan Number.
// CatalanNumber answers many combinatorial questions, such as:
// 1. given nxn grid, what is the number of different paths along
// edges from bottom left to top right corner,
// withouth crossing the diagonal?
// 2. the number of expressions having n pairs of parentheses
// which are correctly matched.
// 3. the number of triangualtion of n+2 contex polygon.
func CatalanNumber(n int) *big.Int {
	x := int64(n)
	ret := Binomial(2*x, x)
	ret.Div(ret, big.NewInt(x+1))
	return ret
}
