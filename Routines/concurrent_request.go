package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Value is a model
type Value struct {
	ID    int    `json:"id"`
	Quote string `json:"quote"`
}

// Quote is a model
type Quote struct {
	Type  string `json:"type"`
	Value Value  `json:"value"`
}

func makeCall(ctr int, url string, ch chan bool) {
	resp, _ := http.Get(url)
	contents, _ := ioutil.ReadAll(resp.Body)
	var q Quote
	json.Unmarshal([]byte(contents), &q)
	log.Printf("Count and Id is %d-%d\n", ctr+1, q.Value.ID)
	ch <- true
}

func main() {
	log.SetFlags(4)
	url := "http://gturnquist-quoters.cfapps.io/api/random"
	ch := make(chan bool)
	for i := 0; i < 5; i++ {
		go makeCall(i, url, ch)
	}
	for i := 0; i < 5; i++ {
		<-ch
	}

}
