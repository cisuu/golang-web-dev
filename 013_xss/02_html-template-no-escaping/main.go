package main

import (
	"html/template"
	"log"
	"os"
)

type page struct {
	Title   string
	Heading string
	Input   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	home := page{
		Title:   "Nothing Escaped",
		Heading: "Nothing is escaped with text/html",
		Input:   `<script>alert("You!");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}
