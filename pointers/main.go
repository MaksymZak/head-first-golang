package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Declare an integer variable
	var num int = 42

	// Get the address of the variable
	var ptr *int = &num

	// Print the address of the variable
	fmt.Printf("Address of num: %p\n", ptr)

	// Print the value at that address
	fmt.Printf("Value at address: %d\n", *ptr)

	// Modify the value using the pointer
	*ptr = 100

	// Print the modified value
	fmt.Printf("Modified value of num: %d\n", num)

	// Print the size of the pointer
	fmt.Printf("Size of pointer: %d bytes\n", unsafe.Sizeof(ptr))
}
