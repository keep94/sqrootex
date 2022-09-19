package main

import (
	"fmt"
	"github.com/keep94/sqroot"
)

func main() {
	n := sqroot.Sqrt(3)
	fmt.Printf("%.50g\n", n)
	fmt.Println(n.Mantissa().FindFirst([]int{0, 5, 0, 8}))
}
