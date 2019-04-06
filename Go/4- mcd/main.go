package main

import "fmt"

func main() {
	var a, b uint64
	fmt.Print("a: ")
	fmt.Scanln(&a)
	fmt.Print("b: ")
	fmt.Scanln(&b)
	fmt.Print("MCD: ", euclides(a, b))
}

func euclides(a uint64, b uint64) (c uint64) {
	for ; a!=b; {
		if (a>b) { a=a-b
		} else { b=b-a }
		c = a
	}
	return
}