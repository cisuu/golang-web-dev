package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type user struct {
	Name  string
	Motto string
	Admin bool
}

func main() {
	u1 := user{
		Name:  "Cisuu",
		Motto: "Failure is not an option",
		Admin: true,
	}

	u2 := user{
		Name:  "Asia",
		Motto: "Loco",
		Admin: true,
	}

	u3 := user{
		Name:  "JJ",
		Motto: "Da King",
		Admin: false,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
