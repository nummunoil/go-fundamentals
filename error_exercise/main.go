package main

import (
	"errors"
	"fmt"
)

// validateLength return false when string length less than 8
// Please change return type to error with message "too short string"
func validateLength(s string) error {
	if len([]rune(s)) < 8 {
		return errors.New("too short string")
	}
	return nil
}

type ErrTooYoung int

func (err ErrTooYoung) Error() string {
	return fmt.Sprintf("%d is too young", err)
}

var ageError = errors.New("your age is negative!")

// in case of too young age please create a new type ErrTooYoung int` with method `Error() string`
// example error message: "17 is too young"
func validateAge(n int) error {
	if n < 0 {
		return ageError
	}
	if n < 18 {
		return ErrTooYoung(n)
	}
	return nil
}

func main() {
	fmt.Println(validateLength("hello"))
	fmt.Println(validateLength("123456789"))
	fmt.Println(validateLength("12345678"))
	fmt.Println(validateAge(17))
	fmt.Println(validateAge(18))
	fmt.Println(validateAge(-1))
}
