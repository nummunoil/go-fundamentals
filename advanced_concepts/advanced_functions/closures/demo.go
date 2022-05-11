package main

import "fmt"

func main() {
	counterFn := newCounterFunc()
	fmt.Println(counterFn())
	fmt.Println(counterFn())
	fmt.Println(counterFn())
}

func newCounterFunc() func() int {
	var i int // zero value is 0
	return func() int {
		i++
		return i
	}
}
