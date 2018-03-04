package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type dish struct {
	Name  string
	Price float64
}

type meal struct {
	Name   string
	Dishes []dish
}

type restaurant struct {
	Name string
	Menu []meal
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m1 := meal{
		Name: "Breakfast",
		Dishes: []dish{
			dish{
				Name:  "1",
				Price: 1,
			},
			dish{
				Name:  "2",
				Price: 2,
			},
		},
	}

	m2 := meal{
		Name: "Lunch",
		Dishes: []dish{
			dish{
				Name:  "3",
				Price: 3,
			},
			dish{
				Name:  "4",
				Price: 4,
			},
		},
	}

	m3 := meal{
		Name: "Dinner",
		Dishes: []dish{
			dish{
				Name:  "5",
				Price: 5,
			},
			dish{
				Name:  "6",
				Price: 6,
			},
		},
	}

	restaurants := []restaurant{
		restaurant{
			Name: "Good Food",
			Menu: []meal{m1, m2, m3},
		},
	}

	if err := tpl.Execute(os.Stdout, restaurants); err != nil {
		log.Fatalln(err)
	}
}
