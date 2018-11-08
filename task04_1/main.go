package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	f := func() {
		defer fmt.Println("World")
		fmt.Println("Hello")
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(f)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Scanln()
}
