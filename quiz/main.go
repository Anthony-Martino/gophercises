package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "The name of the csv containing the problem set.")
	flag.Parse()

	csvFile, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatalln("Error opening csv file", err)
	}

	r := csv.NewReader(csvFile)

	rows, err := r.ReadAll()

	if err != nil {
		log.Fatalln("Error parsing csv file", err)
	}
	for i, row := range rows {
		fmt.Println("Problem", i+1, row[0], "=")
	}

}
