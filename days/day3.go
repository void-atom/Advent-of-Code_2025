package days

import (
	"aoc2025/internal"
	"strconv"
	"strings"
)

func DayThreePartOne(input string) int {
	input_data := strings.Split(input, "\n")
	answer := 0

	for _, num_string := range input_data {
		numbers := strings.TrimSpace(num_string)
		result_array := internal.MaxTillRight(numbers)
		best := -1
		for i, val := range result_array {
			l_digit, _ := strconv.Atoi(string(numbers[i]))
			if val < 0 {
				continue
			}
			candidate := l_digit*10 + val
			best = max(candidate, best)
		}
		if best > 0 {
			answer += best
		}

	}
	return answer
}

func DayThreePartTwo(input string) int {

	return 1
}
