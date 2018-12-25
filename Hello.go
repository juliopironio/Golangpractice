// Package main implements the application's entry point
// Will print a list of command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	Comp1()
	for i, arg := range os.Args[0:] {
		arg = os.Args[i]
		fmt.Printf( "%d %s\n", i, arg)
	}
}
