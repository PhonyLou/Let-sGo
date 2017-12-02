package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint = 1 << 64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const temple = "%T(%v)\n"
	fmt.Printf(temple, ToBe, ToBe)
	fmt.Printf(temple, MaxInt, MaxInt)
	fmt.Printf(temple, z, z)
}
