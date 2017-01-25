// concurrency-channel-3 project main.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 6)
	wg := sync.WaitGroup{}
	wg.Add(1) //wg.Add(2)
	go func() {
		for i := 0; i < 500; i++ {
			//if i == 499 {
			//		close(ch)
			//	}
			if i%5 == 0 {
				ch <- i
			}
		}
		wg.Done()
	}()
	go func() {
		for k := range ch {
			fmt.Println(k)
		}
		wg.Done()

	}()
	wg.Wait()
}
