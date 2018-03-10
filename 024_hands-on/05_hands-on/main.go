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

func dog(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./starting-files/public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", dog)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
