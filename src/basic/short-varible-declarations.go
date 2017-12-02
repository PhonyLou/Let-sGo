package main

import "fmt"

func main() {
	var a, b int = 1, 2
	keyOne := false
	keyTwo := "hello"
	keyOne = true
	fmt.Println(a, b, keyOne, keyTwo)
}
