package days

import (
	util "aoc2025/internal"
	"fmt"
	"strings"
)

func Day_two_part_one() {
	file_path := "input_files/day2.txt"
	input := strings.Split(util.Read_file(file_path), ",")
	fmt.Println(input)
}
