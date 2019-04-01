package main

import (
	"fmt"
	"net/http"
)

// func main() {
// 	links := []string{
// 		"http://www.google.com",
// 		"http://www.amazon.in",
// 		"http://www.flipkart.com",
// 		"http://www.facebook.com",
// 		"http://www.stackoverflow.com",
// 	}

// 	done := make(chan bool)

// 	for _, link := range links {
// 		go checkLink(link, done)
// 	}
// 	for i := 0; i < len(links); i++ {
// 		<-done
// 	}
// }

func checkLink(link string, done chan bool) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		done <- false
	}
	fmt.Println("Website is up", link)
	done <- true
}
