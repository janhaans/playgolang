package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
	Roles  map[string]int
}

type Title struct {
	Title string
}

func main() {
	movies := make([]Movie, 0)

	movies = append(movies, Movie{
		Title:  "movie1",
		Year:   1940,
		Color:  false,
		Actors: []string{"actor1", "actor2"},
		Roles:  map[string]int{"role1": 1, "role2": 2},
	})

	movies = append(movies, Movie{
		Title:  "movie2",
		Year:   1990,
		Color:  true,
		Actors: []string{"actor3", "actor4"},
		Roles:  map[string]int{"role1": 4, "role2": 5},
	})

	movies = append(movies, Movie{
		Title:  "movie3",
		Year:   2010,
		Color:  true,
		Actors: []string{"actor4", "actor5"},
		Roles:  map[string]int{"role1": 2, "role2": 2},
	})

	data, _ := json.MarshalIndent(movies, "", "	")

	fmt.Println(movies)
	fmt.Println(string(data))

	title := []Title{}

	json.Unmarshal(data, &title)
	fmt.Println(title)

}
