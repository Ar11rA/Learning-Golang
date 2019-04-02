package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

type data struct {
	Name  string
	Value int
}

func main() {
	str := `
	{
	"name": "{{ .Name }}",
	"value": "{{ .Value }}"
	}
	`
	t := template.New("anything")
	t, err := t.Parse(str)
	if err != nil {
		log.Fatalf("template.Parse() failed with '%s'\n", err)
	}

	data := data{
		Name:  "abc",
		Value: 5,
	}

	buf := &bytes.Buffer{}

	err = t.Execute(buf, data)

	fmt.Println("**********", buf.String())
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
}
