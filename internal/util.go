package internal

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// CONSTANTS used throughout the project
const (
	FORWARD_LOOK_INDEX = 5
	INDEX_ZERO         = 0
	INDEX_ONE          = 1
	INDEX_TWO          = 2
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

// Count the total neighbouring cells with paper rolls
func CountTotalNeighbouringRolls(input *[]string, idx_x int, idx_y int) int {
	total := 0
	height := len(*input)
	width := len((*input)[0])

	dirs := [][2]int{
		{0, 1},   // Right
		{-1, 1},  // Up Right
		{-1, 0},  // Up
		{-1, -1}, // Up Left
		{0, -1},  // Left
		{1, -1},  // Down Left
		{1, 0},   // Down
		{1, 1},   // Down Right
	}

	for _, d := range dirs {
		ny := idx_y + d[0]
		nx := idx_x + d[1]
		if ny >= 0 && ny < height && nx >= 0 && nx < width {
			if (*input)[ny][nx] == '@' {
				total++
			}
		}
	}

	return total
}

// sort ranges based on the start and then the end
func SortRanges(range_list [][2]int) [][INDEX_TWO]int {

	sort.Slice(range_list, func(i, j int) bool {
		if range_list[i][INDEX_ZERO] != range_list[j][INDEX_ZERO] {
			return range_list[i][INDEX_ZERO] < range_list[j][INDEX_ZERO]
		} else {
			return range_list[i][INDEX_ONE] < range_list[j][INDEX_ONE]
		}
	})
	// fmt.Println(range_list)
	return range_list

}

// Future improvements: Maybe binary search can be used?
func SearchInRange(range_list [][2]int, search_id int) bool {
	for _, r := range range_list {
		if r[0] <= search_id && search_id <= r[1] {
			return true
		}
	}
	return false
}

func MergeRangesSorted(ranges_raw [][2]int) [][2]int {
	if len(ranges_raw) == 0 {
		return nil
	}
	ranges := SortRanges(ranges_raw)
	merged := make([][2]int, 0, len(ranges))
	merged = append(merged, ranges[0])

	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]
		if r[0] <= last[1] { // overlap
			if r[1] > last[1] {
				last[1] = r[1] // extend the last range
			}
		} else {
			merged = append(merged, r) // no overlap, add new range
		}
	}

	return merged
}
