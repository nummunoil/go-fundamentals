package main

import "fmt"

// n := 2
// double(&n)
// n == 4
func double(n *int) {
	*n = *n * 2
}

// name := "Bob"
// appendGreeting(&name)
// name == "Hi, Bob"
func appendGreeting(s *string) {
	*s = fmt.Sprintf("Hi, %s", *s)
}

func main() {
	n := 2
	double(&n)
	fmt.Println(n)

	name := "Bob"
	appendGreeting(&name)
	fmt.Println(name)
}
