package main

import (
	"fmt"
	"math/big"

	"github.com/keep94/sqroot"
)

func main() {
	fmt.Printf("%.50g\n", sqroot.Sqrt(big.NewRat(3, 1)))
}
