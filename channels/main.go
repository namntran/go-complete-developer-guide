package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) // c == channel, make channel of type string

	for _, link := range links {
		go checkLink(link, c) // create a new thread go routine, and run "checkLink()" function with it
		// go routines executes lines of codes (line by line) inside of one function, always use "go" keyword right before function call
	}

	for l := range c {
		go func(link string) { // key word is "go" - creates a second go routines or go engine
			time.Sleep(10 * time.Second)
			checkLink(link, c)
		}(l) //this parentheses executes the function literal
		// for i := 0; i < len(links); i++ {
		// fmt.Println(<-c) // receiving a message through a channel - this is a blocking call
		// fmt.Println(<-c)
		// fmt.Println(<-c)
		// fmt.Println(<-c)
		// fmt.Println(<-c)
	}
}
func checkLink(link string, c chan string) { // go routine runs inside "checkLink" function
	_, err := http.Get(link) // called a blocking call, goroutine runs next iteration of for loop and creates the next goroutine
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "I think it might be down" //goroutine waits to send a message into the channel
		return
	}
	fmt.Println(link, "is up!")
	c <- "It's definitely working hombre"
}
