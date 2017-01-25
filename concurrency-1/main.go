// concurrency-1 project main.go
package main

import (
	"fmt"
	"time"
)

type job struct {
	i, max int
	text   string
}

func outoutText(j *job) {
	for i := j.i; i < j.max; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
	}

}

func main() {

	hello := job{i: 0, max: 100, text: "Hello"}
	world := job{i: 0, max: 300, text: "World"}

	go outoutText(&hello)
	go outoutText(&world)
	for { // IF i remove this for , then does not give out put as the main program ends.

	}

}

//Go routine is given as 100 but it displayed as much as it can execute until the non go routine is executed.
//Notes: When a go routine is invoked , it waits for the blocking code to complete before it executes concurrency
//As soon as the main function ends , all other go routines also closes..hence the time lapse of the execution is
//always till the end of the main..
