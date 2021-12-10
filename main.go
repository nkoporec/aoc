package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"github.com/nkoporec/aoc/days"
	"github.com/nkoporec/aoc/helpers"
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
		days.Day1Second(data)
		break
	case 2:
		days.Day2(data)
		days.Day2Second(data)
	}
}

func getData(filepath string) []string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	records := strings.Split(string(content), "\n")
	records = helpers.RemoveEmptyItemsInSlice(records)

	return records
}

