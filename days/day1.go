package days

import (
	"fmt"
	"strconv"

	"github.com/nkoporec/aoc/helpers"
)


func Day1(data []string) {
	fmt.Println("Measurments increases", day1_count(data))
}

func Day1Second(data []string) {
	fmt.Println("Three-measurement increases", day1_second_count(data))
}

func day1_count(data []string) int {
	count := 0
	prev_item := 0

	for key, line := range data {
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		// Skip the first item
		if key == 0 {
			prev_item = number
			continue
		}

		if (number > prev_item) {
			count ++
			prev_item = number
		} else {
			prev_item = number
		}
	}

	return count
}


func day1_second_count(data []string) int {
	var chunks [][]string
	var count []string
	var index int

	for i := 0; index < len(data); i += 1 {
		index = i + 3
		chunks = append(chunks, data[i:index])
	}

	// Count the chunks sum values.
	for _, chunk := range chunks {
		sum := 0
		chunk = helpers.RemoveEmptyItemsInSlice(chunk)
		for _, line := range chunk {
			number, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			sum += number
		}

		count = append(count, strconv.Itoa(sum))
	}

	result := day1_count(count)

	return result
}
