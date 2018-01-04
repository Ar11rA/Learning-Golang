// file: list_posts.go
package main

import (
	"fmt"
)

func main() {
	links := searchGoogle("bob dylan")
	fmt.Println("------------------- LINKS ----------------------")
	print(links)
	tracks := getTopTen(links)
	fmt.Println("------------------- TRACKS ----------------------")
	print(tracks)
	fmt.Println("------------------- DOWNLOADING ----------------------")
	tracks.download()
}
