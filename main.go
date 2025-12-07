package main

import (
	"aoc2025/days"
	"aoc2025/internal"
	"fmt"
)

func main() {
	// change this choice value to the day you want the answer of
	// Example choice := 1 gives answer for day 1
	choice := 5
	solution_day := days.DaysMap[choice]

	file_path := fmt.Sprintf("input_files/day%d.txt", choice)
	input := internal.ReadFile(file_path)

	// 1. Create channels to receive the results
	part1Result := make(chan int)
	part2Result := make(chan int)

	// 2. Concurrent func calls
	go func() {
		part1Result <- solution_day.PartOne(input)
	}()

	go func() {
		part2Result <- solution_day.PartTwo(input)
	}()

	// 3. Receive the results from the channels (this blocks until a result is available)
	result1 := <-part1Result
	result2 := <-part2Result

	// 4. Print the results
	fmt.Printf("Day %d. Part 1: %d\n", choice, result1)
	fmt.Printf("Day %d. Part 2: %d\n", choice, result2)

}
