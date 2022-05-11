package main

import "fmt"

func whichType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("this is a string %s\n", v)
	case int:
		fmt.Printf("this is an integer %d\n", v)
	case float64:
		fmt.Printf("this is a floating point %f\n", v)
	default:
		fmt.Println("don't know")
	}
}

func main() {
	whichType("hello")
	whichType(10)
	whichType(3.14)
	whichType(true)
}
