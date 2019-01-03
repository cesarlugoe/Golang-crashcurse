package main

import "fmt"

func main() {
	//arrays
	// declare & assing
	fruitArray := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArray[0])
	fmt.Println(fruitArray[1])
	fmt.Println(fruitArray)

	//Slice
	fruitSlice := []string{"Apple", "Orange", "Grapes"}
	fmt.Println("length", len(fruitSlice))

	fruitSlice = append(fruitSlice, "Cherry", "watermelon")
	fmt.Println("appended:", fruitSlice)

	copiedFruit := fruitSlice
	fmt.Println("copied:", copiedFruit)

	slicedFruit := fruitSlice[2:4]
	fmt.Println("Sliced", slicedFruit)

	slicedFruit = fruitSlice[:4]
	fmt.Println("Sliced2", slicedFruit)

	// pointer
	var a = 5
	p := &a // p holds variable a memory address
	fmt.Printf("Address of var a: %p\n", p)
	fmt.Printf("Value of var a: %v\n", *p)

	// Let's change a value (using the initial variable or the pointer)
	*p = 3 // using pointer
	a = 3  // using initial var

	fmt.Printf("Address of var a: %p\n", p)
	fmt.Printf("Value of var a: %v\n", *p)

}
