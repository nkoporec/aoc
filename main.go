package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"github.com/nkoporec/aoc/days"
)

func main() {
	filepath := flag.String("data", "", "AOC data path")
	day := flag.Int("day", 0, "AOC day")
	flag.Parse()

	if *filepath == "" {
		log.Fatal("Please provide a path to the data file")
	}

	if *day ==  0 {
		log.Fatal("Please provide a valid day")
	}

	data := getData(*filepath)
	executeDay(*day, data)
}

func executeDay(day int, data []string) {
	switch day {
	case 1:
		days.Day1(data)
	}
}

func getData(filepath string) []string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	records := strings.Split(string(content), "\n")
	records = removeEmptyLines(records)

	return records
}


// Remove empty lines from the data
func removeEmptyLines(data []string) []string {
	var result []string
	for _, line := range data {
		if line != "" {
			result = append(result, line)
		}
	}	
	return result
}
