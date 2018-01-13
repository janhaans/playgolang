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

type item struct {
	Title       string
	Description string
	Price       float64
}

type meal struct {
	Type  string
	Items []item
}

type menu struct {
	Meals []meal
}

type restaurant struct {
	Name string
	Menu menu
}

type restaurants []restaurant

func main() {
	Restaurants := restaurants{
		restaurant{
			Name: "Het zachte ei",
			Menu: menu{
				Meals: []meal{
					meal{
						Type: "ontbijt",
						Items: []item{
							item{"Omelet", "Gebakken omelet met hele dooier", 6.5},
							item{"Zacht ei", "Zacht gekookt ei", 2.5},
						},
					},
					meal{
						Type: "lunch",
						Items: []item{
							item{"Tonijn", "Broodje verse tonijn", 4.25},
							item{"Zalm", "Broodje verse zalm", 5.25},
						},
					},
					meal{
						Type: "dinner",
						Items: []item{
							item{"Hamburger", "Broodje hamburger met friet", 12.75},
							item{"Biefstuk", "Medium gebakken biefstuk met friet", 19.25},
						},
					},
				},
			},
		},
		restaurant{
			Name: "Het harde ei",
			Menu: menu{
				Meals: []meal{
					meal{
						Type: "ontbijt",
						Items: []item{
							item{"Croissant", "3 hele croissants", 4.5},
							item{"Kaas", "Broodje kaas", 5.5},
						},
					},
					meal{
						Type: "lunch",
						Items: []item{
							item{"Rosbief", "Broodje verse rosbief", 4.25},
							item{"Filet Americain", "Broodje verse filer americain", 5.25},
						},
					},
					meal{
						Type: "dinner",
						Items: []item{
							item{"Pot au fau", "Een franse klassieker", 22.75},
							item{"Groentesoep", "Verse groentesoep", 6.25},
						},
					},
				},
			},
		},
	}

	err := t.Execute(os.Stdout, Restaurants)
	if err != nil {
		log.Fatalln(err)
	}

}
