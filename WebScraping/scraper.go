package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type links []string
type tracks []string

func searchGoogle(query string) links {
	searchQuery := "https://www.google.co.in/search?q=" + formatQuery(query)
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
		track := item.Find("b").Text()
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
