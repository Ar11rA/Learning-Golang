package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type links []string
type tracks []string

func searchGoogle(query string) links {
	searchQuery := "https://www.google.co.in/search?q=" + formatQuery(query) + "+thetoptens"
	doc, err := goquery.NewDocument(searchQuery)
	if err != nil {
		log.Fatal(err)
	}
	results := links{}
	doc.Find("h3").Each(func(index int, item *goquery.Selection) {
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		url := strings.Split(link, "=")[1]
		url = escapeString(url)
		results = append(results, url)
	})
	return results
}

func formatQuery(query string) string {
	return strings.Replace(query, " ", "+", len(query)-1)
}

func getTopTen(l links) tracks {
	doc, err := goquery.NewDocument(l.first())
	if err != nil {
		log.Fatal(err)
	}
	results := tracks{}
	doc.Find(".image100").Each(func(index int, item *goquery.Selection) {
		track := item.Find("b").First().Text()
		results = append(results, track)
	})
	return results
}

func print(l []string) {
	for index, link := range l {
		fmt.Println(index+1, link)
	}
}

func (l links) first() string {
	return l[0]
}

func downloadSong(songName string) {
	_, err := exec.Command("/usr/local/bin/youtube-dl", "--extract-audio", "--audio-format", "mp3", getYoutubeURL(songName)).Output()
	if err != nil {
		fmt.Printf("Error donloading song %s - %s", songName, err)
		os.Exit(1)
	}
	fmt.Println("Download complete for: ", songName)
}

func getYoutubeURL(songName string) string {
	searchQuery := "https://www.google.co.in/search?q=" + formatQuery(songName) + "+youtube"
	doc, err := goquery.NewDocument(searchQuery)
	if err != nil {
		log.Fatal(err)
	}
	results := links{}
	doc.Find("h3").Each(func(index int, item *goquery.Selection) {
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		url := strings.Split(link, "=")[1]
		results = append(results, url)
	})
	return escapeString(results.first())
}

func (t tracks) download() {
	for _, track := range t {
		downloadSong(track)
	}
}

func escapeString(url string) string {
	url = strings.Replace(url, "%3F", "?", len(url)-1)
	url = strings.Replace(url, "%3D", "=", len(url)-1)
	url = strings.Replace(url, "&sa", "", len(url)-1)
	return url
}
