package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// associate URLs requested to functions that handle requests
	http.HandleFunc("/hello", helloRequest)
	http.HandleFunc("/", getRequest)

	// start web server
	log.Println("Listening on http://localhost:9998/")
	log.Fatal(http.ListenAndServe(":9998", nil))
}

// basic handler for /hello request
func helloRequest(w http.ResponseWriter, r *http.Request) {

	// Fprint writes to a writer, in this case the http response
	fmt.Fprint(w, "Hello API!")
	return
}

// this function simply serves up the file requested, or
// an index list if a directory is requested
func getRequest(w http.ResponseWriter, r *http.Request) {
	fileRequested := "./" + r.URL.Path
	http.ServeFile(w, r, fileRequested)
	return
}
