package main

import "fmt"
func main() {
	i, j := 21, 2210

	p := &i
	fmt.Println(*p)

	*p = 66
	fmt.Println(i)

	p = &j
	*p = *p / 12

	fmt.Println(j)
}
