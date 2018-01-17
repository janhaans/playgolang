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

func index(w http.ResponseWriter, req *http.Request) {
	err := t.ExecuteTemplate(w, "index.html", "")
	if err != nil {
		log.Fatal(err)
	}
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := (t.ExecuteTemplate(w, "dog.html", ""))
	if err != nil {
		log.Fatal(err)
	}
}

func me(w http.ResponseWriter, req *http.Request) {
	err := t.ExecuteTemplate(w, "me.html", "")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileserver := http.FileServer(http.Dir("./img"))
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)
	http.Handle("/images/", http.StripPrefix("/images", fileserver))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
