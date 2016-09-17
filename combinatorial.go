package handy

import (
	"math/big"
)

func Factorial(n int64) *big.Int {
	ret := big.NewInt(0)
	ret.MulRange(1, n)
	return ret
}

func Binomial(n, k int64) *big.Int {
	ret := big.NewInt(0)
	ret.Binomial(n, k)
	return ret
}

func Permutation(n, k int64) *big.Int {
	ret := big.NewInt(0)
	ret.MulRange(n-k+1, n)
	return ret
}
