package main

import "fmt"

// greeting("Pallat") should returns "Hello, Pallat"
func greeting(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

// greetingWithAge("Pallat", 30) should returns "Hello, Pallat. You are 30 years old."
func greetingWithAge(name string, age uint) string {
	return fmt.Sprintf("Hello, %s. You are %d years old.", name, age)
}

// greetingWithAgeAndDrink("Pallat", 30, "Cola") should returns "Hello, Pallat. You are 30 years old and your favorite drink is Cola."
func greetingWithAgeAndDrink(name string, age uint, drink string) string {
	return fmt.Sprintf("Hello, %s. You are %d years old and your favorite drink is %s.", name,
		age, drink)
}

func main() {
	fmt.Println(greeting("Pallat"))
	fmt.Println(greetingWithAge("Pallat", 30))
	fmt.Println(greetingWithAgeAndDrink("Pallat", 30, "Cola"))
}
