package main

import "fmt"

func main() {
	// the first [string] refers to the keys and the 2nd to the values
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b745",
	}

	// adding a key
	colors["white"] = "#fffff"

	// deleting a key
	delete(colors, "green")
	printMap(colors)
	fmt.Println(colors)
}

// iterating with maps
func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("hex code for ", color, " is ", hex)
	}
}
