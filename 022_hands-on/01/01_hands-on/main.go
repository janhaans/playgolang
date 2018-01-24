package main

import (
	"html/template"
	"log"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("template/*"))
}

type Person struct {
	Fname string
	Bname string
	Age   int32
}

var person Person

func index(w http.ResponseWriter, req *http.Request) {
	err := t.ExecuteTemplate(w, "index.html", person)
	if err != nil {
		log.Fatal(err)
	}
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := (t.ExecuteTemplate(w, "dog.html", person))
	if err != nil {
		log.Fatal(err)
	}
}

func me(w http.ResponseWriter, req *http.Request) {
	err := t.ExecuteTemplate(w, "me.html", person)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	person = Person{"Anna-Maria", "Gobati", 50}
	fileserver := http.FileServer(http.Dir("./img"))
	http.Handle("/", http.HandlerFunc(index))
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)
	http.Handle("/images/", http.StripPrefix("/images", fileserver))
	http.Handle("/try", http.NotFoundHandler())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
