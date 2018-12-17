package main

import "fmt"

// Function to be called before anything else...
func init() {
	fmt.Printf("Initializing...\n")
	fmt.Printf("....working....\n")
	fmt.Printf("...init finished\n")
}
