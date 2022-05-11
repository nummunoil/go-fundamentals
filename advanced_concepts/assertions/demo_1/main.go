package main

import "fmt"

func main() {
	var a interface{}

	a = 10
	fmt.Printf("%T %v\n", a, a)

	var i int
	// i = a // can't do this
	i = a.(int)
	fmt.Printf("%T %v\n", i, i)

	// ------or------

	if i2, ok := a.(int); ok {
		fmt.Printf("%T %v\n", i2, i2)
	}
}
