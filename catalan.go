package handy

import (
	"math/big"
)

func CatalanNumber(n int) *big.Int {
	x := int64(n)
	ret := Binomial(2*x, x)
	ret.Div(ret, big.NewInt(x+1))
	return ret
}
