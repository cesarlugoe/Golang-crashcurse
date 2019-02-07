package main

import (
	"fmt"
)

func main() {

	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a) //%T refers to the type of value of the variable
	fmt.Printf("%T\n", b)

	// to get the value that 'b' points to use *

	fmt.Println(*b)

	//change value of 'a' with the pointer of 'b'

	*b = 10

	fmt.Println(a)
}
