package main

import "encoding/csv"
import "fmt"
import "flag"
import "os"

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file of questions and answers")
	flag.Parse()
	_ = csvFilename

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file : %s", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse csv file")
	}
	fmt.Println(lines)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
