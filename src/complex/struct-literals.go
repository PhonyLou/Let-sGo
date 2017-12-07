package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{Y: 66, X: 1}
	v3 = Vertex{}
	p = &Vertex{1, 3}
)

func main() {
	fmt.Println(v1, v2, v3, *p)
}
