package main

import "fmt"

type float = float64

const (
	name = iota
	age
	_
	age3
)

func main() {
	fmt.Printf("%T %.2f\n", 5/2.3, 5/2.3)
}
