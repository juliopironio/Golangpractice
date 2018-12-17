package main

import "fmt"

func main() {
	init()
	fmt.Printf("hello, world\n")
	fmt.Printf("modified today\n")
}

// Function to be called before anything else...
func init() {

	fmt.Printf("Initializing...\n")

	fmt.Printf("....working....\n")

	fmt.Printf("...init finished succersfully.\n")

}
