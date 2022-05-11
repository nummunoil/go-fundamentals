package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(3)

	start := time.Now()
	go doSomeThing(wg)
	go doSomeThing(wg)
	go doSomeThing(wg)

	wg.Wait()
	fmt.Println(time.Since(start))

}

func doSomeThing(wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("do something")
	wg.Done()
}
