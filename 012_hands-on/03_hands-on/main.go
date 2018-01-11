package main

import (
	"html/template"
	"log"
	"os"
)

var t *template.Template

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

func init() {
	t = template.Must(template.ParseFiles("t.gohtml"))
}

func main() {
	hotels := []Hotel{
		Hotel{"Northside", "Street North", "LA", "1234", "Northern"},
		Hotel{"Centralside", "Street Central", "San Francisco", "1234", "Central"},
		Hotel{"Southside", "Street South", "San Jose", "1234", "Southern"},
	}

	err := t.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
