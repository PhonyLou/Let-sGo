package main

import "fmt"

type Vertex struct {
	X int
	Y int
}



func main() {
	ver := Vertex{1, 2}
	fmt.Println(ver.X)

	ver.X = 66
	fmt.Println(ver.X)
}
