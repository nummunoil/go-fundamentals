package main

import "fmt"

func isOdd(n int) bool {
	if n%2 != 0 {
		return true
	} else {
		return false
	}
}

func isOdd2(n int) bool {
	if result := n%2 == 0; result {
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Println(isOdd(1))
	fmt.Println(isOdd(2))
	fmt.Println(isOdd2(1))
	fmt.Println(isOdd2(2))
}
