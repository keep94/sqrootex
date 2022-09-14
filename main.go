package main

import (
	"fmt"
	"math/big"

	"github.com/keep94/sqroot"
)

func main() {
	n := sqroot.Sqrt(big.NewRat(3, 1))
	fmt.Printf("%.50g\n", n)
	fmt.Println(n.Mantissa().FindFirst([]int{0, 5, 0, 8}))
}
