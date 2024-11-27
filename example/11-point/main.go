// This program demonstrates the difference between passing a value and passing a pointer to a function in Go.
// The `add2` function takes an integer value and adds 2 to it, but since it operates on a copy of the value,
// the original value remains unchanged.
// The `add2ptr` function takes a pointer to an integer and adds 2 to the value at that memory address,
// modifying the original value.
package main

import "fmt"

func add2(n int) {
	n += 2
}

func add2ptr(n *int) {
	*n += 2
}

func main() {
	n := 5
	add2(n)
	fmt.Println(n) // 5
	add2ptr(&n)
	fmt.Println(n) // 7
}
