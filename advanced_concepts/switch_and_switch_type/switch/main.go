package main

import (
	"fmt"
	"time"
)

func whatDoToday(t time.Time) {
	switch t.Weekday() { // break by default
	case time.Saturday:
		fmt.Println("go shopping")
	case time.Sunday:
		fmt.Println("take some rest")
	case time.Monday:
		fallthrough
	case time.Tuesday:
		fmt.Println("go planing")
	default:
		fmt.Println("go working")
	}
}

func whatDoThisDay(t time.Time) {
	fmt.Println(t.Weekday())
	switch { // no condition
	case t.Day() == 1:
		fallthrough
	case t.Day() == 16:
		fmt.Println("Did you buy a lotto?")
	case t.Weekday() == time.Sunday:
		fmt.Println("Netflix day!!")
	}
}

func main() {
	whatDoToday(time.Now())
	whatDoToday(time.Date(2022, time.May, 7, 0, 0, 0, 0, time.UTC))  // Saturday go shopping
	whatDoToday(time.Date(2022, time.May, 8, 0, 0, 0, 0, time.UTC))  // Sunday take some rest
	whatDoToday(time.Date(2022, time.May, 9, 0, 0, 0, 0, time.UTC))  // Monday go planning // fallthrough
	whatDoToday(time.Date(2022, time.May, 10, 0, 0, 0, 0, time.UTC)) // Tuesday go planning
	whatDoToday(time.Date(2022, time.May, 11, 0, 0, 0, 0, time.UTC)) // Wednesday go working

	whatDoThisDay(time.Now())
	whatDoThisDay(time.Date(2022, time.May, 1, 0, 0, 0, 0, time.UTC))  // Sunday 1 // Did you buy a lotto? // fallthrough and first condition
	whatDoThisDay(time.Date(2022, time.May, 16, 0, 0, 0, 0, time.UTC)) // Did you buy a lotto?
	whatDoThisDay(time.Date(2022, time.May, 15, 0, 0, 0, 0, time.UTC)) // Netflix day!!
}
