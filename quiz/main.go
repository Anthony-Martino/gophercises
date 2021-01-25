package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "The filename of the csv.")
	timeLimit := flag.Int("timeLimit", 5, "Time limit of the quiz")
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

	problems := parseProblems(rows)
	correct := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

tag:
	for i, problem := range problems {
		fmt.Print("Problem ", i+1, ": ", problem.q, " = ")

		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scan(&answer)
			answerChannel <- answer
		}()
		select {
		case <-timer.C:
			break tag
		case answer := <-answerChannel:
			if answer == problem.a {
				correct++
			}
		}
	}
	fmt.Println("You answered", correct, "out of", len(rows), "questions correctly.")
}

func parseProblems(rows [][]string) []problem {
	problems := make([]problem, len(rows))
	for i, row := range rows {
		problems[i] = problem{
			q: row[0],
			a: row[1],
		}
	}
	return problems
}
