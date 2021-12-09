package days

import (
	"fmt"
	"strconv"
)

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

func Day1(data []string) {
	count := day1_count(data)

	fmt.Println("Measurments increases", count)
}
