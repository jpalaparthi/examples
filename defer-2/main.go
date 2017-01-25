// defer-2 project main.go
package main

import (
	"fmt"
)

func main() {
	val := new(int)

	defer fmt.Println(*val)

	for i := 0; i < 100; i++ {
		*val = *val + i
	}

	k := 0
	defer fmt.Println(k)

	for i := 0; i < 100; i++ {
		k = k + i
	}
}

// even though the defer function is executed after the normal execution, the value of the val state does not change.
//
