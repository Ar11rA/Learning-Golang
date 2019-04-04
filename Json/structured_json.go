package main

import (
	"encoding/json"
	"log"
)

type Bird struct {
	Species     string
	Description string
}

type BirdWithDimension struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

type Dimensions struct {
	Height int
	Width  int
}

type Bird2 struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

func main() {
	str1 := `{
		"species": "pigeon",
		"description": "likes to perch on rocks"
	}`

	str2 := `[
  {
    "species": "pigeon",
    "description": "likes to perch on rocks"
  },
  {
    "species":"eagle",
    "description":"bird of prey"
  }
	]`

	str3 := `{
  "species": "pigeon",
  "description": "likes to perch on rocks",
  "dimensions": {
    "height": 24,
    "width": 10
  }
	}`
	var bird Bird
	var birds []Bird
	var birdWithDimension BirdWithDimension
	json.Unmarshal([]byte(str1), &bird)
	json.Unmarshal([]byte(str2), &birds)
	json.Unmarshal([]byte(str3), &birdWithDimension)
	log.Printf("str1: %v", bird)
	log.Printf("str2: %v", birds)
	log.Printf("str3: %v", birdWithDimension)
}
