// concurrency-channel-2 project main.go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	for i := range ch { // receiver
		fmt.Println(i)
	}
}

//the sending should close the channel or else there would be a deadlock by the receiver in the range
