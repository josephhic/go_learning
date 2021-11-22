package main

import "fmt"

func main() {
	a := Vector{1, 5, 3}
	b := Vector{2, 6, -1}

	d := a.dot(&b)
	fmt.Println(d)
}
