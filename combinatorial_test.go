package handy

import (
	"math/big"
	"testing"
)

func TestPermuation(t *testing.T) {
	p6_3 := Permutation(6, 3)
	if p6_3.Cmp(big.NewInt(120)) != 0 {
		t.Fatal("Bad Permuataion")
	}
}

func TestFactorial(t *testing.T) {
	f5 := Factorial(5)
	if f5.Cmp(big.NewInt(120)) != 0 {
		t.Fatal("Bad Permuataion")
	}

}
