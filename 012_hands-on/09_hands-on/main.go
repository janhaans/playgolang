package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"os"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("t.gohtml"))
}

func main() {
	f, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	err = t.Execute(os.Stdout, records)
	if err != nil {
		log.Fatalln(err)
	}

}
