package main

import (
	"aoc2025/days"
	"aoc2025/internal"
	"fmt"
)

func main() {
	file_path := "input_files/day1.txt"
	input := internal.Read_file(file_path)
	fmt.Printf("Part one answer = %d\n", days.Day_one_part_one(input))
	fmt.Printf("Part two answer = %d\n", days.Day_one_part_two(input))

}
