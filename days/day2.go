package days

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2(data []string) {
	fmt.Println("Three-measurement increases", day2_calculate(data))
}

func day2_calculate(data []string) int {
	var result int
	var horizontal int
	var depth int

	for _, line := range data {
		i := strings.Split(line, " ")

		if len(i) > 2 || len(i) < 2 {
			panic(i)
		}

		direction := i[0]

		amount, err := strconv.Atoi(i[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			horizontal += amount
			break

		case "down":
			depth += amount
			break

		case "up":
			depth -= amount
			break
		}
	}

	result = horizontal * depth

	return result
}
