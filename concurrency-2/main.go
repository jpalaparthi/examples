// concurrency-1 project main.go
package main

import (
	"fmt"
	"sync"
	"time"
)

type job struct {
	i, max int
	text   string
}

func outoutText(j *job, goGroup *sync.WaitGroup) {
	for i := j.i; i < j.max; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
	}
	goGroup.Done()

}

func main() {
	goGroup := &sync.WaitGroup{}
	hello := job{i: 0, max: 500, text: "Hello"}
	world := job{i: 0, max: 300, text: "World"}

	go outoutText(&hello, goGroup)
	go outoutText(&world, goGroup)
	goGroup.Add(2)
	goGroup.Wait()

}

//wait group delays the execution of the main until all go routines are completed
//waitt group object to be made done once it is done.
