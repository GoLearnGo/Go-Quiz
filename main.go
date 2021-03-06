/******************************************************************************
Go Timed Quiz
	Create a simple, timed quiz to practice programming in Go

Author: Jason Flinn
OG Author: Gophercises
Date: 1/15/19
******************************************************************************/

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename) // note this is a pointer to the csvFilename string
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
		os.Exit(1)
	}

	// create a csv reader
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer) // this will not work if answers are multiple word strings
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

//might want to make a validator for the csv, but not needed for this exercise
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines)) // assume every line is a problem
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), // removes any unneccessary spaces to prevent false errors
		}
	}
	return ret
}

// a struct makes it easier to not have to change the code too much if the input file format/type changes
type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
