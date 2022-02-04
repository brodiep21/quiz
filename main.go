package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file.")
	}
	problems := parseLines(lines)
	correct := 0
	incorrect := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Println("Incorrect :(")
			incorrect++
		}
	}
	fmt.Printf("You scored %d out of %d! Missing %d questions\n", correct, len(problems), incorrect)
}

func parseLines(lines [][]string) []problem {
	returner := make([]problem, len(lines))
	for i, line := range lines {
		returner[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return returner
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
