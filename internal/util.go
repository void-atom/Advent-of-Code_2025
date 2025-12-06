package internal

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// Check if the string is composed of repeated substring present in it
// eg: if 121212 -> return true
// 	   if 12122 -> returns false

func RepeatedSubstring(number string) bool {
	test := number + number

	// fmt.Println(test, test[1:len(test)-1])
	value_to_return := strings.Contains(test[1:len(test)-1], number)
	// fmt.Println(value_to_return)
	return value_to_return

}

// This stack code is AI generated
type Stack[T any] struct {
	items []T
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack.
// It returns the item and a boolean indicating success, or an error if the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index] // Slice off the top element
	return item, nil
}

// Peek returns the top item of the stack without removing it.
// It returns the item and a boolean indicating success, or an error if the stack is empty.
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Print the values in stack
func (s *Stack[T]) Print() {
	fmt.Print("[")

	for i := 0; i < len(s.items); i++ {
		fmt.Print(s.items[i])
		if i < len(s.items)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}
