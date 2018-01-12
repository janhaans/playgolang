package main

import (
	"html/template"
	"log"
	"os"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("t.gohtml"))
}

type MenuItem struct {
	Title       string
	Description string
	Price       float64
}

type Menu struct {
	Breakfast []MenuItem
	Lunch     []MenuItem
	Dinner    []MenuItem
}

func main() {
	restaurant := Menu{
		Breakfast: []MenuItem{
			MenuItem{"Omelet", "Gebakken omelet met hele dooier", 6.5},
			MenuItem{"Zacht ei", "Zacht gekookt eitje", 2.5},
		},
		Lunch: []MenuItem{
			MenuItem{"Tonijn", "Broodje verse tonijn", 4.25},
			MenuItem{"Zalm", "Broodje verse zalm", 5.25},
		},
		Dinner: []MenuItem{
			MenuItem{"Hamburger", "Broodje hamburger met friet", 12.75},
			MenuItem{"Biefstuk", "Medium gebakken biefstuk met friet", 19.25},
		},
	}

	err := t.Execute(os.Stdout, restaurant)
	if err != nil {
		log.Fatalln(err)
	}

}
