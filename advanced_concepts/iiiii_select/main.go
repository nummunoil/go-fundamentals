package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	qCh := make(chan struct{}) // q channel should be type of struct{} or type of [0]func

	go fibonacci(ch, qCh)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	qCh <- struct{}{}
}

func fibonacci(ch chan int, qCh chan struct{}) {
	a, b := 0, 1
	for {
		select {
		case ch <- a:
			a, b = b, a+b
		case <-qCh:
			return
		}
	}
}
