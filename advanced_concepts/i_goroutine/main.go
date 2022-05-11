package main

import (
	"fmt"
	"time"
)

func main() {
	// start := time.Now()
	// fmt.Println("start")
	// go doSomeThing()
	// go doSomeThing()
	// go doSomeThing()
	// time.Sleep(time.Second)
	// fmt.Println("end")
	// fmt.Println(time.Since(start))

	example()
}

func doSomeThing() {
	time.Sleep(time.Second)
	fmt.Println("do something")
}

func example() {
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Print("-")
		}
	}()
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Print("+")
		}
	}()
	time.Sleep(time.Second)
}
