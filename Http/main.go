package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Track json format
type Track struct {
	Title          string `json:"title"`
	Artist         string `json:"artist"`
	URL            string `json:"url"`
	Image          string `json:"image"`
	ThumbnailImage string `json:"thumbnail_image"`
}

// Tracks ....
type Tracks struct {
	Collection []Track
}

func main() {
	response, err := http.Get("http://rallycoding.herokuapp.com/api/music_albums")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	// albums := make([]byte, 99999)
	// response.Body.Read(albums)
	// fmt.Println(string(albums))

	bs, err := ioutil.ReadAll(response.Body)
	tracks := make([]Track, 0)
	data := []byte(string(bs))
	_ = json.Unmarshal(data, &tracks)
	fmt.Printf("Results: %v\n", tracks)

}
