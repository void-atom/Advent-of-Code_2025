package internal

import (
	"os"
	"strconv"
)

// function to just check errs from std lib functions
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

// Reads the entire input and returns a string of the input
func ReadFile(file_path string) string {
	input_data, err := os.ReadFile(file_path)
	CheckError(err)
	// fmt.Print(string(input_data))
	return string(input_data)
}

// extracts the number from the input.
// Example: input = "L123", initial_index = 1 (The initial_index is always insured to be an integer)
// Returns -> 3 [the total digits before the start of newline]
// Checks: out of bound errors and total look forward chars

func FindNumberIndex(input string, initial_index int) int {
	// Assuming that the max number that can occur is upto 4 digits
	// Example: 0-9999
	if initial_index >= len(input) {
		return 0
	}
	const FORWARD_LOOK_INDEX = 5
	total_digits := 0

	for i := initial_index; i < i+initial_index+FORWARD_LOOK_INDEX && i < len(input); i++ {
		if input[i] == '\n' || input[i] == '\r' {
			break
		}

		total_digits++
	}
	//Returns total digits
	// Example 12 => 2
	return total_digits
}

// Gives sum of a slice
func SumIntSlice(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Returns an array where each index has the max value for all numbers behind this index
func MaxTillRight(numbers string) []int {
	max_array := make([]int, len(numbers))
	curr_max := -1

	for i := len(numbers) - 1; i > -1; i-- {
		max_array[i] = curr_max
		candidate, _ := strconv.Atoi(string(numbers[i]))
		curr_max = max(candidate, curr_max)
	}
	return max_array
}
