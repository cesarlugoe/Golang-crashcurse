package main

import (
	"fmt"
)

func main() {
	//define map
	newMap := make(map[string]string)

	//assign key-values
	newMap["Cesar"] = "Lugo"
	newMap["Jeniree"] = "Parra"
	newMap["Maria"] = "Del Mar"

	fmt.Println(newMap)
	fmt.Println(newMap["Cesar"])
	fmt.Println(len(newMap))

	//delete
	delete(newMap, "Maria")
	fmt.Println(newMap)

	// declare & assign
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
