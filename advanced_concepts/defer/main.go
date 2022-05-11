package main

import (
	"fmt"
)

func main() {
	stack()
	doNotForget(4)
	tricky()

}

func stack() {
	defer fmt.Println("derfer#1")
	defer fmt.Println("derfer#2")
	defer fmt.Println("derfer#3") // prints first

	fmt.Println("end")
}

func doNotForget(n int) {
	defer fmt.Println(n) // n = 4
	defer func() {
		fmt.Println(n)
	}() // n = 8
	n += n
}

func tricky() {
	for i := 0; i < 3; i++ {
		defer func() { fmt.Println(i) }() // i always equals 3
	}
}
