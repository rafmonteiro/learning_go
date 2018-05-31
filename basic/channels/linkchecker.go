package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// send msg into the channel
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link

}

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://www.twohat.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// creates go routines
		go checkLink(link, c)
	}
	// alternative for loop
	for l := range c {
		// function literal!!!
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}
