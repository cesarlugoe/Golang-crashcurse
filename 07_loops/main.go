package main

import (
	"fmt"
)

func main() {
	// long for loop method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//short method

	for i := 0; i <= 10; i++ {
		fmt.Printf("number %d \n", i)
	}

}
