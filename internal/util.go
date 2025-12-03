package internal

import (
	"os"
)

// function to just check errs from std lib functions
func Check_error(e error) {
	if e != nil {
		panic(e)
	}
}

// Reads the entire input and returns a string of the input
func Read_file(file_path string) string {
	input_data, err := os.ReadFile(file_path)
	Check_error(err)
	// fmt.Print(string(input_data))
	return string(input_data)
}

// extracts the number from the input.
// Example: input = "L123", initial_index = 1 (The initial_index is always insured to be an integer)
// Returns -> 3 [the total digits before the start of newline]
// Checks: out of bound errors and total look forward chars

func Find_number_index(input string, initial_index int) int {
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

func Sum_Int_Slice(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
