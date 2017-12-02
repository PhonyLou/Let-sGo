package main

import (
	"fmt"
	"math"
)

func main() {
	var i, j int = 3, 4
	d := math.Sqrt(float64(i + j))
	var e uint = uint(d)
	fmt.Println(i, j, d, e)

}
