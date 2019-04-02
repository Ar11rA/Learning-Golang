package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	path := "resource.txt"
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("os.Open() failed with '%s'\n", err)
	}
	defer f.Close()

	d, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("ioutil.ReadAll() failed with '%s'\n", err)
	}

	lines := bytes.Split(d, []byte{'\n'})
	s := string(d)
	log.Printf("File %s has %d lines\n", path, len(lines))
	log.Println(s)
}
