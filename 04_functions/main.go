package main

import "fmt"

// the second string is to declare the type of the return value
func greetings(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}
func main() {
	fmt.Println(greetings("Cesar"))
	fmt.Println(getSum(3, 5))
}
