package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./starting-files/templates/*.gohtml"))
}

func doggo(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.Handle(
		"/resources/",
		http.StripPrefix(
			"/resources",
			http.FileServer(
				http.Dir("./starting-files/public"),
			),
		),
	)
	http.HandleFunc("/", doggo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
