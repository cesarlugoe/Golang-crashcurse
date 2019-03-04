package main

import (
	"fmt"
	"net/http"
)

// example of checking status of different webpages
// the problem with this aproach is that is syncronous, the program
// stops to wait for the function get result before starting again
// so for a big list, it would take hours.

//update now with correct aproach

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	//create a channel for communication between go routines &
	// pass it to the function using a new go routine
	c := make(chan string)

	for _, link := range links {
		// Adding the "go" routine keyword opens a new routine
		// that executes the http.Get function inside checkLink, and waits for its
		// response, while waiting, it signals back the main routine to continue
		// letting the for loop to call checkLink again with next argument and so on..
		go checkLink(link, c)
	}

	//infinite loop will only be execute internal code when receiving new
	// response from checklink.
	for {
		// '<-c' is a blocking code line, it'll make the program stop, until
		// it recieves the message from the channel to print, so when the
		// first go routine completes and sends the message from the channel
		// this for loop calls checklink again with recieving channel parameter
		//and waits for the next message.

		go checkLink(<-c, c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		//passing the link back to the channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

// * Finally we are changing the program so that it makes continous
// calls to the links until they get an error
