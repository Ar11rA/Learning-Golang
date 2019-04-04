package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gojektech/heimdall/httpclient"
)

type data struct {
	Name  string
	Value int
}

func getResourceContents(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("os.Open() failed with '%s'\n", err)
	}
	defer f.Close()

	d, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("ioutil.ReadAll() failed with '%s'\n", err)
	}
	return d
}

func formRequest(contents []byte, d data) *bytes.Buffer {
	t := template.New("request_template")
	s := string(contents)
	t, _ = t.Parse(s)
	buf := &bytes.Buffer{}
	err := t.Execute(buf, d)
	if err != nil {
		log.Fatalf("templating failed with '%s'\n", err)
	}
	log.Println("Final request:", buf.String())
	return buf
}

func makePostRequest(remoteURL string, buf *bytes.Buffer, out chan string) {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	buff := []io.Reader{buf}
	combined := io.MultiReader(buff...)
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	response, _ := client.Post(remoteURL, combined, headers)
	byteResponse, _ := ioutil.ReadAll(response.Body)
	out <- string(byteResponse)
}

func main() {
	path := "./resources/request_template.txt"
	data1 := data{
		Name:  "def",
		Value: 10,
	}
	data2 := data{
		Name:  "abc",
		Value: 5,
	}
	data3 := data{
		Name:  "xyz",
		Value: 15,
	}
	fileContents := getResourceContents(path)
	requestTemplate1 := formRequest(fileContents, data1)
	requestTemplate2 := formRequest(fileContents, data2)
	requestTemplate3 := formRequest(fileContents, data3)
	out := make(chan string)
	go makePostRequest("http://localhost:3000/ping", requestTemplate1, out)
	go makePostRequest("http://localhost:3000/ping", requestTemplate2, out)
	go makePostRequest("http://localhost:3000/ping", requestTemplate3, out)
	log.Println(<-out)
	log.Println(<-out)
	log.Println(<-out)
}
