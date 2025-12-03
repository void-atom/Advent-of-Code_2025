package days

import (
	"strconv"
)

// // function to just check errs from std lib functions
// func check_error(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

// // Reads the entire input and returns a string of the input
// func readFile(file_path string) string {
// 	input_data, err := os.ReadFile(file_path)
// 	check_error(err)
// 	// fmt.Print(string(input_data))
// 	return string(input_data)
// }

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

// Logic:
// 1. Initialize counter and dial to 0 and 50
// 2. Loop through the input and extract the immediate number after L and R
// 3. If L is found, subtract number from dial and bring the result between 0-99 by using modulo 100 and addition by 100 if needed
// 4. If R is found, add number to dial and bring the result between 0-99 by using modulo 100.
// 5. If the result is 0, increment counter

func Day_one_part_one(input string) int {
	dial_value := 50
	zero_count := 0
	for i := 0; i < len(input); {
		if i >= len(input) {
			break
		}
		if input[i] == '\n' || input[i] == '\r' {
			i++
		} else if input[i] == 'L' {
			total_digits := Find_number_index(input, i)
			number, _ := strconv.Atoi(input[i+1 : i+total_digits])
			dial_value = (dial_value - number) % 100
			if dial_value < 0 {
				dial_value += 100
			}
			if dial_value == 0 {
				zero_count += 1
			}
			i += total_digits

		} else if input[i] == 'R' {
			total_digits := Find_number_index(input, i)
			number, _ := strconv.Atoi(input[i+1 : i+total_digits])
			dial_value = (dial_value + number) % 100
			if dial_value == 0 {
				zero_count += 1
			}
			i += total_digits

		}

	}
	return zero_count
}

func Day_one_part_two(input string) int {
	dial_value := 50
	zero_count := 0
	for i := 0; i < len(input); {

		if i >= len(input) {
			break
		}
		if input[i] == '\n' || input[i] == '\r' {
			i++
		} else if input[i] == 'L' {
			total_digits := Find_number_index(input, i)

			number, _ := strconv.Atoi(input[i+1 : i+total_digits])
			zero_crossings := (dial_value - number) / 100

			add_factor := 1
			if dial_value == 0 {
				add_factor = 0
			}
			if dial_value-number <= 0 {
				zero_count += (zero_crossings * -1) + add_factor
			}
			dial_value = (dial_value - number) % 100
			if dial_value < 0 {
				dial_value += 100
			}

			i += total_digits + 1

		} else if input[i] == 'R' {
			total_digits := Find_number_index(input, i)
			number, _ := strconv.Atoi(input[i+1 : i+total_digits])
			zero_crossings := (dial_value + number) / 100
			if (dial_value + number) >= 100 {
				zero_count += zero_crossings
			}
			dial_value = (dial_value + number) % 100

			i += total_digits + 1

		}

	}
	return zero_count
}
