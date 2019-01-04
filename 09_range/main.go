package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5, 85, 47}

	// loop through ids
	// range returns 2 values, in the case of array/slice its index (i) and value (id)
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//if not using one of the returned values (this case index) "_" must substitute its assignment or error
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// acumulator
	sum := 0

	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum of IDs: %d\n", sum)

	// range with map
	userEmails := map[string]string{
		"Bob":    "bob@gmail.com",
		"Sharon": "sharon@gmail.com",
		"Cesar":  "cesar@gmail.com",
	}

	for key, value := range userEmails {
		fmt.Printf("%s's email: %s\n", key, value)
	}
}
