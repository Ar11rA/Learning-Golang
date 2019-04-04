package main

import (
	"encoding/json"
	"log"
)

func main() {
	str := `{
		"birds": {
			"pigeon":"likes to perch on rocks",
			"eagle":"bird of prey"
		},
		"animals": 1
	}
	`
	var result map[string]interface{}
	json.Unmarshal([]byte(str), &result)
	log.Println(result)
	log.Println(result["birds"])
	log.Println(result["birds"].(map[string]interface{})["pigeon"])
	log.Println(result["birds"].(map[string]interface{})["eagle"])
	log.Println(result["animals"])
}
