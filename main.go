package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename) // note this is a pointer to the csvFilename string
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s\n", *csvFilename)
		os.Exit(1)
	}
	_ = file
}
