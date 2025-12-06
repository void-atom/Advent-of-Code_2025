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

// Logic similar to evaluating mathematical expression
func DayThreePartTwo(input string) int {
	TOTAL_NUM_BATTERIES := 12
	input_data := strings.Split(input, "\n")
	answer := 0

	for _, num_string := range input_data {
		highest_joltage := internal.Stack[int]{}
		numbers := strings.TrimSpace(num_string)

		for idx, ch := range numbers {
			remaining := len(numbers) - idx - 1
			curr_num, _ := strconv.Atoi(string(ch))
			for {
				top, ok := highest_joltage.Peek()
				if ok == nil && top < curr_num && remaining+highest_joltage.Size() >= TOTAL_NUM_BATTERIES {
					highest_joltage.Pop()
				} else {
					break
				}
			}
			if highest_joltage.Size() < TOTAL_NUM_BATTERIES {
				highest_joltage.Push(curr_num)
			}
		}
		highest_value := ""
		for !highest_joltage.IsEmpty() {
			val, _ := highest_joltage.Pop()
			highest_value = strconv.Itoa(val) + highest_value
		}
		highest_value_int, _ := strconv.Atoi(highest_value)
		answer += highest_value_int

	}
	return answer

}
