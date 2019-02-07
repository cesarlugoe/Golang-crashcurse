package main

import "fmt"

// the second string is to declare the type of the return value
func greetings(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func say(comment string, name string) string {
	return name + " says " + comment
}
func main() {
	myName := "Cesar"
	sumValueA, sumValueB := 3, 5

	fmt.Println(greetings(myName))
	fmt.Println(getSum(sumValueA, sumValueB))
	fmt.Println(say(myName, "'Que mundos'"))
}
