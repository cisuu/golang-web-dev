package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name string
	}{
		Name: "Cisuu",
	}
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "dog")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Cisuu")
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}
