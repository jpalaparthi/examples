// concurrency-noofThreads-1 project main.go
package main

import (
	"fmt"
	"runtime"
)

func listofThreads() int {
	threads := runtime.GOMAXPROCS(0) // No parameter or 0 means not to change the mode of it.. even if 3 or4 given
	// it gives available threads only
	return threads

}
func main() {
	runtime.GOMAXPROCS(4)
	fmt.Println("Threads available in GO:", listofThreads())
}
