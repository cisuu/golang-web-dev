package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, City, Zip string
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h1 := hotel{
		Name: "Hutnik",
		City: "Stw",
		Zip:  "37-459",
	}

	h2 := hotel{
		Name: "Airport Inn",
		City: "Richmond",
		Zip:  "???",
	}
	h3 := hotel{
		Name: "Hutnik3",
		City: "Stw3",
		Zip:  "37-4593",
	}

	h4 := hotel{
		Name: "Airport Inn4",
		City: "Richmond4",
		Zip:  "???4",
	}
	h5 := hotel{
		Name: "Hutnik5",
		City: "Stw5",
		Zip:  "37-4595",
	}

	h6 := hotel{
		Name: "Airport Inn6",
		City: "Richmond6",
		Zip:  "???6",
	}

	r1 := region{
		Region: "Southern",
		Hotels: []hotel{h1, h2},
	}

	r2 := region{
		Region: "Central",
		Hotels: []hotel{h3, h4},
	}

	r3 := region{
		Region: "Northern",
		Hotels: []hotel{h5, h6},
	}

	regions := []region{r1, r2, r3}

	err := tpl.Execute(os.Stdout, regions)
	if err != nil {
		log.Fatalln(err)
	}
}
