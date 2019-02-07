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

	for i := 1; i <= 100; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println("FizzBuzz")
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
		}
		fmt.Println(i)
	}

}
