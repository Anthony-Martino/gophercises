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

	problems := make([]problem, len(rows))
	for i, row := range rows {
		problems[i] = problem{
			q: row[0],
			a: row[1],
		}
	}

	correct := 0
	for i, problem := range problems {
		fmt.Print("Problem ", i+1, ": ", problem.q, " = ")

		var answer string
		fmt.Scan(&answer)
		if answer == problem.a {
			correct++
		}
	}

	fmt.Println("You answered", correct, "out of", len(rows), "questions correctly.")

}
