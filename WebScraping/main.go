package main

import "fmt"

func main() {
	fmt.Print("Enter text: ")
	var artistName string
	fmt.Scanln(&artistName)
	links := searchGoogle(artistName)
	tracks := getTopTen(links)
	print(tracks)
	tracks.download()
}
