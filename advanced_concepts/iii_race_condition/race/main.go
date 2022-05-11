package main

import (
	"fmt"
)

var i int

func main() {
	go func() {
		for {
			fmt.Print(read())
		}
	}()

	for i := 0; i < 10; i++ {
		write(i)
	}
}

func write(n int) {
	i = n
}

func read() int {
	return i
}
