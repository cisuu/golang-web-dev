package main

import (
	"bufio"
	"encoding/csv"
	"html/template"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type action struct {
	Date time.Time
	Open float64
}

var (
	actions []action
	tpl     *template.Template
)

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	csvData, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer csvData.Close()
	parseCSV(csvData)

	err = tpl.Execute(os.Stdout, actions)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseCSV(csvData *os.File) {
	reader := csv.NewReader(bufio.NewReader(csvData))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		if record[0] == "Date" {
			continue
		}
		date, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			log.Println("Skipping because of invalid data: ", err)
			continue
		}
		open, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Println("Skipping because of invalid float value: ", err)
			continue
		}
		act := action{
			Date: date,
			Open: open,
		}
		actions = append(actions, act)
	}
}
