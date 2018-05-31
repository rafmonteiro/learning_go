package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// send msg into the channel
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up!")
	c <- "It's UP"
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
	// this for loop works like a while!
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
	//fmt.Println(<-c) // <--- this is a blocking code.
}
