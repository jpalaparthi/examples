// concurrency-1 project main.go
package main

import (
	"fmt"
)

func main() {
	i := new(int) // new always returns the pointer
	defer func() {
		for j := 0; j < 100; j++ {
			fmt.Println(j + *i)
		}
	}()
	defer fmt.Println("this defer prints second", i)
	defer fmt.Println("this defer function will be printed first")
	fmt.Println(*i) // since printing i is deferred , it prints the value..
	fmt.Println("Checking when defer really work")
}

// defers always work in the Fist in last out manner..  that means first defer will be executed at last.
