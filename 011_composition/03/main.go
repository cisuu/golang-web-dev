package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Fall GO", "4"},
				course{"CSCI-402", "Introduction to Fall GO2", "4"},
				course{"CSCI-403", "Introduction to Fall GO3", "4"},
			},
		},
		Spring: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Spring GO", "4"},
				course{"CSCI-402", "Introduction to Spring GO2", "4"},
				course{"CSCI-403", "Introduction to Spring GO3", "4"},
			},
		},
		Summer: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Summer GO", "4"},
				course{"CSCI-402", "Introduction to Summer GO2", "4"},
				course{"CSCI-403", "Introduction to Summer GO3", "4"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
