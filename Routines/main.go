package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.amazon.in",
		"http://www.flipkart.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
	}

	c := make(chan bool)

	for _, link := range links {
		go checkLink(link, c)
	}
	for i := range links {
		fmt.Println(i+1, <-c)
	}
}

func checkLink(link string, c chan bool) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		c <- false
	}
	fmt.Println("Website is up")
	c <- true
}
