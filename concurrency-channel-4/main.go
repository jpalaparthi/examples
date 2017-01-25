package main

import (
	"fmt"
	"sync"
	"time"
)

var e, o int

func main() {
	var c sync.Mutex

	go func() {

		for i := 0; i < 500; i++ {
			if i%2 == 0 {
				c.Lock()
				e = i
				c.Unlock()
				time.Sleep(time.Millisecond * 10)
			} else {
				c.Lock()
				e = i
				c.Unlock()
			}
		}
	}()

	go func() {
		for {
			c.Lock()
			fmt.Println(e)
			fmt.Println(o)
			c.Unlock()

		}
	}()

	_, _ = fmt.Scanln()

}
