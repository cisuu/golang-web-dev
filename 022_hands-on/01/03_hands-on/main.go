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
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
