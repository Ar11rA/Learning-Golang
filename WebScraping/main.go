package main

func main() {
	links := searchGoogle("bob dylan")
	tracks := getTopTen(links)
	print(tracks)
	tracks.download()
}
