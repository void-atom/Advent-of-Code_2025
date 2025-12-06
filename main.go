package main

import (
	"aoc2025/days"
	"aoc2025/internal"
	"fmt"
)

func main() {
	// change this choice value to the day you want the answer of
	// Example choice := 1 gives answer for day 1
	choice := 2
	solution_day := days.DaysMap[choice]

	file_path := fmt.Sprintf("input_files/day%d.txt", choice)
	input := internal.ReadFile(file_path)

	fmt.Printf("Day %d. Part 1:  %d\n", choice, solution_day.PartOne(input))
	fmt.Printf("Day %d. Part 2:  %d\n", choice, solution_day.PartTwo(input))

}
