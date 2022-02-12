// Modify the `echo` program to print the index and value of each of its argument,
// one per line.

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}